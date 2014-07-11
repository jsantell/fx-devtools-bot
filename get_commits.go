package main

import (
	"fmt"
	"os"
	"time"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
)

var githubToken string = os.Getenv("FX_DEVTOOLS_BOT_GITHUB_TOKEN")

const BRANCH = "fx-team"
const REPO_OWNER = "mozilla"
const REPO_NAME = "gecko-dev"
const CLIENT_PATH = "browser/devtools"
const SERVER_PATH = "toolkit/devtools"
const STYLES_PATH = "browser/themes/shared/devtools"

func GetCommits(since time.Time) ([]github.RepositoryCommit, error) {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: githubToken},
	}
	client := github.NewClient(t.Client())

	allCommits := []github.RepositoryCommit{}

	for _, path := range []string{CLIENT_PATH, SERVER_PATH, STYLES_PATH} {
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
			if !isDuplicate && IsValidCommit(*commit.Commit.Message) {
				allCommits = append(allCommits, commit)
			}
		}
	}

	filtered := FilterCommits(DB_NAME, allCommits)
	fmt.Println("Fetched commits:", len(filtered))

	return filtered, nil
}

func FilterCommits(dbName string, commits []github.RepositoryCommit) []github.RepositoryCommit {
	filtered := []github.RepositoryCommit{}

	for _, commit := range commits {
		if GetSHA(dbName, *commit.SHA) == false {
			filtered = append(filtered, commit)
		}
	}

	return filtered
}
