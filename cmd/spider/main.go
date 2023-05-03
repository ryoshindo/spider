package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/ryoshindo/spider"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	opt := &spider.Option{
		ConfigFilePath: "../spider-test/spider.yml",
	}

	app, err := spider.New(context.Background(), opt)
	if err != nil {
		panic(err)
	}

	// if err := app.Apply(ctx); err != nil {
	// 	panic(err)
	// }

	if err := app.Destroy(ctx); err != nil {
		panic(err)
	}
}
