package main

import (
	"log"
)

type Node struct {
	Str   []Token
	Par   *Node
	Child []*Node
}

var tape []Token
var header int

func parse(tokens []Token) *Node {
	tape = tokens
	header = 0

	var ret Node

	var st = header

	var init = parseEXPR()
	init.Par = &ret
	ret.Child = append(ret.Child, init)

	for header < len(tape) {
		if tape[header].Type != SCOLON {
			log.Panicf("missing ';'")
		}
		header++

		if header < len(tape) && tape[header].Type == SCOLON {
			continue
		}
		if header >= len(tape) {
			break
		}

		var ch = parseEXPR()
		ch.Par = &ret
		ret.Child = append(ret.Child, ch)
	}

	ret.Str = tape[st:header]

	return &ret
}
func parseEXPR() *Node {
	var ret Node

	var st = header

	var init = parseMUL()
	init.Par = &ret
	ret.Child = append(ret.Child, init)

	for header < len(tape) {
		if tape[header].Type != ADD && tape[header].Type != SUB {
			break
		}
		ret.Child = append(ret.Child, &Node{
			Str: []Token{tape[header]},
		})
		header++

		var ch = parseMUL()
		ch.Par = &ret
		ret.Child = append(ret.Child, ch)
	}

	ret.Str = tape[st:header]

	return &ret
}
func parseMUL() *Node {
	var ret Node

	var st = header

	var init = parseFACTOR()
	init.Par = &ret
	ret.Child = append(ret.Child, init)

	for header < len(tape) {
		if tape[header].Type != MUL && tape[header].Type != DIV {
			break
		}
		ret.Child = append(ret.Child, &Node{
			Str: []Token{tape[header]},
		})
		header++

		var ch = parseEXPR()
		ch.Par = &ret
		ret.Child = append(ret.Child, ch)
	}

	ret.Str = tape[st:header]

	return &ret
}
func parseFACTOR() *Node {
	var ret Node

	var st = header

	if tape[header].Type == OB {
		header++
		var init = parseEXPR()
		init.Par = &ret
		ret.Child = append(ret.Child, init)

		if tape[header].Type != CB {
			log.Panicf("missing ')'")
		}
		header++
	} else {
		if tape[header].Type != NUM {
			log.Panicf(
				"unexpected %s",
				tape[header].Str,
			)
		}

		ret.Str = []Token{tape[header]}
		header++
	}

	ret.Str = tape[st:header]

	return &ret
}
