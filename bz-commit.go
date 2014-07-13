package main

import (
	"strings"

	"github.com/google/go-github/github"
)

type BzCommit struct {
	Message   string
	BugNumber string
	Component string
	SHA       string
}

func NewBzCommit(repoCommit *github.RepositoryCommit) (*BzCommit, error) {
	c := new(BzCommit)

	c.Message = *repoCommit.Commit.Message
	c.SHA = *repoCommit.SHA
	c.BugNumber = GetBugNumber(c.Message)

	bugData, err := GetBugData(c.BugNumber)

	if err != nil {
		return c, err
	}

	component, err := bugData.Get("component").String()
	c.Component = component

	return c, nil
}

// Returns the tweetable message for the commit
func (c *BzCommit) FormatMessage() string {
	return CreateMessage(c.Message, c.BugNumber)
}

// Confirms whether or not this bug is in the Developer Tools component
func (c BzCommit) IsValid() bool {
	return strings.Index(c.Component, "Developer Tools") == 0
}
