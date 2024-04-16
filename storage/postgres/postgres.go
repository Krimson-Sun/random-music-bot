package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"random-music-bot/storage"
)

type Storage struct {
	db *sql.DB
}

type dbParams struct {
	dbName   string
	host     string
	user     string
	password string
	port     int
}

func New(params dbParams) (*Storage, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("dbname=%s host=%s user=%s password=%s port=%d sslmode=disable",
		params.dbName, params.host, params.user, params.password, params.port))
	if err != nil {
		return nil, fmt.Errorf("err while opening db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("err while pinging db: %w", err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Save(ctx context.Context, m *storage.Music) error {
	_, err := s.db.ExecContext(ctx, "INSERT INTO music (username, id) values ($1, $2)", m.UserName, m.ID)
	if err != nil {
		return fmt.Errorf("err while inserting music: %w", err)
	}

	return nil
}

func (s *Storage) PickRandom(ctx context.Context, userName string) (*storage.Music, error) {
	var id string
	row := s.db.QueryRowContext(ctx, "SELECT id FROM music WHERE username = $1 ORDER BY random() LIMIT 1", userName)

	err := row.Scan(&id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("err while scanning music: %w", err)
	}

	m := storage.Music{
		ID:       id,
		UserName: userName,
	}
	return &m, nil
}

func (s *Storage) IsExists(ctx context.Context, m *storage.Music) (bool, error) {
	var count int

	row := s.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM music WHERE username = $1 AND id = $2", m.UserName, m.ID)
	if err := row.Scan(&count); err != nil {
		return false, fmt.Errorf("can't check music existance: %w", err)
	}

	return count > 0, nil
}

func (s *Storage) Init(ctx context.Context) error {
	_, err := s.db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS music (
		id VARCHAR(255) PRIMARY KEY,
		username VARCHAR(255)
	)`)
	if err != nil {
		return fmt.Errorf("err while creating table: %w", err)
	}

	return nil
}
