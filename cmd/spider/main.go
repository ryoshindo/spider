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

	exitCode, err := spider.Cli(ctx, spider.ParseCli)
	if err != nil {
		panic(err)
	}

	os.Exit(exitCode)
}
