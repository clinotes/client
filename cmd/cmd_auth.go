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

	"github.com/spf13/cobra"
)

var authToken string
var authMail string

type jsonDataAuth struct {
	Address string
	Token   string
}

var authHandler = func(cmd *cobra.Command, args []string) {
	var jsonData jsonDataAuth

	if authToken != "" || authMail != "" {
		// Show error if no token is set
		if authToken == "" {
			failNice(`Missing token. Use "--token" or see "--help"`)
		}

		// Show error if no mail address is set
		if authMail == "" {
			failNice(`Missing email address. Use "--mail" or see "--help"`)
		}

		jsonData = jsonDataAuth{authMail, authToken}
	} else {
		// Use data from ~/.clinotes.yaml if available or raise error
		if APIAddress != "" && APIToken != "" {
			doneNice("Using credentials from ~/.clinotes.yaml …")

			jsonData = jsonDataAuth{APIAddress, APIToken}
		} else {
			failNice("Credentials in ~/.clinotes.yaml not valid")
		}
	}

	if err := newRequest("/auth").post(jsonData); err == nil {
		doneNice("Token is valid for " + jsonData.Address + "!")

		// Prepare configuration content
		config := []byte(fmt.Sprintf(
			"CLINOTES_API_HOSTNAME: %s\nCLINOTES_API_USERNAME: %s\nCLINOTES_API_TOKEN: %s",
			APIHostname,
			jsonData.Address,
			jsonData.Token,
		))

		// Get current system user
		usr, err := user.Current()
		if err != nil {
			failNice("Could not access home directory!")
		}

		// Write data in $HOME/.clinotes.yaml
		err = ioutil.WriteFile(usr.HomeDir+"/.clinotes.yaml", config, 0644)
		if err != nil {
			failNice(`Failed to store credentials in ~/.clinotes.yaml`)
		}

		// Done
		doneNice("Stored credentials in ~/.clinotes.yaml")
	} else {
		failNice("Failed to authorize token.")
	}
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authorize clinot.es client",
	Run:   authHandler,
}

func init() {
	authCmd.Flags().StringVar(&authMail, "mail", "", "mail address")
	authCmd.Flags().StringVar(&authToken, "token", "", "pass a valid auth token")

	RootCmd.AddCommand(authCmd)
}
