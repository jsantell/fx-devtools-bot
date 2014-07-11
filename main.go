package main

import "time"

var since time.Time = time.Date(2014, 6, 1, 0, 0, 0, 0, time.UTC)

const FREQUENCY = time.Hour * 2
const RANGE = time.Hour * 96

func tick() {
	since := time.Now().Local().Add(-RANGE)
	commits, err := GetCommits(since)
	if err == nil {
		for _, commit := range commits {
			Tweet(FormatMessage(commit))
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
