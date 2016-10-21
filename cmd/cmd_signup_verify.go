package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var signupVerifyToken string
var signupVerifyMail string

type jsonDataSignupVerify struct {
	Address string
	Token   string
}

var signupVerifyHandler = func(cmd *cobra.Command, args []string) {
	if signupVerifyToken == "" {
		fail(`Missing token. Use "--token" or see "--help"`)
	}

	if signupVerifyMail == "" {
		fail(`Missing email address. Use "--mail" or see "--help"`)
	}

	jsonData := jsonDataSignupVerify{signupVerifyMail, signupVerifyToken}
	if _, err := newRequest("/auth/user/verify").post(jsonData); err == nil {
		fmt.Println("Verified account for " + signupVerifyMail + "!")
	} else {
		fail("Failed to verify account.")
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
