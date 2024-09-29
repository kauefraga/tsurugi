package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/kauefraga/tsurugi/internal/blocker"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:     "stop",
	Short:   "Unblock websites",
	Aliases: []string{"unblock", "allow"},
	Run: func(cmd *cobra.Command, args []string) {
		currentHostsFile, err := os.ReadFile(blocker.HOSTS_PATH)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var newHostsFile []string
		isTsurugiBlockList := false

		lines := strings.Split(string(currentHostsFile), "\n")
		for _, line := range lines {
			if strings.Contains(line, "# tsurugi blocks start") {
				isTsurugiBlockList = true
			}

			if strings.Contains(line, "# tsurugi blocks end") {
				isTsurugiBlockList = false
				continue
			}

			if !isTsurugiBlockList {
				newHostsFile = append(newHostsFile, line)
			}
		}

		os.WriteFile(
			blocker.HOSTS_PATH,
			[]byte(strings.Join(newHostsFile, "\n")),
			0666)
	},
}
