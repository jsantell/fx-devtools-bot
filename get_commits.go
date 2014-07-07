package main

import (
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
		opts := &github.CommitsListOptions{SHA: BRANCH, Since: since, Path: path}
		commits, _, err := client.Repositories.ListCommits(REPO_OWNER, REPO_NAME, opts)

		if err != nil {
			return nil, err
		}

		// ensure that the SHA hasn't already been appended, in the case
		// where both server and client sides are in the same commit
		for _, commit := range commits {
			isDuplicate := false
			for _, storedCommit := range allCommits {
				if *storedCommit.SHA == *commit.SHA {
					isDuplicate = true
				}
			}
			if !isDuplicate {
				allCommits = append(allCommits, commit)
			}
		}
	}

	return allCommits, nil
}
