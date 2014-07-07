package main

import (
	"fmt"
	"time"
)

var since time.Time = time.Date(2014, 6, 1, 0, 0, 0, 0, time.UTC)

func main() {
	commits, err := GetCommits(since)

	if err == nil {
		for _, commit := range commits {
			fmt.Println(FormatMessage(*commit.Commit.Message))
		}
	}
}
