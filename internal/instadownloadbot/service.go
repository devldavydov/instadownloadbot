package instadownloadbot

import (
	"context"
	"fmt"
	"time"

	tele "gopkg.in/telebot.v3"
)

type Service struct {
	settings ServiceSettings
}

func NewService(settings ServiceSettings) (*Service, error) {
	return &Service{settings: settings}, nil
}

func (s *Service) Run(ctx context.Context) error {
	pref := tele.Settings{
		Token:  s.settings.Token,
		Poller: &tele.LongPoller{Timeout: s.settings.PollTimeOut},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		return err
	}

	b.Handle("/start", func(c tele.Context) error {
		return c.Send(fmt.Sprintf("Hello! Current unix time: %d", time.Now().Unix()))
	})

	b.Start()

	return nil
}
