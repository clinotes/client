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

	"github.com/spf13/cobra"
)

var signupMail string

type jsonDataSignup struct {
	Address string
}

var signupHandler = func(cmd *cobra.Command, args []string) {
	if signupMail == "" {
		fail(`Missing email address. Use "--mail" or see "--help"`)
	}

	jsonData := jsonDataSignup{signupMail}
	if err := newRequest("/account/create").post(jsonData); err == nil {
		fmt.Println("Created an account for " + signupMail + "! Please check your mails to verify your account …")
	} else {
		fail("Failed to create account.")
	}
}

var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "Sign up for a clinot.es account",
	Run:   signupHandler,
}

func init() {
	signupCmd.Flags().StringVar(&signupMail, "mail", "", "mail address")

	RootCmd.AddCommand(signupCmd)
}
