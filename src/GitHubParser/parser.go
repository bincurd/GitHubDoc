package GitHubParser

import (
	"fmt"
	"time"
)

var (
	GitHubRoot       string         = "https://github.com"
	GitHubRaw        string         = "https://raw.github.com"
	GitHubRepository string         = ""
	LocationRoot     string         = "C:"
	image            *ImageParser   = NewImageParser()
	tree             *TreeParser    = NewTreeParser()
	article          *ArticleParser = NewArticleParser()
	stat             int            = 0
	pool             chan int       = make(chan int, 16)
)

func Start(url string) {
	GitHubRepository = url

	t := time.Now()

	fire(url)
	for stat > 0 {
		time.Sleep(time.Duration(1) * time.Second)
	}

	fmt.Println("Completed: ", time.Now().Sub(t))
}

func fire(url string) {
	stat++
	go parse(url)
}

func parse(url string) {
	pool <- 1
	fmt.Println("Request: ", url)

	if image.IsImage(url) {

		if err := image.Parse(url); err != nil {
			fmt.Println("Error: ", err.Error(), url)
			return
		}

	} else {

		content, err := GetContent(url)
		if err != nil {
			fmt.Println("Error: ", err.Error(), url)
			return
		}

		switch {
		case tree.IsTree(content):
			urls, err := tree.Parse(content)
			if err != nil {
				fmt.Println("Error: ", err.Error(), url)
				return
			}

			for _, u := range urls {
				fire(GitHubRoot + u)
			}
		case article.IsArticle(content):
			if err := article.Parse(url, content); err != nil {
				fmt.Println("Error: ", err.Error(), url)
				return
			}
		}
	}

	<-pool
	stat--
}
