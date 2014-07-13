package main

import (
	"os"
	"time"

	"github.com/jsantell/go-githubstream"
)

const FREQUENCY = time.Hour
const OWNER = "mozilla"
const REPO = "gecko-dev"
const BRANCH = "fx-team"
const DB_NAME = "__db.txt"

var TOKEN string = os.Getenv("FX_DEVTOOLS_BOT_GITHUB_TOKEN")

/*
		if !GetSHA(DB_NAME, *commit.SHA) {
			AddSHA(DB_NAME, *commit.SHA)
			message := FormatMessage(commit)
			fmt.Println("Tweeting: ", message)
			_, err := Tweet(message)
			if err != nil {
				//panic(err)
			}
		}
	}
}
*/
func main() {
	ghs := githubstream.NewGithubStream(FREQUENCY, OWNER, REPO, BRANCH, TOKEN)

	for commits := range ghs.Start() {
		// Filter out already used commits from DB_NAME, and invalid commits
		// like merges, changesets, automated commits to store on overhead of querying
		// bugzilla
		commits = FilterCommits(DB_NAME, commits)
	}
}
