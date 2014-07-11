package main

import "time"

const FREQUENCY = time.Minute * 30
const RANGE = time.Hour * 300
const DB_NAME = "__db.txt"

func tick() {
	since := time.Now().Local().Add(-RANGE)
	commits, err := GetCommits(since)
	if err != nil {
		panic(err)
	}

	for _, commit := range commits {
		if !GetSHA(DB_NAME, *commit.SHA) {
			AddSHA(DB_NAME, *commit.SHA)
			message := FormatMessage(commit)
			_, err := Tweet(message)
			if err != nil {
				//panic(err)
			}
		}
	}
}

func main() {
	ticker := time.NewTicker(FREQUENCY)

	// Every FREQUENCY, check GitHub for new commits.
	// The range of commits we check is FREQUENCY - RANGE, as GitHub queries
	// by commit date, not push, so in the case where a commit that's a few days old
	// finally is pushed, it's date is a few days ago, so we have a large range to account for
	// those cases
	go func() {
		tick()
		for _ = range ticker.C {
			tick()
		}
	}()

	// TODO
	// Holds open the main routine, this is shitty
	timer := time.NewTimer(time.Hour * 10000)
	<-timer.C
}
