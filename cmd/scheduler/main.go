package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"employee_attendance/internal/scheduler"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	schedule, err := scheduler.NewScheduler(ctx)
	if err != nil {
		log.Fatalf("scheduler.NewScheduler: %v", err)
	}

	err = schedule.Run(ctx)
	if err != nil {
		log.Fatalf("scheduler.Run: %v", err)
	}
}
