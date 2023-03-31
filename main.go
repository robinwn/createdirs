package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// get the root directory from the command-line arguments
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: createdirs <root_dir>")
		return
	}
	rootDir := args[1]

	// get the permissions of the root directory
	rootInfo, err := os.Stat(rootDir)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	perms := rootInfo.Mode().Perm()

	// create a scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// loop through each line of input
	for scanner.Scan() {
		// get the directory path from the input line
		dirPath := scanner.Text()

		// create the directory under the root directory with inherited permissions
		fullPath := filepath.Join(rootDir, dirPath)
		err := os.MkdirAll(fullPath, perms)
		if err != nil {
			// if there was an error creating the directory, print an error message
			fmt.Printf("Error creating directory '%s': %s\n", fullPath, err)
		} else {
			// if the directory was created successfully, print a success message
			// fmt.Printf("Directory '%s' created successfully\n", fullPath)
		}
	}

	// check for any errors that occurred while scanning input
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading standard input:", err)
	}
}
