package GitHubParser

import (
	"io/ioutil"
	"strings"
)

type Parser struct {
	TagHead string
	TagTail string
}

/* for parse tree page
 * most of tree page has article section else
 * so check tree first
 */
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
		return nil, &GHError{"Not GitHub tree"}
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

/* for parse article page
 * .md
 */
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

/* for parse image page
 * https://github.com/astaxie/build-web-application-with-golang/blob/master/images/1.1.mac.png
 * ->
 * https://raw.github.com/astaxie/build-web-application-with-golang/master/images/1.1.cmd.png
 */
type ImageParser struct {
	Exts []string
}

func (this *ImageParser) Init() {
	this.Exts = append(this.Exts, ".png")
}

func (this *ImageParser) IsImage(url string) bool {
	for _, ext := range this.Exts {
		if strings.LastIndex(url, ext) == (len(url) - len(ext)) {
			return true
		}
	}
	return false
}

func (this *ImageParser) Parse(url string) string {
	uri := strings.Replace(url, "https://github.com/", "https://raw.github.com/", 1)
	return strings.Replace(uri, "/blob/", "/", 1)
}
