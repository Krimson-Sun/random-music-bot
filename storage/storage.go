package storage

import "context"

type Storage interface {
	Save(ctx context.Context, m *Music) error
	PickRandom(ctx context.Context, UserName string) (*Music, error)
	IsExists(ctx context.Context, m *Music) (bool, error)
}

type Music struct {
	ID       string
	UserName string
}
