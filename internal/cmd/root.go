package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
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
			color.Red("Error: could not get hosts list")
			os.Exit(1)
		}

		err = blocker.WriteHosts(hosts)
		if err != nil {
			color.Red("Error: could not block hosts")
			fmt.Println("Might be lack of permissions, try again with sudo")
			fmt.Println(
				color.MagentaString("$"),
				color.HiBlackString("sudo tsurugi %s", args[0]),
			)
			os.Exit(1)
		}

		fmt.Println("New blocked hosts:")
		for _, host := range hosts {
			color.Green("+ %s", host)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
