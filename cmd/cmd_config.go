package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configHandler = func(cmd *cobra.Command, args []string) {
	if APIHostname != "" {
		fmt.Println("CLINOTES_API_HOSTNAME: " + APIHostname)
	}

	if APIAddress != "" {
		fmt.Println("CLINOTES_API_USERNAME: " + APIAddress)
	}

	if APIToken != "" {
		fmt.Println("CLINOTES_API_TOKEN: " + APIToken)
	}
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show clinot.es client configuration",
	Run:   configHandler,
}

func init() {
	RootCmd.AddCommand(configCmd)
}
