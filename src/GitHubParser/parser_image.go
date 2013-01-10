package GitHubParser

import (
	"os"
	"path/filepath"
	"strings"
)

/* for parse image page
 * https://github.com/astaxie/build-web-application-with-golang/blob/master/images/1.1.mac.png
 * ->
 * https://raw.github.com/astaxie/build-web-application-with-golang/master/images/1.1.cmd.png
 */
var image_exts = [...]string{".png"}

type ImageParser struct {
}

func NewImageParser() *ImageParser {
	return &ImageParser{}
}

func (this *ImageParser) IsImage(url string) bool {
	for _, ext := range image_exts {
		if strings.LastIndex(url, ext) == (len(url) - len(ext)) {
			return true
		}
	}
	return false
}

func (this *ImageParser) Parse(url string) error {
	file := LocationRoot + url[len(GitHubRepository):]
	dir := filepath.Dir(file)

	if err := os.MkdirAll(dir, 0600); err != nil {
		return err
	}

	uri := strings.Replace(url, GitHubRoot, GitHubRaw, 1)
	uri = strings.Replace(uri, "/blob/", "/", 1)

	return GetImage(uri, file)
}
