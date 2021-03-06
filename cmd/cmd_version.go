/**
 * clinot.es client
 * Copyright (C) 2016 Sebastian Müller
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

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

func versionHandler(cmd *cobra.Command, args []string) {
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
				fmt.Printf("v%s (supports client >= v%s)\n", data.Version, data.Client)
				return
			}
		}
	}

	fmt.Printf("n/a\n")
}

func init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Show clinot.es client version",
		Run:   versionHandler,
	})
}
