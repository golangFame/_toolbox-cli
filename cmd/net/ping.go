/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var hostUrl string
var client http.Client

func ping(hostURL string) (statusCode int, err error) {
	url := "https://" + hostURL + "/ping"

	resp, err := client.Head(url)
	if err != nil {
		return
	}

	return resp.StatusCode, err
}

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Pings a host",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ping called")

		client = http.Client{
			Timeout: time.Second * 2,
		}

		statusCode, err := ping(hostUrl)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Status Code: %d\n", statusCode)
		}

	},
}

func init() {

	pingCmd.Flags().StringVarP(&hostUrl, "hostUrl", "u", "", "Host to ping")

	if err := pingCmd.MarkFlagRequired("hostUrl"); err != nil {
		fmt.Println(err)
	}

	NetCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
