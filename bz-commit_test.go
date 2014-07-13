// Keeping these tests light, as they do ping the bugzilla API.

package main

import (
	"testing"

	"github.com/google/go-github/github"
)

func createGHCommit(message string, sha string) *github.RepositoryCommit {
	repoCommit := new(github.RepositoryCommit)
	commit := new(github.Commit)
	commit.Message = &message
	commit.SHA = &sha
	repoCommit.SHA = &sha
	repoCommit.Commit = commit
	return repoCommit
}

func TestNewBzCommit(t *testing.T) {
	message := "Bug 1034670 - The canvas graphs should draw the background separately, r=pbrosset"
	sha := "abcdefgh"
	ghCommit := createGHCommit(message, sha)
	bz, _ := NewBzCommit(ghCommit)

	if bz.Message != message {
		t.Errorf("Expected `Message` property '%v' to equal '%v'", bz.Message, message)
	}
	if bz.BugNumber != "1034670" {
		t.Errorf("Expected `BugNumber` property '%v' to equal '%v'", bz.BugNumber, "1034670")
	}
	if bz.Component != "Developer Tools: Profiler" {
		t.Errorf("Expected `Component` property '%v' to equal '%v'", bz.Component, "Developer Tools: Profiler")
	}
	if bz.SHA != sha {
		t.Errorf("Expected `SHA` property '%v' to equal '%v'", bz.SHA, sha)
	}
}

func TestBzCommit_FormatMessage(t *testing.T) {
	messages := [][]string{
		[]string{
			"Bug 1034670 - The canvas graphs should draw the background separately, r=pbrosset",
			"The canvas graphs should draw the background separately http://bugzil.la/1034670",
		},
		[]string{
			"Bug 1034670: The canvas graphs should draw the background separately and one two three four five six seven eight nine ten eleven twelve thirteen fourteen fifteen sixteen, r=pbrosset",
			"The canvas graphs should draw the background separately and one two three four five six seven eight nine ten ele... http://bugzil.la/1034670",
		},
	}

	for _, duple := range messages {
		commit, _ := NewBzCommit(createGHCommit(duple[0], "abcdefg"))
		formatted := commit.FormatMessage()
		if formatted != duple[1] {
			t.Error("Expected formatted message\n'" + formatted + "'\nto have been formatted to\n'" + duple[1] + "'")
		}
	}
}

func TestBzCommit_IsValid(t *testing.T) {
	dtCommit, _ := NewBzCommit(createGHCommit("bug 980506 - some stuff going here", "abcdefg"))
	nonDTCommit, _ := NewBzCommit(createGHCommit("bug 1000000: happy one million, bring your own drinks", "abcdefgfg"))

	if dtCommit.IsValid() != true {
		t.Error("Expected valid Developer Tools commit IsValid() to be valid.")
	}
	if nonDTCommit.IsValid() == true {
		t.Error("Expected non Developer Tools commit IsValid() to be invalid.")
	}
}
