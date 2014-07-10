package main

import (
	"os"
	"testing"

	"github.com/google/go-github/github"
)

func seed() {
	os.Remove(TEST_DB_NAME)

	for _, sha := range SHAS {
		AddSHA(TEST_DB_NAME, sha)
	}
}

func isIn(commit github.RepositoryCommit, commits []github.RepositoryCommit) bool {
	for _, c := range commits {
		if commit.SHA == c.SHA {
			return true
		}
	}
	return false
}

func createCommit(sha string) github.RepositoryCommit {
	commit := new(github.RepositoryCommit)
	commit.SHA = &sha
	return *commit
}

func TestFilterCommits(t *testing.T) {
	seed()
	commits := []github.RepositoryCommit{
		createCommit(SHAS[0]),
		createCommit(SHAS[3]),
		createCommit("799dc73909bee352b3645cbc753c48c772e514d3"),
	}

	filtered := FilterCommits(TEST_DB_NAME, commits)

	if isIn(commits[0], filtered) == true {
		t.Error("Expected previously stored commit to be removed from filtered commits.")
	}
	if isIn(commits[1], filtered) == true {
		t.Error("Expected previously stored commit to be removed from filtered commits.")
	}
	if isIn(commits[2], filtered) == false {
		t.Error("Expected previously unstored commit to be in filtered commits.")
	}

	os.Remove(TEST_DB_NAME)
}
