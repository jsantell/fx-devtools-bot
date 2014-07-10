package main

import "testing"
import "os"

var filename string = "test_db.txt"
var shas []string = []string{
	"873dde64d69092a86d254e0892822aaaf4065422",
	"082a3d0875558bd25acbab62de69940d8bdd1158",
	"51c58b7891ae0dd16a71f0a3b1d5c07ec2e663dc",
	"b50c5fb71ce712c22d9bdadc3c21122d7d01c46c",
}

func TestAddSHA(t *testing.T) {
	os.Remove(filename)

	for _, sha := range shas {
		AddSHA(filename, sha)
	}

	for _, sha := range shas {
		if GetSHA(filename, sha) == false {
			t.Error("Expected " + sha + " to be in " + filename)
		}
	}

	os.Remove(filename)
}

func TestGetSHA(t *testing.T) {
	os.Remove(filename)

	for _, sha := range shas {
		AddSHA(filename, sha)
	}

	if GetSHA(filename, "51c58b7891ae0dd16a71f0a3b1d5c07ec2e663dc") != true {
		t.Error("Expected SHA to be in " + filename)
	}

	if GetSHA(filename, "51c58b789") == true {
		t.Error("Expected subset of valid SHA to not return true.")
	}

	if GetSHA(filename, "0688d7437161f7832a1ed0779cfe8b604f61533e") == true {
		t.Error("Expected SHA to not be found in " + filename)
	}

	os.Remove(filename)
}
