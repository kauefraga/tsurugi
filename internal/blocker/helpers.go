package blocker

import "strings"

func getHostsList(file []byte) []string {
	lines := strings.Split(string(file), "\n")

	var hosts []string

	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			hosts = append(hosts, line)
		}
	}

	return hosts
}
