package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Declare a subcommand using NewFlagSet and assign flags to it
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// Define a new flag set for a different subcommand
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// Expect at least 1 subcommand to be used
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// Check which subcommand is invoked
	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Fprintln(os.Stderr, "expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
	// Try commands:
	// go run cli_subc.go foo -enable --name=Joe
	// go run cli_subc.go bar -level 8 a1
	// go run cli_subc.go baz
}
