package main

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/ktakenaka/gosample/cmd/internal/config"
	"github.com/ktakenaka/gosample/cmd/internal/logger"
	ifrLogger "github.com/ktakenaka/gosample/infrastructure/logger"
)

func main() {
	cfg, err := config.Initialize()
	if err != nil {
		panic(err)
	}

	l, err := logger.Initialize(cfg)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	l.RequestErrorWithPerson(
		"userID",
		ifrLogger.WARN,
		&http.Request{
			Method: "GET",
			URL:    &url.URL{Scheme: "https"},
		},
		errors.New("dummy error"),
	)
}
