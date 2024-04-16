package telegram

import "random-music-bot/clients/telegram"

type Processor struct {
	tg     telegram.Client
	offset int
}

func New(client telegram.Client, storage) *Processor {
	return &Processor{
		tg: client,
	}
}
