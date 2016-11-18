/**
 * clinot.es client
 * Copyright (C) 2016 Sebastian Müller
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

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
	// RootCmd.AddCommand(subscribeCmd)
}
