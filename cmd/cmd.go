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
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// VERSION is set during build
	VERSION string
	// APIHostname is the hostname for the API
	APIHostname string
	// APIAddress is the username for the API
	APIAddress string
	// APIToken is the username for the API
	APIToken string
	// FallbackHostname is used if no custom hostname
	FallbackHostname = "https://api.clinot.es"
)

// APIErrorResponse stores information from the API
type APIErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Done    bool   `json:"done"`
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "clinotes",
	Short: "Access clinot.es from the command line",
}

func ensureCredentials() {
	if APIAddress == "" || APIToken == "" {
		fail("No credentials found!\n\nSee `cn --help` or visit https://clinot.es to read more …")
	}
}

func doneNice(message string) {
	fmt.Fprintf(os.Stderr, " %s "+message+"\n", color.New(color.FgGreen).SprintFunc()("✓"))
}

func failNice(message string) {
	fmt.Fprintf(os.Stderr, " %s "+message+"\n", color.New(color.FgRed).SprintFunc()("✘"))
	os.Exit(1)
}

func fail(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, "error: "+format+"\n", a...)
	os.Exit(1)
}

// Execute adds all child commands to the root command
func Execute(version string) {
	VERSION = version

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigName(".clinotes")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	APIHostname = viper.GetString("CLINOTES_API_HOSTNAME")
	APIAddress = viper.GetString("CLINOTES_API_USERNAME")
	APIToken = viper.GetString("CLINOTES_API_TOKEN")
}
