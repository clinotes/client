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
	"strings"

	fb "github.com/sbstjn/feedback"
	"github.com/spf13/cobra"
)

type jsonDataSignup struct {
	Address string
}

func signupHandler(cmd *cobra.Command, args []string) {
	if len(args) != 1 || args[0] == "" || !strings.Contains(args[0], "@") {
		fb.Fail("Invalid email address. Use `cn signup mail@example.com`")
	}

	jsonData := jsonDataSignup{args[0]}
	if err := newRequest("/account/create").post(jsonData); err == nil {
		fb.Done("A verification token will be delivered to your inbox!")
		verifyAccount(args[0], fb.Ask("Please enter your token:"))
		requestToken(args[0])
	} else {
		fb.Fail("Failed to create account!")
	}
}

func init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "signup",
		Short: "Sign up for a clinot.es account",
		Run:   signupHandler,
	})
}
