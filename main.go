package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jsantell/go-githubstream"
)

const FREQUENCY = time.Hour * 1
const OWNER = "mozilla"
const REPO = "gecko-dev"
const BRANCH = "fx-team"
const DB_NAME = "__db.txt"

var TOKEN string = os.Getenv("FX_DEVTOOLS_BOT_GITHUB_TOKEN")

func main() {
	ghs := githubstream.NewGithubStream(FREQUENCY, OWNER, REPO, BRANCH, TOKEN)

	for commits := range ghs.Start() {
		fmt.Println(time.Now().Local())
		fmt.Println("Fetched commits: ", len(commits))
		// Filter out already used commits from DB_NAME, and invalid commits
		// like merges, changesets, automated commits to store on overhead of querying
		// bugzilla
		filtered := FilterCommits(DB_NAME, commits)
		fmt.Println("Non-merge/changeset/automated commits: ", len(filtered))

		for _, commit := range filtered {
			c, err := NewBzCommit(&commit)

			// If error created during creation of BzCommit,
			// Bugzilla couldn't be queried from either not finding a bug in the
			// commit message, or the API being down, or something else. Whatever the case
			// just ignore for now.
			if err != nil {
				fmt.Println("Failed creating BzCommit: " + c.Message)
				continue
			}

			// In this case if it's valid, meaning in the Developer Tools component
			// in this case, tweet it.
			if c.IsValid() {
				fmt.Println("Tweeting Valid Commit: " + c.Message)
				_, err := Tweet(c.FormatMessage())

				if err == nil {
					fmt.Println("Adding SHA for " + c.Message)
					AddSHA(DB_NAME, c.SHA)
				}
			}
		}

	}
}
