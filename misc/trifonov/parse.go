package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

// IgnoredConfig defines arguments which can be accepted for compatibility with rsync,
// but are ignored by exodus-rsync.
type IgnoredConfig struct {
	Recursive bool `short:"r"`
	Delete    bool
	Rsh       string `short:"e"`
	Timeout   int
}

func main() {
	out := IgnoredConfig{}
	kong.Parse(&out)
	if out.Recursive {
		fmt.Println("Recursive Tom")
	}
}

// "string" is an interpreted string literal. It can't contain newlines.
// `string` is a raw string literal.

// backticks/backquotes in structs definition is "tags"
// The tags are made visible through a reflection interface and take part
// in type identity for structs but are otherwise ignored.

/*
Reference:

Tag:
https://stackoverflow.com/questions/30681054/what-is-the-usage-of-backtick-in-golang-structs-definition
*/
