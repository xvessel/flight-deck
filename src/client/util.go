package main

import (
	"fmt"
	"os"
)
func ExitIfError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}
}

func Yes() bool {
	fmt.Printf("? Y/[any other key cancel]:")
	b := make([]byte, 1)
	_, err := os.Stdin.Read(b)
	ExitIfError(err)
	if b[0] == 'Y' {
		return true
	} else {
		return false
	}
}


