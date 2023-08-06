package net

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NetCmd represents the net command
var NetCmd = &cobra.Command{
	Use:   "net",
	Short: "access net utils",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("net called")
		cmd.Help()
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// NetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// NetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
