package main

import (
	"GitHubParser"
	"fmt"
	"os"
)

func main() {
	params := os.Args
	if params == nil || len(params) <= 1 {
		fmt.Println("Usage: GitHubDoc <GitHub URL>")
		return
	}

	ghRoot, err := GitHubParser.GetContent(params[1])
	if err != nil {
		fmt.Println("Usage: GitHubDoc <GitHub URL>")
		return
	}

	fmt.Println(ghRoot)
}
