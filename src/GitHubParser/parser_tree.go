package GitHubParser

import (
	"strings"
)

/* for parse tree page
 * most of tree page has article section else
 * so check tree first
 */
type TreeParser struct {
	Parser
}

func NewTreeParser() *TreeParser {
	return &TreeParser{
		Parser{
			TagHead: "<tbody class=\"tree-entries",
			TagTail: "</tbody>",
		},
	}
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
