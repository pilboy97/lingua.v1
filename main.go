package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	buf, err := os.ReadFile("test.txt")
	if err != nil {
		log.Panic(err)
	}

	str := addSColon(string(buf))
	tokens := Tokenize(str)

	root := parse(tokens)
	for _, ch := range root.Child {
		fmt.Printf("%v\n", exec(ch))
	}
}
