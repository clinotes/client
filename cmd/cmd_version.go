package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show clinot.es client version",
	Run: func(cmd *cobra.Command, args []string) {
		// Check if running a local checkout or the built command
		if VERSION == "latest" {
			fmt.Println(RootCmd.Use + " @" + VERSION)
		} else {
			fmt.Println(RootCmd.Use + " v" + VERSION)
		}
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
