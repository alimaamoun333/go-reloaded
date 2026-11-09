package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage : go run . input.txt output.txt")
		os.Exit(1)
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read the input file
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file : ", err)
		os.Exit(1)
	}
	// Process the content (for demonstration, we just convert it to uppercase)
	text := string(content)
	text = transformText(text)
	text = fixPunctuation(text)
	text = fixQuotes(text)
	text = fixArticles(text)

	// Write the processed content to the output file
	err = os.WriteFile(outputFile, []byte(text), 0644)
	if err != nil {
		fmt.Println(" Error writing to output file : ", err)
		os.Exit(1)
	}
}
