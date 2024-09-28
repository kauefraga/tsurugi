package blocker

import (
	"io"
	"net/http"
	"net/url"
)

func isUrl(s string) bool {
	if len(s) == 0 {
		return false
	}

	url, err := url.Parse(s)
	if err != nil || url.Host == "" {
		return false
	}

	if url.Scheme == "" && url.Fragment == "" && url.Opaque == "" {
		return false
	}

	return url.Scheme == "http" || url.Scheme == "https"
}

func fetchHostsFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	file, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return file, nil
}
