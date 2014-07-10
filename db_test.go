package main

import "testing"
import "os"

var TEST_DB_NAME string = "test_db.txt"
var SHAS []string = []string{
	"873dde64d69092a86d254e0892822aaaf4065422",
	"082a3d0875558bd25acbab62de69940d8bdd1158",
	"51c58b7891ae0dd16a71f0a3b1d5c07ec2e663dc",
	"b50c5fb71ce712c22d9bdadc3c21122d7d01c46c",
}

func TestAddSHA(t *testing.T) {
	os.Remove(TEST_DB_NAME)

	for _, sha := range SHAS {
		AddSHA(TEST_DB_NAME, sha)
	}

	for _, sha := range SHAS {
		if GetSHA(TEST_DB_NAME, sha) == false {
			t.Error("Expected " + sha + " to be in " + TEST_DB_NAME)
		}
	}

	os.Remove(TEST_DB_NAME)
}

func TestGetSHA(t *testing.T) {
	os.Remove(TEST_DB_NAME)

	for _, sha := range SHAS {
		AddSHA(TEST_DB_NAME, sha)
	}

	if GetSHA(TEST_DB_NAME, "51c58b7891ae0dd16a71f0a3b1d5c07ec2e663dc") != true {
		t.Error("Expected SHA to be in " + TEST_DB_NAME)
	}

	if GetSHA(TEST_DB_NAME, "51c58b789") == true {
		t.Error("Expected subset of valid SHA to not return true.")
	}

	if GetSHA(TEST_DB_NAME, "0688d7437161f7832a1ed0779cfe8b604f61533e") == true {
		t.Error("Expected SHA to not be found in " + TEST_DB_NAME)
	}

	os.Remove(TEST_DB_NAME)
}
