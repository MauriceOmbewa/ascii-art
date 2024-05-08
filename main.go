package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"ascii/ascii"
)

func main() {
	//Dealing with the length of arguments error.
	if len(os.Args) < 2 {
		fmt.Println("Arguments less than required")
	} else if len(os.Args) > 2 {
		fmt.Println("Arguments more than required")
	}

	arguments := os.Args[1]

	if arguments == "" { //if arguments is empty, nothing is printed and the program just returns
		return
	} else if arguments == "\\n" || arguments == "\\t" { // newline is printed in the event of "\\n" or "\\t" as the arguments
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


	words := strings.Split(arguments, "\\n") // splitting the arguments every instance of \n

	for i := 0; i < len(arguments); i++ {
		if arguments[i] < 32 || arguments[i] > 126 { //dealing with anything outside the scope of ascii characters required
			fmt.Println("Input is not within the range of ascii characters wanted")
			return
		}
	}
	ascii.AsciiArt(words, lines)
}
