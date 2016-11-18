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
	"fmt"
	"io/ioutil"
	"os/user"

	fb "github.com/sbstjn/feedback"
	"github.com/spf13/cobra"
)

type jsonDataAuth struct {
	Address string
	Token   string
}

type jsonDataAuthRequest struct {
	Address string
}

func checkCredentials(address string, token string) {
	jsonData := jsonDataAuth{address, token}
	if err := newRequest("/auth").post(jsonData); err == nil {
		fb.Done("Token is valid for " + jsonData.Address + "!")

		if address != APIAddress && token != APIToken {
			saveCredentials(APIHostname, address, token)
		}
	} else {
		fb.Fail("Failed to authorize token.")
	}
}

func requestToken(address string) {
	jsonData := jsonDataAuthRequest{address}
	if err := newRequest("/token/create").post(jsonData); err == nil {
		fb.Done("An access token will be delivered to your inbox!")

		checkCredentials(address, fb.Ask("Please enter your token:"))
	} else {
		fb.Fail("Failed to request token!")
	}
}

func saveCredentials(hostname string, address string, token string) {
	config := []byte(fmt.Sprintf(
		"CLINOTES_API_HOSTNAME: %s\nCLINOTES_API_USERNAME: %s\nCLINOTES_API_TOKEN: %s",
		hostname,
		address,
		token,
	))

	// Get current system user
	usr, err := user.Current()
	if err != nil {
		fb.Fail("Could not access home directory!")
	}

	// Write data in $HOME/.clinotes.yaml
	err = ioutil.WriteFile(usr.HomeDir+"/.clinotes.yaml", config, 0644)
	if err != nil {
		fb.Fail(`Failed to store credentials in ~/.clinotes.yaml`)
	}

	// Done
	fb.Done("Stored credentials in ~/.clinotes.yaml")
}

func authHandler(cmd *cobra.Command, args []string) {
	switch len(args) {
	case 0:
		fb.Info("Using credentials from ~/.clinotes.yaml …")
		checkCredentials(APIAddress, APIToken)
	case 1:
		fb.Info("Requesting access token …")
		requestToken(args[0])
	case 2:
		fb.Info("Checking token …")
		checkCredentials(args[0], args[1])
	default:
		fb.Fail("Invalid parameter. Use `cn auth [mail@example.com] [YourToken]`")
	}
}

func init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "auth",
		Short: "Authorize clinot.es client",
		Run:   authHandler,
	})
}
