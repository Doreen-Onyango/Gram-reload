# Go Reloaded

![image](https://miro.medium.com/v2/resize:fit:4800/format:webp/1*IPQm7JhoY5v7ni1YAkSocA.jpeg)

## The Overview
This is a simple text processing tool. it is written in Go language. It performs functionality to perform modifications on text files depending on the rule.

## Features
- **(cap):** Capitalizes the letter at the first position of the word before (cap) and removes the (cap) from the sentence.
- **(up):** Changes to uppercase the word that appears before (up) and removes the (up) from the sentence.
- **(low):** Changes to lowercase the word that appears before (low) and removes (low) from the sentence.
- **(hex)/(bin):** Changes hexadecimal value and binary value to decimal value when it appears before (hex) and (bin) respectively and removes the (hex) and (bin)
- **a/an:** Checks for instance of 'a' and if the next word starts with a vowel or h then replaces it with 'an'. Uppercase and lowercase is checked.
- **punctuations/quotes:** Checks from any instances of a wrongly used punctuations and quotes and rectifies them grammerly.

## Usage
Run the main program by passing two command-line arguments, input file path and output file path:

``` bash
go run . sample.txt result.txt
```
Replace sample.txt file with the file you want to test and result.txt with the output file.

## Testing
This project includes tests that checks the functionality of the programs to ensure they are working well.
To execute the test files:

```bash
go test -v
```




