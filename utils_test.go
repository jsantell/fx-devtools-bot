package main

import (
	"testing"

	"github.com/google/go-github/github"
)

func createCommitMessage(message string) github.RepositoryCommit {
	repoCommit := new(github.RepositoryCommit)
	commit := new(github.Commit)
	commit.Message = &message
	repoCommit.Commit = commit
	return *repoCommit
}

func TestIsValidCommit(t *testing.T) {
	invalid := []string{
		"merge fx-team to mozilla-central a=merge",
		"Merge mozilla-central to fx-team",
		"Bumping manifests a=b2g-bump",
		"Bumping gaia.json for 2 gaia revision(s) a=gaia-bump",
		"Backed out changeset e01dbdf8a218 (bug 1016613)",
	}

	valid := []string{
		"Bug 1034670 - The canvas graphs should draw the background separately, r=pbrosset",
		"Bug 1343: Support asm.js frames in SavedStacks. r=luke",
	}

	for _, str := range invalid {
		if IsValidCommit(str) {
			t.Error("Expected " + str + " to be an invalid commit.")
		}
	}

	for _, str := range valid {
		if !IsValidCommit(str) {
			t.Error("Expected " + str + " to be a valid commit.")
		}
	}
}

func TestCreateMessage(t *testing.T) {
	messages := [][]string{
		[]string{
			"Bug 1034670 - The canvas graphs should draw the background separately, r=pbrosset",
			"The canvas graphs should draw the background separately http://bugzil.la/1034670",
		},
		[]string{
			"Bug 1034670 - The canvas graphs should draw the background separately, a=pbrosset",
			"The canvas graphs should draw the background separately http://bugzil.la/1034670",
		},
		[]string{
			"Bug 1343: Support asm.js frames in SavedStacks. r=luke ",
			"Support asm.js frames in SavedStacks http://bugzil.la/1343",
		},
		[]string{
			"bug 1003450: Make all the cold trees mourn Their branches frozen in sightless motion Waving, reaching for the whipping rain There was silence And the firmament withdrew Revealing all, shapelessly and swiftly",
			"Make all the cold trees mourn Their branches frozen in sightless motion Waving, reaching for the whipping rain T... http://bugzil.la/1003450",
		},
		[]string{
			" Bug 1034668-The `getMappedSelection` method for all canvas graphs should clamp the selection bounds, r=pbrosset,someoneelse, space",
			"The `getMappedSelection` method for all canvas graphs should clamp the selection bounds http://bugzil.la/1034668",
		},
		[]string{
			`Bug 988314 - Rename Inspector tests and supporting documents. r=pbrosset

			--HG--
			rename : browser/devtools/inspector/test/browser_inspector_bug_817558_delete_node.js => browser/devtools/inspector/test/browser_inspector_delete-selected-node-01.js
			rename : browser/devtools/inspector/test/browser_inspector_bug_848731_reset_selection_on_delete.js => browser/devtools/inspector/test/browser_inspector_delete-selected-node-02.js
			rename : browser/devtools/inspector/test/browser_inspector_destroyselection.js => browser/devtools/inspector/test/browser_inspector_delete-selected-node-03.js
			rename : browser/devtools/inspector/test/browser_inspector_bug_840156_destroy_after_navigation.js => browser/devtools/inspector/test/browser_inspector_destroy-after`,
			"Rename Inspector tests and supporting documents http://bugzil.la/988314",
		},
	}

	for _, duple := range messages {
		if CreateMessage(duple[0], GetBugNumber(duple[0])) != duple[1] {
			t.Error("Expected '" + CreateMessage(duple[0], GetBugNumber(duple[0])) + "' to equal '" + duple[1] + "'")
		}
	}
}

func TestGetBugNumber(t *testing.T) {
	messages := [][]string{
		[]string{
			"Bug 1034670 - The canvas graphs should draw the background separately, r=pbrosset",
			"1034670",
		},
		[]string{
			"Bug 1034670 - The canvas graphs should draw the background separately, a=pbrosset",
			"1034670",
		},
		[]string{
			"Bug 1343: Support asm.js frames in SavedStacks. r=luke ",
			"1343",
		},
		[]string{
			" Bug 1034668-The `getMappedSelection` method for all canvas graphs should clamp the selection bounds, r=pbrosset,someoneelse, space",
			"1034668",
		},
		[]string{
			`Bug 988314 - Rename Inspector tests and supporting documents. r=pbrosset

			--HG--
			rename : browser/devtools/inspector/test/browser_inspector_bug_817558_delete_node.js => browser/devtools/inspector/test/browser_inspector_delete-selected-node-01.js
			rename : browser/devtools/inspector/test/browser_inspector_bug_848731_reset_selection_on_delete.js => browser/devtools/inspector/test/browser_inspector_delete-selected-node-02.js
			rename : browser/devtools/inspector/test/browser_inspector_destroyselection.js => browser/devtools/inspector/test/browser_inspector_delete-selected-node-03.js
			rename : browser/devtools/inspector/test/browser_inspector_bug_840156_destroy_after_navigation.js => browser/devtools/inspector/test/browser_inspector_destroy-after`,
			"988314",
		},
	}

	for _, duple := range messages {
		if GetBugNumber(duple[0]) != duple[1] {
			t.Error("Expected '" + duple[0] + "' to clean into '" + duple[1] + "'")
		}
	}
}

func TestGetBugzillaURL(t *testing.T) {
	messages := [][]string{
		[]string{
			"1034670",
			"http://bugzil.la/1034670",
		},
		[]string{
			"1343",
			"http://bugzil.la/1343",
		},
		[]string{
			"",
			"",
		},
	}

	for _, duple := range messages {
		if GetBugzillaURL(duple[0]) != duple[1] {
			t.Error("Expected '" + duple[0] + "' to have a URL: " + duple[1])
		}
	}
}
