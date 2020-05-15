package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// Go requires absolute path to binary we want to exec
	// We can use LookPath for that
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Exec requires arguments in a slice form
	// First arg should be a program name
	args := []string{"ls", "-a", "-l", "-h"}

	// Exec also needs a set of env vars to use. Here we are providing our current env
	env := os.Environ()

	// Actual syscall. If this call is ok, the execution of our process will end here and be
	// Replaced by the /bin/ls -a -l -h process
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
	// Go does NOT provide the classic fork function as starting goroutines and executing
	// processes covers most of the cases
}
