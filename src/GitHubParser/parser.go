package GitHubParser

import (
	"io/ioutil"
	"strings"
)

type Parser struct {
	TagHead string
	TagTail string
}

type TreeParser struct {
	Parser
}

func (this *TreeParser) Init() {
	this.TagHead = "<tbody class=\"tree-entries"
	this.TagTail = "</tbody>"
}

func (this *TreeParser) IsTree(content string) bool {
	return strings.Contains(content, this.TagHead)
}

func (this *TreeParser) Parse(content string) ([]string, error) {
	posHead := strings.Index(content, this.TagHead)
	posTail := strings.Index(content, this.TagTail)

	if posHead < 0 || posTail < 0 {
		return nil, &GHError{"Not GitHub tree: " + string(posHead) + ", " + string(posTail)}
	}

	tbody := content[posHead:posTail]

	tagAH := "<td class=\"content\"><a href=\""
	tagAT := "\" class="
	urls := []string{}

	for {
		posAH := strings.Index(tbody, tagAH)
		if posAH < 0 {
			break
		}

		tbody = tbody[posAH+len(tagAH):]

		posAT := strings.Index(tbody, tagAT)
		if posAT < 0 {
			break
		}

		urls = append(urls, tbody[:posAT])

		tbody = tbody[posAT+len(tagAT):]
	}

	if len(urls) <= 0 {
		return nil, &GHError{"Not GitHub tree"}
	}

	return urls, nil
}

type ArticleParser struct {
	Parser
}

func (this *ArticleParser) Init() {
	this.TagHead = "<article class=\"markdown-body entry-content"
	this.TagTail = "</article>"
}

func (this *ArticleParser) IsArticle(content string) bool {
	return strings.Contains(content, this.TagHead)
}

func (this *ArticleParser) Parse(file string, content string) error {
	posHead := strings.Index(content, this.TagHead)
	posTail := strings.Index(content, this.TagTail)

	if posHead < 0 || posTail < 0 {
		return &GHError{"Not GitHub .md file"}
	}

	return ioutil.WriteFile(file, []byte(content[posHead:posTail+len(this.TagTail)]), 0600)
}
