package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var authToken string
var authMail string

var authHandler = func(cmd *cobra.Command, args []string) {
	if authToken == "" {
		fail(`Missing token. Use "--token" or see "--help"`)
	}

	if authMail == "" {
		fail(`Missing email address. Use "--mail" or see "--help"`)
	}

	if _, err := postToAPI("/auth", `{"address":"`+authMail+`", "token": "`+authToken+`"}`); err == nil {
		fmt.Println("Token is valid for " + authMail + "!")
	} else {
		fail("Failed to request token.")
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
