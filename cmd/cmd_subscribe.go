package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type jsonDataSubscribe struct {
	Address string
	Token   string
	Number  string
	Expire  string
	CVC     string
}

var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribe a clinot.es pro account",
	Run: func(cmd *cobra.Command, args []string) {
		// Fail of either APIAddress or APIToken is missing
		ensureCredentials()

		reader := bufio.NewReader(os.Stdin)

		fmt.Println("")
		fmt.Print("   Subscribing to paid clinot.es account …\n")
		fmt.Println("")
		fmt.Print("   card number: ")
		cardNumber, _ := reader.ReadString('\n')

		fmt.Print("   exp (mo/yr): ")
		cardExpire, _ := reader.ReadString('\n')

		fmt.Print("           cvc: ")
		cardCVC, _ := reader.ReadString('\n')

		fmt.Println("")
		fmt.Print("   … processing payment request\n")

		jsonData := jsonDataSubscribe{APIAddress, APIToken, strings.TrimSpace(cardNumber), strings.TrimSpace(cardExpire), strings.TrimSpace(cardCVC)}
		if err := newRequest("/subscribe").post(jsonData); err == nil {
			fmt.Println("   … Done. Thank you so much!")
			fmt.Println("")
		} else {
			fmt.Println("   … Failed. Something went wrong, Sorry!")
			fmt.Println("")
		}
	},
}

func init() {
	RootCmd.AddCommand(subscribeCmd)
}
