package main

import (
	"strings"

	"github.com/google/go-github/github"
)

import "net/http"
import "github.com/bitly/go-simplejson"

import "regexp"

// MAX_LENGTH determines how long a commit message can be after trimmed,
// while leaving room for the URL and trailing ellipses to remain
// under 140 characters.
const MAX_LENGTH = 112

var compMap = map[string]string{
	"3D View":                         "Tilt",
	"Canvas Debugger":                 "Canvas",
	"Graphic Commandline and Toolbar": "gcli",
	"Object Inspector":                "Inspector", // Alias for brevity and clarity to end users
	"Responsive Mode":                 "Responsive",
	"Web Audio Editor":                "Audio",
	"WebGL Shader Editor":             "Shader",
}

var cleanRegExp = regexp.MustCompile("(?i)[\\s]*bug [\\d]{4,9}[\\s]*[\\:\\-]+[\\s]*(.*\\w)\\W*\\s(a|r)=")
var stripBugRegExp = regexp.MustCompile("(?i)[\\s]*bug [\\d]{4,9}[\\s]*[\\:\\-]+(.*)")
var bugNumberRegExp = regexp.MustCompile("^(?i)[\\s]*bug ([\\d]{1,9})")
var changesetRegExp = regexp.MustCompile("^(?i)Backed out changeset")
var mergeRegExp = regexp.MustCompile("^(?i)merge ")
var bumpRegExp = regexp.MustCompile("^(?i)Bumping (gaia|mani)")
var subcomponentRegExp = regexp.MustCompile("^(?i)Developer Tools: (.*)$")

// Returns a boolean indicating whether or not this commit message
// is useful, ignoring merges, backouts and automated commits.
func IsValidCommit(message string) bool {
	return !changesetRegExp.MatchString(message) &&
		!mergeRegExp.MatchString(message) &&
		!bumpRegExp.MatchString(message)
}

// Takes a DB file store path and a slice of github.RepositoryCommits
// and returns a slice of commits that are valid based off of:
//
// * Have not yet been tweeted
// * Is not a changeset, merge or bump commit
func FilterCommits(dbName string, commits []github.RepositoryCommit) []github.RepositoryCommit {
	var filtered []github.RepositoryCommit

	for _, commit := range commits {
		if IsValidCommit(*commit.Commit.Message) &&
			GetSHA(dbName, *commit.SHA) == false {

			filtered = append(filtered, commit)
		}
	}

	return filtered
}

func GetJson(url string) (*simplejson.Json, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return simplejson.NewFromReader(res.Body)
}

func CreateMessage(message string, bugNumber string, subcomponent string) string {
	url := GetBugzillaURL(bugNumber)
	result := cleanRegExp.FindStringSubmatch(message)

	// Clean up the bug number and everything after the message, like the `r=jsantell` comments and after
	if len(result) > 1 {
		message = result[1]
		// If couldn't parse, just remove the bug number and leave everything else
	} else {
		result = stripBugRegExp.FindStringSubmatch(message)
		// If still not cleaned, who cares, let Twitter deal with it
		if len(result) > 1 {
			message = result[1]
		}
	}

	message = strings.Trim(message, " ")

	if subcomponent != "" {
		message = "[" + subcomponent + "]" + " " + message
	}

	if len(message) > MAX_LENGTH {
		message = message[0:MAX_LENGTH] + "..."
	}

	return message + " " + url
}

func GetBugzillaURL(bugNum string) string {
	if bugNum != "" {
		return "http://bugzil.la/" + bugNum
	} else {
		return ""
	}
}

func GetBugData(bugNum string) (*simplejson.Json, error) {
	j, err := GetJson("https://bugzilla.mozilla.org/rest/bug/" + bugNum)

	if err != nil {
		return nil, err
	}

	return j.Get("bugs").GetIndex(0), nil
}

// Get bug number from a commit message
func GetBugNumber(str string) string {
	result := bugNumberRegExp.FindStringSubmatch(str)
	if len(result) == 2 {
		return result[1]
	} else {
		return ""
	}
}

func GetSubComponent(c string) string {
	result := subcomponentRegExp.FindStringSubmatch(c)

	if len(result) > 1 {
		mapping := compMap[result[1]]
		if mapping != "" {
			return mapping
		} else {
			return result[1]
		}
	}

	return ""
}
