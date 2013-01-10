package GitHubParser

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

/* for parse article page
 * .md
 */
type ArticleParser struct {
	Parser
}

func NewArticleParser() *ArticleParser {
	return &ArticleParser{
		Parser{
			TagHead: "<article class=\"markdown-body entry-content",
			TagTail: "</article>",
		},
	}
}

func (this *ArticleParser) IsArticle(content string) bool {
	return strings.Contains(content, this.TagHead)
}

func (this *ArticleParser) Parse(url string, content string) error {
	posHead := strings.Index(content, this.TagHead)
	posTail := strings.Index(content, this.TagTail)

	if posHead < 0 || posTail < 0 {
		return &GHError{"Not GitHub .md file"}
	}

	file := LocationRoot + url[len(GitHubRepository):] + ".html"
	dir := filepath.Dir(file)

	if err := os.MkdirAll(dir, 0600); err != nil {
		return err
	}

	return ioutil.WriteFile(file, []byte(content[posHead:posTail+len(this.TagTail)]), 0600)
}
