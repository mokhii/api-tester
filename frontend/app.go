package main

import (
	"context"
	"fmt"
    "time"

    "github.com/go-resty/resty/v2"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// SendRequest performs an HTTP request and returns a short summary
func (a *App) SendRequest(url string, method string) string {
    client := resty.New()
    start := time.Now()
    resp, err := client.R().Execute(method, url)
    if err != nil {
        return fmt.Sprintf("Error: %v", err)
    }
    duration := time.Since(start)
    return fmt.Sprintf("Status: %d, Time: %v, Size: %d", resp.StatusCode(), duration, len(resp.Body()))
}
