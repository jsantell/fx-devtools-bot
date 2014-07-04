package main

import (
	"fmt"
	"time"

	"github.com/google/go-github/github"
)

const BRANCH = "fx-team"
const PATH = "browser/devtools"

// TODO Figure out why `time.Time` can't be `const`
var since time.Time = time.Date(2014, 6, 1, 0, 0, 0, 0, time.UTC)

func main() {
	client := github.NewClient(nil)
	opts := &github.CommitsListOptions{SHA: BRANCH, Since: since, Path: PATH}
	repos, _, err := client.Repositories.ListCommits("mozilla", "gecko-dev", opts)

	for i := 0; i < len(repos); i++ {
		repo := repos[i]
		fmt.Println(i)
		fmt.Println(repo)
		fmt.Println("\n\n\n\n\n\n\n\n")
		/*
			if repo.Author != nil {
				fmt.Println(*repo.Author.Login)
			}
			if repo.Message != nil {
				fmt.Println(repo.Message)
			}
		*/
	}

	fmt.Println(err)
}
