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
	fb "github.com/sbstjn/feedback"
	"github.com/spf13/cobra"
)

type jsonDataSignupVerify struct {
	Address string
	Token   string
}

func verifyAccount(address string, token string) {
	if address == "" || token == "" {
		fb.Fail("Invalid parameter. Use `cn verify mail@example.com YourToken`")
	}

	jsonData := jsonDataSignupVerify{address, token}
	if err := newRequest("/account/verify").post(jsonData); err == nil {
		fb.Done("Account verified!")
	} else {
		fb.Fail("Failed to verify account!")
	}
}

func signupVerifyHandler(cmd *cobra.Command, args []string) {
	switch len(args) {
	case 1:
		verifyAccount(args[0], fb.Ask("Please enter your token:"))
	case 2:
		verifyAccount(args[0], args[1])
	default:
		fb.Fail("Invalid parameter. Use `cn verify mail@example.com [YourToken]`")
	}
}

func init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "verify",
		Short: "Verify a new clinot.es account",
		Run:   signupVerifyHandler,
	})
}
