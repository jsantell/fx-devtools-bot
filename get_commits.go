package main

import (
	"fmt"
	"time"

	"github.com/google/go-github/github"
)

const BRANCH = "fx-team"
const REPO_OWNER = "mozilla"
const REPO_NAME = "gecko-dev"
const CLIENT_PATH = "browser/devtools"
const SERVER_PATH = "toolkit/devtools"

func GetCommits(since time.Time) ([]github.RepositoryCommit, error) {
	client := github.NewClient(nil)

	allCommits := []github.RepositoryCommit{}

	for _, path := range []string{CLIENT_PATH, SERVER_PATH} {
		fmt.Println("Getting commits for " + path)
		opts := &github.CommitsListOptions{SHA: BRANCH, Since: since, Path: path}
		commits, _, err := client.Repositories.ListCommits(REPO_OWNER, REPO_NAME, opts)
		fmt.Println("Got commits for " + path)

		if err != nil {
			return nil, err
		}

		for _, commit := range commits {
			// TODO ensure that the SHA hasn't already been appended, in the case
			// where both server and client sides are in the same commit
			allCommits = append(allCommits, commit)
		}

	}

	for _, commit := range allCommits {
		fmt.Println(*commit.Commit.Author.Name, *commit.Commit.Message)
	}

	return allCommits, nil
}
