package cmd

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var (
	url    string
	method string
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a single HTTP request",
	Run: func(cmd *cobra.Command, args []string) {
		client := resty.New()
		start := time.Now()

		resp, err := client.R().
			EnableTrace().
			Execute(method, url)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		duration := time.Since(start)

		fmt.Printf("Status: %d\n", resp.StatusCode())
		fmt.Printf("Duration: %v\n", duration)
		fmt.Printf("Size: %d bytes\n", len(resp.Body()))
	},
}

func init() {
	sendCmd.Flags().StringVarP(&url, "url", "u", "", "Request URL")
	sendCmd.Flags().StringVarP(&method, "method", "m", "GET", "HTTP method")
	sendCmd.MarkFlagRequired("url")
}
