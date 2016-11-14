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

import "github.com/spf13/cobra"

var authRequestMail string

type jsonDataAuthRequest struct {
	Address string
}

var authRequestHandler = func(cmd *cobra.Command, args []string) {
	if authRequestMail == "" {
		fail(`Missing email address. Use "--mail" or see "--help"`)
	}

	jsonData := jsonDataAuthRequest{authRequestMail}
	if err := newRequest("/token/create").post(jsonData); err == nil {
		doneNice("Requested a token for " + authRequestMail + "! Please check your mails for your token …")
	} else {
		failNice("Failed to request token.")
	}
}

var authRequestCmd = &cobra.Command{
	Use:   "request",
	Short: "Authorize clinot.es client",
	Run:   authRequestHandler,
}

func init() {
	authRequestCmd.Flags().StringVar(&authRequestMail, "mail", "", "mail address")

	authCmd.AddCommand(authRequestCmd)
}
