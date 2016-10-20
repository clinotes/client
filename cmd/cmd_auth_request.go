package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var authRequestMail string

var authRequestHandler = func(cmd *cobra.Command, args []string) {
	if authRequestMail == "" {
		fail(`Missing email address. Use "--mail" or see "--help"`)
	}

	if _, err := postToAPI("/auth/token/create", `{"address":"`+APIUsername+`"}`); err == nil {
		fmt.Println("Requested a token for " + APIUsername + "! Please check your mails for your token …")
	} else {
		fail("Failed to request token.")
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
