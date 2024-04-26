package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"ascii/ascii"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Arguments less than required")
	} else if len(os.Args) > 2 {
		fmt.Println("Arguments more than required")
	}

	arguments := os.Args[1]

	if arguments == "" {
		return
	} else if arguments == "\\n" || arguments == "\\t" {
		fmt.Print("\n")
		return
	}

	inputfile, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Error while opening file")
	}

	scaner := bufio.NewScanner(inputfile)

	var lines []string

	for scaner.Scan() {
		lines = append(lines, scaner.Text())
	}

	// scanner := strings.Split(string(inputfile), "\n")

	words := strings.Split(arguments, "\\n")

	for i := 0; i < len(arguments); i++ {
		if arguments[i] < 32 || arguments[i] > 126 {
			fmt.Println("Input is not within the range of ascii characters wanted")
			return
		}
	}
	ascii.AsciiArt(words, lines)
}
