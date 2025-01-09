package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) > 1 {
		text := os.Args[1]

		branch, err := shellCommand("git", "branch")
		if err != nil {
			panic(err)
		}
		diff, err := shellCommand("git", "diff")
		if err != nil {
			panic(err)
		}
		fmt.Printf("branch: %v\n", string(branch))
		fmt.Printf("diff: %v\n", string(diff))
		fmt.Printf("text: %v\n", text)

	}
}

func shellCommand(command ...string) (string, error) {
	cmd := exec.Command(command[0], command[1:]...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
