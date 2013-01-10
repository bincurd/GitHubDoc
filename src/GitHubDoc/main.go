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

	GitHubParser.Start(params[1])
}
