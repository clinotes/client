package feedback

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

// Ask for the user input
func Ask(message string) string {
	Info(message)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("  > ")
	text, _ := reader.ReadString('\n')

	return strings.TrimSpace(text)
}

// Info prints out an information
func Info(message string) {
	fmt.Fprintf(os.Stderr, "%s "+message+"\n", color.New(color.FgHiCyan).SprintFunc()("⋅"))
}

// Done prints out a confirmation
func Done(message string) {
	fmt.Fprintf(os.Stderr, "%s "+message+"\n", color.New(color.FgGreen).SprintFunc()("✓"))
}

// Fail prints an error and stops the process
func Fail(message string) {
	fmt.Fprintf(os.Stderr, "%s "+message+"\n", color.New(color.FgRed).SprintFunc()("✘"))
	os.Exit(0)
}
