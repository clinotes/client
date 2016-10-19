package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// VERSION is set during build
	VERSION string
	// APIHostname is the hostname for the API
	APIHostname string
	// APIUsername is the username for the API
	APIUsername string
	// APIToken is the username for the API
	APIToken string
)

// APIErrorResponse stores information from the API
type APIErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "clinotes",
	Short: "Access clinot.es from the command line",
}

func fail(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, "error: "+format+"\n", a...)
	os.Exit(1)
}

func postToAPI(action string, jsonString string) (*APIErrorResponse, error) {
	req, err := http.NewRequest("POST", APIHostname+action, bytes.NewBuffer([]byte(jsonString)))
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, errors.New("Failed to send request")
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("Received HTTP error code")
	}

	var data APIErrorResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	defer resp.Body.Close()

	if err != nil {
		return nil, errors.New("Failed to parse response")
	}

	return &data, nil
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
	APIUsername = viper.GetString("CLINOTES_API_USERNAME")
	APIToken = viper.GetString("CLINOTES_API_TOKEN")
}
