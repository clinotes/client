package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
		fmt.Println("Requested a token for " + authRequestMail + "! Please check your mails for your token …")
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
