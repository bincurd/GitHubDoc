package GitHubParser

import (
	"io/ioutil"
	"strings"
)

type Parser interface {
	Parse(string, string) error
}

type MdParser struct {
	TagHead string
	TagTail string
}

func init() {
	TagHead := "<artical "
	TagTail := "</artical>"
}

func (this *MdParser) Parse(file string, content string) error {
	posHead := strings.Index(content, this.TagHead)
	posTail := strings.Index(content, this.TagTail)

	if posHead < 0 || posTail < 0 {
		return &GHError{"Not GitHub .md file"}
	}

	return ioutil.WriteFile(file, []byte(content[posHead:posTail+len(this.TagTail)]), 0600)
}
