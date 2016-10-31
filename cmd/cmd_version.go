package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type versionCmdResponse struct {
	Version string `json:"version"`
	Client  string `json:"client"`
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show clinot.es client version",
	Run: func(cmd *cobra.Command, args []string) {
		// Check if running a local checkout or the built command
		prefix := "v"
		if VERSION == "latest" {
			prefix = "@"
		}

		fmt.Printf("client: %s%s\n", prefix, VERSION)
		fmt.Printf("server: ")

		if APIHostname != "" {
			r, err := http.Get(APIHostname + "/version")

			if err == nil {
				defer r.Body.Close()
				decoder := json.NewDecoder(r.Body)
				var data versionCmdResponse
				err = decoder.Decode(&data)

				if err == nil && data.Version != "" && data.Client != "" {
					fmt.Printf("v%s  (supports client >= v%s)\n", data.Version, data.Client)
					return
				}
			}
		}

		fmt.Printf("n/a\n")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
