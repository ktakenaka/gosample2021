package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/ktakenaka/gosample/cmd/config"
	"github.com/ktakenaka/gosample/cmd/db"
	"github.com/ktakenaka/gosample/cmd/logger"
	"github.com/ktakenaka/gosample/ent/office"
	ifrLogger "github.com/ktakenaka/gosample/infrastructure/logger"
)

func main() {
	cfg, err := config.Initialize()
	if err != nil {
		log.Fatalf("failed initializing cofig: %v", err)
	}
	cfg.DB.IsLogging = true

	client, err := db.Initialize(cfg)
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	office, err := client.Office.Query().Where(office.ID(1)).Only(ctx)
	if err != nil {
		panic(err)
	}

	samples, err := office.QuerySamples().All(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", samples)

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
