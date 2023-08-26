package main

import (
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		str := scanner.Text()

		tokens := Tokenize(str)
		root := parse(tokens)
		result := exec(root)

		println(result.Value)
	}
}
