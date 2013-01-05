package main

import (
	"GitHubParser"
	"fmt"
	"os"
	"path/filepath"
)

const (
	GitHubRoot string = "https://github.com"
)

var (
	GitHubRepository string
	pImage           GitHubParser.ImageParser
	pTree            GitHubParser.TreeParser
	pArticle         GitHubParser.ArticleParser
)

func init() {
	pImage = GitHubParser.ImageParser{}
	pImage.Init()

	pTree = GitHubParser.TreeParser{}
	pTree.Init()

	pArticle = GitHubParser.ArticleParser{}
	pArticle.Init()
}

func main() {
	params := os.Args
	if params == nil || len(params) <= 1 {
		fmt.Println("Usage: GitHubDoc <GitHub URL>")
		return
	}

	GitHubRepository = params[1]

	parse(GitHubRepository)
}

func parse(url string) {
	fmt.Println("Request: ", url)

	if pImage.IsImage(url) {

		file := "C:" + url[len(GitHubRepository):]
		dir := filepath.Dir(file)

		if err := os.MkdirAll(dir, 0600); err != nil {
			fmt.Println(err.Error())
			return
		}

		if err := GitHubParser.GetImage(pImage.Parse(url), file); err != nil {
			fmt.Println(err.Error())
			return
		}

	} else {

		content, err := GitHubParser.GetContent(url)
		if err != nil {
			fmt.Println("Usage: GitHubDoc <GitHub URL>")
			return
		}

		switch {
		case pTree.IsTree(content):
			urls, err := pTree.Parse(content)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			for _, u := range urls {
				parse(GitHubRoot + u)
			}

			return
		case pArticle.IsArticle(content):
			file := "C:" + url[len(GitHubRepository):] + ".html"
			dir := filepath.Dir(file)

			if err := os.MkdirAll(dir, 0600); err != nil {
				fmt.Println(err.Error())
				return
			}

			if err := pArticle.Parse(file, content); err != nil {
				fmt.Println(err.Error())
				return
			}

			return
		}
	}
}
