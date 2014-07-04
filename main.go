package main

import (
	"fmt"
	"time"
)

// TODO Figure out why `time.Time` can't be `const`
var since time.Time = time.Date(2014, 6, 1, 0, 0, 0, 0, time.UTC)

func main() {
	fmt.Println("Getting commits...")
	GetCommits(since)

	fmt.Println("Done getting commits.")
}
