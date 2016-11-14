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

var signupVerifyToken string
var signupVerifyMail string

type jsonDataSignupVerify struct {
	Address string
	Token   string
}

var signupVerifyHandler = func(cmd *cobra.Command, args []string) {
	if signupVerifyToken == "" {
		failNice(`Missing token. Use "--token" or see "--help"`)
	}

	if signupVerifyMail == "" {
		failNice(`Missing email address. Use "--mail" or see "--help"`)
	}

	jsonData := jsonDataSignupVerify{signupVerifyMail, signupVerifyToken}
	if err := newRequest("/account/verify").post(jsonData); err == nil {
		doneNice("Verified account for " + signupVerifyMail + "!")
	} else {
		failNice("Failed to verify account.")
	}
}

var signupVerifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify a new clinot.es account",
	Run:   signupVerifyHandler,
}

func init() {
	signupVerifyCmd.Flags().StringVar(&signupVerifyToken, "token", "", "token for account verification")
	signupVerifyCmd.Flags().StringVar(&signupVerifyMail, "mail", "", "mail for account verification")

	signupCmd.AddCommand(signupVerifyCmd)
}
