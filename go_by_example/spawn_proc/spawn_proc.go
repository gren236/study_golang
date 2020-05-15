package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	// Start with a simple command with no args
	// Command creates an object representing a command
	dateCmd := exec.Command("date")

	// Output runs the command, waits for it to finish and collects the output as byte array
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// Example where we pipe data to process stdin and collect results from stdout
	grepCmd := exec.Command("grep", "hello")

	// Here we grab in/out pipes and start the process
	// Then write some input to it, read the result and wait for the process to exit
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()

	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	// We only collect StdoutPipe results, but we could collect StderrPipe as well
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// When we are spawning commands we need to provide command and argument array
	// To spawn a full command with a string bash's -c can be used
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
