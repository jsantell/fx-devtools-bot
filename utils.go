package main

import (
	"github.com/google/go-github/github"
)
import "regexp"

// MAX_LENGTH determines how long a commit message can be after trimmed,
// while leaving room for the URL and trailing ellipses to remain
// under 140 characters.
const MAX_LENGTH = 112

// TODO probably can group these regexps together in some cases, parsing out chunks
// of a commit message
var bugRegExp = regexp.MustCompile("(?i)[\\s]*bug [\\d]{4,9}[:\\-\\s]*")
var reviewRegExp = regexp.MustCompile("(?i)\\W*[\\s]+(a|r)=[\\s\\w,]+$")

var bugNumberRegExp = regexp.MustCompile("^(?i)[\\s]*bug ([\\d]{1,9})")
var changesetRegExp = regexp.MustCompile("^(?i)Backed out changeset")
var mergeRegExp = regexp.MustCompile("^(?i)merge ")
var bumpRegExp = regexp.MustCompile("^(?i)Bumping (gaia|mani)")

// Replaces bug number and reviewd comments from string.
func CleanMessage(str string) string {
	str = bugRegExp.ReplaceAllLiteralString(str, "")
	str = reviewRegExp.ReplaceAllLiteralString(str, "")
	return str
}

// Creates a Bugzilla URI from a commit message.
func CreateBugzillaURL(str string) string {
	result := bugNumberRegExp.FindStringSubmatch(str)
	if len(result) == 2 {
		return "http://bugzil.la/" + result[1]
	} else {
		return ""
	}
}

// Takes a commit and formats it into a digestable tweet.
func FormatMessage(commit github.RepositoryCommit) string {
	message := *commit.Commit.Message
	url := CreateBugzillaURL(message)
	message = CleanMessage(message)

	if len(message) > MAX_LENGTH {
		message = message[0:MAX_LENGTH] + "..."
	}

	return message + " " + url
}

// Returns a boolean indicating whether or not this commit message
// is useful, ignoring merges, backouts and automated commits.
func IsValidCommit(message string) bool {
	return !changesetRegExp.MatchString(message) &&
		!mergeRegExp.MatchString(message) &&
		!bumpRegExp.MatchString(message)
}
