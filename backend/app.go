package main

import (
	"api-tester/backend/internal"
	"context"
)

type App struct{}

func (a *App) Startup(ctx context.Context) {}

func (a *App) SendRequest(url string, method string) string {
	return internal.Send(url, method)
}
