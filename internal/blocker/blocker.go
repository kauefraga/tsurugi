package blocker

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const HOSTS_PATH string = "/etc/hosts"

func getHosts(file []byte) []string {
	lines := strings.Split(string(file), "\n")

	var hosts []string

	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			hosts = append(hosts, line)
		}
	}

	return hosts
}

// TODO: check invalid url
func fetchHostsFile(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	file, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return getHosts(file), nil
}

// plain text -> hosts list
// obs.: empty strings/blank lines are being returned
func GetHosts(hostsFile string) ([]string, error) {
	isUrl := strings.Contains(hostsFile, "https") || strings.Contains(hostsFile, "http")

	if isUrl {
		return fetchHostsFile(hostsFile)
	}

	file, err := os.ReadFile(hostsFile)
	if err != nil {
		return nil, err
	}

	return getHosts(file), nil
}

// hosts list -> /etc/hosts
func WriteHosts(hosts []string) error {
	var blockList []string

	for _, host := range hosts {
		block := fmt.Sprintf("%s %s", "127.0.0.1", host)

		blockList = append(blockList, block)
	}

	blockFile := strings.Join(blockList, "\n")

	return os.WriteFile(HOSTS_PATH, []byte(blockFile), 0666)
}
