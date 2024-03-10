package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/zSnails/VRAirHockey-Backend/db"
	"github.com/zSnails/VRAirHockey-Backend/router"
	"github.com/zSnails/VRAirHockey-Backend/session"
)

var (
	wait   time.Duration
	logger = logrus.WithField("service", "entry-point")
)

func init() {

	logger.Logger.SetFormatter(&logrus.TextFormatter{})
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	flag.DurationVar(&wait, "wait", time.Second*10, "Graceful shutdown timeout")
	flag.Parse()
}

func main() {
	err := db.Init()
	if err != nil {
		panic(err)
	}

	session.Setup()

	server := http.Server{
		Addr:                         ":8080",
		Handler:                      router.NewRouter(),
		DisableGeneralOptionsHandler: false,
		WriteTimeout:                 time.Second * 15,
		ReadTimeout:                  time.Second * 15,
		IdleTimeout:                  time.Second * 15,
	}

	go func() {
		logger.Info("Listening on address", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			logger.Println(err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	logger.Info("Shutting down the server")
	err = server.Shutdown(ctx)
	if err != nil {
		logger.Fatal(err)
	}
}
