package main

type Commit interface {
	FormatMessage() string
	IsValid() bool
}
