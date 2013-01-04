package GitHubParser

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	GitHubRoot string = "https://github.com/"
)

func GetContent(url string) (string, error) {
	if !strings.Contains(url, GitHubRoot) {
		return "", &GHError{"Not GitHub URL"}
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
