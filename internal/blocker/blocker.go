package blocker

import (
	"fmt"
	"os"
	"strings"
)

const HOSTS_PATH string = "/etc/hosts"

// read plain text -> create hosts list
func GetHosts(file string) ([]string, error) {
	if isUrl(file) {
		hostsFile, err := fetchHostsFile(file)

		return getHostsList(hostsFile), err
	}

	hostsFile, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return getHostsList(hostsFile), nil
}

// read hosts list -> write /etc/hosts
func WriteHosts(hosts []string) error {
	var blockList []string

	for _, host := range hosts {
		block := fmt.Sprintf("%s %s", "127.0.0.1", host)

		blockList = append(blockList, block)
	}

	currentHostsFile, err := os.ReadFile(HOSTS_PATH)
	if err != nil {
		return err
	}

	newHostsFile := string(currentHostsFile) + fmt.Sprint(
		"# tsurugi blocks start (DO NOT REMOVE THIS COMMENT)\n",
		strings.Join(blockList, "\n")+"\n",
		"# tsurugi blocks end (DO NOT REMOVE THIS COMMENT)\n",
	)

	return os.WriteFile(HOSTS_PATH, []byte(newHostsFile), 0666)
}

// restore /etc/hosts
func RemoveHosts() error {
	currentHostsFile, err := os.ReadFile(HOSTS_PATH)
	if err != nil {
		return err
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

	newHostsFileContent := strings.Join(newHostsFile, "\n")

	return os.WriteFile(HOSTS_PATH, []byte(newHostsFileContent), 0666)
}
