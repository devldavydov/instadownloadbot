package instadownloadbot

import (
	"bytes"
	"fmt"
	"net/url"

	"github.com/devldavydov/instadownloadbot/internal/instadownloadbot/images"
	tele "gopkg.in/telebot.v3"
)

func (s *Service) onStart(c tele.Context) error {
	return c.Send(
		fmt.Sprintf("Hello, %s!\nWelcome to InstaDownloadBot!\nBuild: %s", c.Sender().Username, s.settings.BuildCommit),
	)
}

func (s *Service) onText(c tele.Context) error {
	instaURL := c.Text()
	if _, err := url.ParseRequestURI(instaURL); err != nil {
		return c.Send("Invalid URL")
	}

	if err := c.Send("We are working on it :)"); err != nil {
		return err
	}

	img := &tele.Photo{File: tele.FromReader(bytes.NewReader(images.UnderCon))}
	return c.Send(img)
}
