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

	blockFile := strings.Join(blockList, "\n")

	return os.WriteFile(HOSTS_PATH, []byte(blockFile), 0666)
}
