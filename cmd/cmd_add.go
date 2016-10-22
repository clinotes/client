package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addStart bool
var addStop bool

type jsonDataAdd struct {
	Address string
	Token   string
	Note    string
}

var addHandler = func(cmd *cobra.Command, args []string) {
	// Fail of either APIAddress or APIToken is missing
	ensureCredentials()

	// Make sure nobody uses --start and --stop at the same time
	if addStart == true && addStop == true {
		fail(`Use "--start" OR "--stop" but not both!`)
	}

	// Read note message from arguments
	note := strings.Join(args, " ")

	data := jsonDataAdd{APIAddress, APIToken, note}
	if err := newRequest("/add").post(data); err == nil {
		fmt.Println("Posted")
	} else {
		fail("Failed to submit note.")
	}
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a note",
	Run:   addHandler,
}

func init() {
	addCmd.Flags().BoolVar(&addStart, "start", false, "start counter")
	addCmd.Flags().BoolVar(&addStop, "stop", false, "stop counter")

	RootCmd.AddCommand(addCmd)
}
