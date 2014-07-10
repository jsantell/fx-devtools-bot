package main

import (
	"bufio"
	"os"
)

func AddSHA(filename string, sha string) {
	// Create database file if DNE
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		_, err = os.Create(filename)
		if err != nil {
			panic(err)
		}
	}

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	if _, err = file.WriteString(sha + "\n"); err != nil {
		panic(err)
	}
}

func GetSHA(filename string, sha string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == sha {
			return true
		}
	}

	return false
}
