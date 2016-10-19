package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var signupMail string

var signupHandle = func(cmd *cobra.Command, args []string) {
	if signupMail == "" {
		fail(`Missing email address. Use "--mail" or see "--help"`)
	}

	if _, err := postToAPI("/auth/user/create", `{"address":"`+signupMail+`"}`); err == nil {
		fmt.Println("Created an account for " + signupMail + "! Please check your mails to verify your account …")
	} else {
		fail("Failed to create account.")
	}
}

var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "Sign up for a clinot.es account",
	Run:   signupHandle,
}

func init() {
	signupCmd.Flags().StringVar(&signupMail, "mail", "", "mail address")

	RootCmd.AddCommand(signupCmd)
}
