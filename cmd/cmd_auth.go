package cmd

import (
	"fmt"
	"io/ioutil"
	"os/user"

	"github.com/spf13/cobra"
)

var authToken string
var authMail string

type jsonDataAuth struct {
	Address string
	Token   string
}

var authHandler = func(cmd *cobra.Command, args []string) {
	var jsonData jsonDataAuth

	if authToken != "" || authMail != "" {
		// Show error if no token is set
		if authToken == "" {
			fail(`Missing token. Use "--token" or see "--help"`)
		}

		// Show error if no mail address is set
		if authMail == "" {
			fail(`Missing email address. Use "--mail" or see "--help"`)
		}

		jsonData = jsonDataAuth{authMail, authToken}
	} else {
		// Use data from ~/.clinotes.yaml if available or raise error
		if APIAddress != "" && APIToken != "" {
			fmt.Println("Using credentials from ~/.clinotes.yaml …")

			jsonData = jsonDataAuth{APIAddress, APIToken}
		} else {
			fail("Credentials in ~/.clinotes.yaml not valid")
		}
	}

	if err := newRequest("/auth").post(jsonData); err == nil {
		fmt.Println("Token is valid for " + jsonData.Address + "!")

		// Prepare configuration content
		config := []byte(fmt.Sprintf(
			"CLINOTES_API_USERNAME: %s\nCLINOTES_API_TOKEN: %s",
			jsonData.Address,
			jsonData.Token,
		))

		// Get current system user
		usr, err := user.Current()
		if err != nil {
			fail(`Could not access home directory!`)
		}

		// Write data in $HOME/.clinotes.yaml
		err = ioutil.WriteFile(usr.HomeDir+"/.clinotes.yaml", config, 0644)
		if err != nil {
			fmt.Println(err)
			fail(`Failed to store credentials in ~/.clinotes.yaml`)
		}

		// Done
		fmt.Println("Stored credentials in ~/.clinotes.yaml")
	} else {
		fail("Failed to authorize token.")
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
