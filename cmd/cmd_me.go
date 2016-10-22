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
	Notes        int
	Token        int
	Subscription bool
}

var meHandler = func(cmd *cobra.Command, args []string) {
	// Fail of either APIAddress or APIToken is missing
	ensureCredentials()

	// Make sure nobody uses --start and --stop at the same time
	if addStart == true && addStop == true {
		fail(`Use "--start" OR "--stop" but not both!`)
	}

	// Read note message from arguments
	note := strings.Join(args, " ")

	data := jsonDataAdd{APIAddress, APIToken, note}
	var resp jsonDataMeResponse
	if err := newRequest("/me").postScan(data, &resp); err == nil {
		fmt.Printf(
			"\n Account: %s\n Created: %s\n Notes:   %d\n Token:   %d\n\n Subscription: %v\n\n",
			resp.Address,
			resp.Created,
			resp.Notes,
			resp.Token,
			resp.Subscription,
		)
	} else {
		fail("Unable to retrieve account data")
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
