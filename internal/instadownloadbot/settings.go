package instadownloadbot

import "time"

type ServiceSettings struct {
	Token       string
	PollTimeOut time.Duration
}

func NewServiceSettings(token string, pollTimeout time.Duration) (*ServiceSettings, error) {
	return &ServiceSettings{
		Token:       token,
		PollTimeOut: pollTimeout,
	}, nil
}
