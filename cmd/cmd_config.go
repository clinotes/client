package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show clinot.es client configuration",
	Run: func(cmd *cobra.Command, args []string) {
		if APIHostname != "" {
			fmt.Println("CLINOTES_API_HOSTNAME: " + APIHostname)
		}

		if APIUsername != "" {
			fmt.Println("CLINOTES_API_USERNAME: " + APIUsername)
		}

		if APIToken != "" {
			fmt.Println("CLINOTES_API_TOKEN: " + APIToken)
		}
	},
}

func init() {
	RootCmd.AddCommand(configCmd)
}
