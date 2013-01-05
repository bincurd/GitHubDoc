package GitHubParser

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func GetContent(url string) (string, error) {
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

func GetImage(url string, file string) error {
	resp, err1 := http.Get(url)
	if err1 != nil {
		return err1
	}
	defer resp.Body.Close()

	out, err2 := os.Create(file)
	if err2 != nil {
		return err2
	}
	defer out.Close()

	_, err3 := io.Copy(out, resp.Body)
	if err3 != nil {
		return err3
	}

	return nil
}
