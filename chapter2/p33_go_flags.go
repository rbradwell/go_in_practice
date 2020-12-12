package main

import (
	flags "github.com/jessevdk/go-flags" // import go-flags aliased to the name flags
)

var opts struct {
	Name    string `short:"n" long:"name" default:"World" description:"A name to say hello to"`
	Spanish bool   `short:"s" long:"spanish" description:"Use spanish Language"`
}

func main() {
	flags.Parse(&opts)
}
