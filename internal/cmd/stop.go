package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/kauefraga/tsurugi/internal/blocker"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:     "stop",
	Short:   "Unblock websites",
	Aliases: []string{"unblock", "allow"},
	Run: func(cmd *cobra.Command, args []string) {
		err := blocker.RemoveHosts()
		if err != nil {
			color.Red("Error: could not unblock hosts")
			fmt.Println("Might be lack of permissions, try again with sudo")
			fmt.Println(
				color.MagentaString("$"),
				color.HiBlackString("sudo tsurugi stop"),
			)
			os.Exit(1)
		}

		color.Green("You can access the websites again, enjoy!")
	},
}
