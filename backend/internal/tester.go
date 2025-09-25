// internal/tester.go
package internal

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

func Send(url, method string) string {
	client := resty.New()
	start := time.Now()
	resp, err := client.R().Execute(method, url)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	duration := time.Since(start)
	return fmt.Sprintf("Status: %d, Time: %v, Size: %d",
		resp.StatusCode(), duration, len(resp.Body()))
}
