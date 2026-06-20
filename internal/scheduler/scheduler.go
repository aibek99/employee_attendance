package scheduler

import (
	"context"
	"fmt"

	"employee_attendance/internal/app"
)

type Scheduler struct {
	a *app.App
}

func NewScheduler(ctx context.Context) (*Scheduler, error) {
	a, err := app.NewApp(ctx)
	if err != nil {
		return nil, fmt.Errorf("app.NewApp: %w", err)
	}

	return &Scheduler{
		a: a,
	}, nil
}

func (s *Scheduler) Run(ctx context.Context) error {
	return nil
}
