package main

import (
	"fmt"
	"time"
)

var since time.Time = time.Date(2014, 6, 1, 0, 0, 0, 0, time.UTC)

const FREQUENCY = time.Second * 8

func tick(prev time.Time) {
	fmt.Println(prev)
	/*
		commits, err := GetCommits(prev)

		if err == nil {
			for _, commit := range commits {
				fmt.Println(FormatMessage(*commit.Commit.Message))
			}
		}
	*/
}

func main() {
	ticker := time.NewTicker(FREQUENCY)

	go func() {
		tick(time.Now().Local().Add(-FREQUENCY))

		prev := time.Now().Local()
		for _ = range ticker.C {
			tick(prev)
			prev = time.Now().Local()
		}
	}()

	// TODO
	// Holds open the main routine, this is shitty
	timer := time.NewTimer(time.Hour * 10000)
	<-timer.C
}
