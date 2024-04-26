package ascii

import (
	"fmt"
)
// This function gets to receive 2 args both of a string array data type.
// It first checks if the word is an empty string, if it is, a new line 
// is printed otherwise it goes ahead and loops through 
// every character in the word string and prints
// the individual characters with their matching ASCII art.
// Afterwards a new line is printed.

func AsciiArt(words []string, contents2 []string){
	for _, word := range words {
		if word != "" {	
				for i := 0; i < 8; i++ {
					for _, char := range word {
						fmt.Print(contents2[int(char-' ')*9+1+i])
					}
				fmt.Println()
			
				} 
			} else {
			fmt.Println()
		}

	}	
}