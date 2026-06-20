package bot

import (
	"context"
	"fmt"

	"employee_attendance/internal/app"
)

type Bot struct {
	a *app.App
}

func NewBot(ctx context.Context) (*Bot, error) {
	a, err := app.NewApp(ctx)
	if err != nil {
		return nil, fmt.Errorf("app.NewApp: %w", err)
	}

	return &Bot{
		a: a,
	}, nil
}

func (b *Bot) Run(ctx context.Context) error {
	return nil
}
