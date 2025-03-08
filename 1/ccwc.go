package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"unicode/utf8"
	"strings"
)

type countStats struct {
	lines int
	words int
	bytes int
	chars int
}

func countFromReader(r io.Reader) countStats {
	scanner := bufio.NewScanner(r)
	stats := countStats{}

	for scanner.Scan() {
		line := scanner.Text()
		stats.lines++
		stats.words += len(strings.Fields(line)) // splitting line into words
		stats.bytes += len(line) + 1 // +1 for newline character
		stats.chars += utf8.RuneCountInString(line) + 1 // +1 for newline character
	}

	return stats
}

func main() {
	args := os.Args

	var flag string
	var fileName string
	var reader io.Reader

	// Determine input soruce
	if (len(args) == 1) {
		flag = "default"
		reader = os.Stdin
	} else if (len(args) == 2) {
		// only flag provided, no filename
		if (strings.HasPrefix(args[1], "-")) {
			flag = args[1]
			reader = os.Stdin
		} else {
			flag = "default"
			fileName = args[1]
		}
	} else {
		flag = args[1]
		fileName = args[2]
	}

	if fileName != "" {
		// open file if filename is provided
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening file: " + fileName)
			os.Exit(1)
		}
	
		defer file.Close()
		reader = file
	} else {
		// read from standard input
		reader = os.Stdin
	}	

	// Count words, lines and bytes
	stats := countFromReader(reader)

	switch flag {
	case "-c":
		fmt.Printf("%d %s", stats.bytes, fileName)
	case "-l":
		fmt.Printf("%d %s", stats.lines, fileName)
	case "-w":
		fmt.Printf("%d %s", stats.words, fileName)
	case "-m":
		fmt.Printf("%d %s", stats.chars, fileName)
	default:
		fmt.Printf("%d %d %d %s", stats.lines, stats.words, stats.bytes, fileName)
	}
	fmt.Println()
}