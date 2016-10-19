package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var authRequest bool
var authToken string

var authHandle = func(cmd *cobra.Command, args []string) {
	if !authRequest && authToken == "" {
		fail(`Use either "--request" or "--token" for auth command. See "--help" for more`)
	}

	if authRequest {
		if _, err := postToAPI("/auth/token/create", `{"address":"`+APIUsername+`"}`); err == nil {
			fmt.Println("Requested a token for " + APIUsername + "! Please check your mails for your token …")
		} else {
			fail("Failed to request token.")
		}
	} else {
		if _, err := postToAPI("/auth", `{"address":"`+APIUsername+`", "token": "`+authToken+`"}`); err == nil {
			fmt.Println("Authenticated with valid token for " + APIUsername + "!")
		} else {
			fail("Failed to authorize account.")
		}
	}
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authorize clinot.es client",
	Run:   authHandle,
}

func init() {
	authCmd.Flags().BoolVar(&authRequest, "request", false, "request a new auth token")
	authCmd.Flags().StringVar(&authToken, "token", "", "pass a valid auth token")

	RootCmd.AddCommand(authCmd)
}
