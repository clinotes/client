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
	if _, err := newRequest("/auth/user/create").post(jsonData); err == nil {
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
