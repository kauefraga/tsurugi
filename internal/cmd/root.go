package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/kauefraga/tsurugi/internal/blocker"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "tsurugi blocklist",
	Version: "1.0.0",
	Short:   "Easily block distracting websites using tsurugi",
	Example: "tsurugi\ntsurugi blocklist.txt\ntsurugi url",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hosts, err := blocker.GetHosts(args[0])
		if err != nil {
			fmt.Println("Error: could not get hosts list")
			os.Exit(1)
		}

		fmt.Println("The following hosts are being blocked:")
		fmt.Println(strings.Join(hosts, "\n"))

		err = blocker.WriteHosts(hosts)
		if err != nil {
			fmt.Println("Could not block hosts. Try running tsurugi with sudo.")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
