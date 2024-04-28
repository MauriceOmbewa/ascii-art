### ASCII Art Generator

This simple ASCII art generator program takes a string input and prints out the corresponding ASCII art. It also allows for the input of special characters such as newline (\n) to produce formatted output.
## Usage

    __Installation__: Clone the repository to your local machine.

    **Run the Program**: Navigate to the directory containing the program and execute it.

bash

``` go run main.go <input_string> ```

Replace <input_string> with the string you want to convert to ASCII art.
# Example

bash

```go run main.go "HELLO\nWORLD"```

## Input Restrictions

    The input string should consist of ASCII characters within the range of 32 to 126.
    Special characters such as newline (\n) can be used for formatting.

## Implementation Details

The ASCII art generation is handled by the AsciiArt function in the ascii package. This function takes two arguments: an array of words and an array containing ASCII art representations for each character.

### ASCII Package Documentation
## Function: AsciiArt

go

```func AsciiArt(words []string, contents2 []string)```

This function generates ASCII art for the provided words using the ASCII art contents provided. It prints each word's ASCII representation, with each character represented in its corresponding ASCII art form.

    words: A string array containing the words to be converted to ASCII art.
    contents2: A string array containing the ASCII art representation for each character.