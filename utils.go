package main

const MAX_LENGTH = 100

func FormatMessage(s string) string {
	if len(s) < MAX_LENGTH {
		return s
	} else {
		return s[0:MAX_LENGTH] + "..."
	}
}
