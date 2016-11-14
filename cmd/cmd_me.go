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
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type jsonDataMe struct {
	Address string
	Token   string
}

type jsonDataMeResponse struct {
	Address      string
	Created      time.Time
	Subscription bool
}

var meHandler = func(cmd *cobra.Command, args []string) {
	// Fail of either APIAddress or APIToken is missing
	ensureCredentials()

	// Make sure nobody uses --start and --stop at the same time
	if addStart == true && addStop == true {
		failNice(`Use "--start" OR "--stop" but not both!`)
	}

	// Read note message from arguments
	note := strings.Join(args, " ")

	data := jsonDataAdd{APIAddress, APIToken, note}
	var resp jsonDataMeResponse
	if err := newRequest("/account").postScan(data, &resp); err == nil {
		sub := "No"
		if resp.Subscription == true {
			sub = "Yes"
		}

		fmt.Printf(
			"\n Account: %s\n Created: %s\n\n Subscription: %v\n\n",
			resp.Address,
			resp.Created,
			sub,
		)
	} else {
		failNice("Unable to retrieve account data")
	}
}

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Show account information",
	Run:   meHandler,
}

func init() {

	RootCmd.AddCommand(meCmd)
}
