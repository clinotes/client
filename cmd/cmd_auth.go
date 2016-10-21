package cmd

import (
	"fmt"
	"io/ioutil"
	"os/user"

	"github.com/spf13/cobra"
)

var authToken string
var authMail string

var authHandler = func(cmd *cobra.Command, args []string) {
	var fromParameter bool
	var token, address string

	if authToken != "" || authMail != "" {
		fromParameter = true

		if authToken == "" {
			fail(`Missing token. Use "--token" or see "--help"`)
		}

		if authMail == "" {
			fail(`Missing email address. Use "--mail" or see "--help"`)
		}

		token = authToken
		address = authMail
	} else {
		fromParameter = false

		if APIUsername != "" && APIToken != "" {
			fmt.Println("Using credentials from ~/.clinotes.yaml …")

			token = APIToken
			address = APIUsername
		}
	}

	if _, err := postToAPI("/auth", `{"address":"`+address+`", "token": "`+token+`"}`); err == nil {
		fmt.Println("Token is valid for " + address + "!")

		if fromParameter {
			config := []byte(fmt.Sprintf("CLINOTES_API_USERNAME: %s\nCLINOTES_API_TOKEN: %s", address, token))

			usr, err := user.Current()
			if err != nil {
				fail(`Could not access home directory!`)
			}

			err = ioutil.WriteFile(usr.HomeDir+"/.clinotes.yaml", config, 0644)
			if err != nil {
				fmt.Println(err)
				fail(`Failed to store credentials in ~/.clinotes.yaml`)
			}

			fmt.Println("Stored credentials in ~/.clinotes.yaml")
		}
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
