package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	handler "carma_guard/internal/handler"
	"carma_guard/util"

	"github.com/go-telegram/bot"
)

func main() {

	conf, err := util.LoadConfig("./configs/")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler.DefaultHandler),
	}

	b, err := bot.New(conf.BOT_TOKEN, opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}
