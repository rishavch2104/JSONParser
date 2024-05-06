package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	filePath := flag.Args()[0]
	fmt.Print(validateJson(filePath))
}

func validateJson(path string) int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 1
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if isEmptyFile(scanner) {
		return 2
	}

	stack := Stack{}
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			if isOpeningBracket(char) {
				stack.Push(char)
			} else {
				if isClosingBracket(char) {
					stack.Pop()
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return 1
	}
	if stack.IsEmpty() {
		return 0
	}
	return 2

}

func isOpeningBracket(char rune) bool {
	return char == '{'
}

func isClosingBracket(char rune) bool {
	return char == '}'
}

func isEmptyFile(scanner *bufio.Scanner) bool {
	hasContent := scanner.Scan()
	return !hasContent
}
