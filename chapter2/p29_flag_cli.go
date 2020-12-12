package main

/*
// create a new variable from a flag
var name = flag.String("name", "World", "A name to say hello to.")

var spanish bool // new variable to store flag value

func init() {
	// set variable to the flag value
	flag.BoolVar(&spanish, "spanish", false, "Use spanish language") // used for long name flag i.e -spanish
	flag.BoolVar(&spanish, "s", false, "Use spanish language")       // used for short name flag i.e. -s
	// flag functions exist for each variable type see the flag package (http://golang.org/pkg/flag)
}

func main() {
	flag.Parse() // parse the flags, placing values into variables
	if spanish == true {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
	// visit all the flags and print the help etc e.g.
	// Hola rich!
	// -name: A name to say hello to. (Default: 'World')
	// -s: Use spanish language (Default: 'false')
	// -spanish: Use spanish language (Default: 'false')
	flag.VisitAll(func(flag *flag.Flag) {
		format := "\t-%s: %s (Default: '%s')\n"
		fmt.Printf(format, flag.Name, flag.Usage, flag.DefValue)
	})
}
*/
// go run p29_flag_cli.go -s -name rich
// Hola rich!

// note the above are not linux style flags, with the main issue you can't mix short flags
// e.g. ls -la
