package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"employee_attendance/internal/bot"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	app, err := bot.NewBot(ctx)
	if err != nil {
		log.Fatalf("bot.NewBot: %v", err)
	}

	err = app.Run(ctx)
	if err != nil {
		log.Fatalf("bot.Run: %v", err)
	}
}
