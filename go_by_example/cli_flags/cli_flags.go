package main

import (
	"flag"
	"fmt"
	"time"
)

// Flags is a common way to specify options for a program like "-l" or "-lm"
func main() {
	// Flag is a convenient package to support flags
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")
	// Duration should receive a param acceptable by time.ParseDuration
	durPtr := flag.Duration("dur", 3 * time.Second, "a time duration")
	// We declared params, arguments given are: name, value and a short description
	// Also, we can declare a flag using an existing var. Note, that pointer is passed to function
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// After all declarations, use Parse to get execute cli parsing
	flag.Parse()

	// Same flag given twice overrides previous one
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("fork:", *durPtr)
	fmt.Println("svar:", svar)
	// This also dumps all tailing positional arguments
	fmt.Println("tail:", flag.Args())
	// Commands to try:
	// go run cli_flags.go -word=opt -numb=7 -fork -svar=flag
	// go run cli_flags.go -word=opt
	// go run cli_flags.go -word=opt a1 a2 a3
	// go run cli_flags.go -word=opt a1 a2 a3 -numb=7
	// go run cli_flags.go --dur=1.5h
	// Note: all flags should come before positional arguments. Otherwise, they'll be parsed as regular positional args
	// Use -h or -help flag to get autogenerated flags descriptions!
	// go run cli_flags.go -h
	// If non-declared flag is given, program shows error and help message
	// go run cli_flags.go -wat
	// One and 2 minus signs are equivalent!
	// go run cli_flags.go --word=opt
	// Also, "=" sign is optional for non-boolean flags
	// go run cli_flags.go --word opt
	// Boolean flags: 1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False
	// go run cli_flags.go --fork=T
}
