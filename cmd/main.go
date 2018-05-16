package main

import (
	"github.com/fatih/color"
	"github.com/sawadashota/waitopen"
	"github.com/spf13/cobra"
	"net/url"
	"os"
)

var (
	urlString string
	interval  int
	retry     int
)

func main() {
	listenOptions()
	rootCmd.Execute()
}

func listenOptions() {
	rootCmd.Flags().StringVarP(&urlString, "URL", "u", "", "URL")
	rootCmd.Flags().IntVarP(&interval, "interval", "i", 30, "Interval")
	rootCmd.Flags().IntVarP(&retry, "retry", "r", 5, "Retry")
}

var rootCmd = &cobra.Command{
	Use:   "wait-open",
	Short: "Wait until access then open URL",
	Run: func(cmd *cobra.Command, args []string) {
		if urlString == "" {
			color.Red("Invalid URL or empty")
			os.Exit(1)
		}

		URL, err := url.Parse(urlString)

		if err != nil {
			color.Red("%v", err.Error())
			os.Exit(1)
		}

		client := waitopen.New(URL)
		retryOption := waitopen.SetRetry(retry)
		intervalOption := waitopen.SetInterval(interval)

		client.WaitOpen(retryOption, intervalOption)
	},
}
