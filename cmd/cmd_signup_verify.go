package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var signupVerifyToken string
var signupVerifyMail string

var signupVerifyHandle = func(cmd *cobra.Command, args []string) {
	if signupVerifyToken == "" {
		fail(`Missing token. Use "--token" or see "--help"`)
	}

	if signupVerifyMail == "" {
		fail(`Missing email address. Use "--mail" or see "--help"`)
	}

	if _, err := postToAPI("/auth/user/verify", `{"address":"`+signupVerifyMail+`", "token": "`+signupVerifyToken+`"}`); err == nil {
		fmt.Println("Verified account for " + signupVerifyMail + "!")
	} else {
		fail("Failed to verify account.")
	}
}

var signupVerifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify a new clinot.es account",
	Run:   signupVerifyHandle,
}

func init() {
	signupVerifyCmd.Flags().StringVar(&signupVerifyToken, "token", "", "token for account verification")
	signupVerifyCmd.Flags().StringVar(&signupVerifyMail, "mail", "", "mail for account verification")

	signupCmd.AddCommand(signupVerifyCmd)
}
