package main

import (
	"context"
	"flag"
	"log"
	"random-music-bot/clients/telegram"
	"random-music-bot/storage/postgres"
)

const hostTgBot = "api.telegram.org"

func mustToken() string {
	// bot -tg-bot-token
	token := flag.String("tg-bot-token", "", "token for tg bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("err while getting token")
	}
	return *token
}

func main() {
	var dbParams = postgres.dbParams{}

	s, err := postgres.New()
	if err != nil {
		log.Fatalf("err while creating storage: %v", err)
	}

	s.Init(context.TODO())

	tgClient = telegram.New(hostTgBot, mustToken())

	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor)
}
