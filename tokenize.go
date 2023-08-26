package main

import (
	"log"
)

func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}
func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}
func isSpace(b byte) bool {
	return b == ' ' || b == '\n' || b == '\t'
}
func searchDict(str string) bool {
	_, ok := tokenDict[str]
	return ok
}
func addSColon(str string) string {
	bytes := []byte(str)
	place := make([]int, 0)

	bytes = append(bytes, ';')

	for i := 0; i < len(bytes); i++ {
		if bytes[i] == '\n' {
			if i > 0 && !(bytes[i] == '(' || bytes[i] == '[' || bytes[i] == '{' || bytes[i] == ',') {
				place = append(place, i)
			}
		}
	}

	for i := len(place) - 1; i >= 0; i-- {
		v := place[i]
		if v+1 < len(bytes) {
			bytes = append(bytes[:v+1], bytes[v:]...)
			bytes[v] = ';'
		}
	}

	return string(bytes)
}
func Tokenize(str string) []Token {
	res := []Token{}
	bytes := []byte(str)

	r, c := 1, 0
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == '\n' {
			r++
		}
		switch {
		case isSpace(bytes[i]):
			for i < len(bytes) && isSpace(bytes[i]) {
				i++
				c++
			}
			i--
		case isNumber(bytes[i]):
			st := i
			for i < len(bytes) && isNumber(bytes[i]) {
				i++
				c++
			}

			res = append(res, Token{
				Type: NUM,
				Str:  string(bytes[st:i]),
			})
			i--
		case isAlpha(bytes[i]) || bytes[i] == '_':
			st := i
			i++
			c++

			for i < len(bytes) && (bytes[i] == '_' || isAlpha(bytes[i]) || isNumber(bytes[i])) {
				i++
				c++
			}

			str := string(bytes[st:i])

			if searchDict(str) {
				res = append(res, Token{
					Type: tokenDict[str],
					Str:  str,
				})
			} else {
				res = append(res, Token{
					Type: NAME,
					Str:  str,
				})
			}
			i--
		case i+1 < len(bytes) && searchDict(string(bytes[i:i+2])):
			res = append(res, Token{
				Type: tokenDict[string(bytes[i:i+2])],
				Str:  string(bytes[i : i+2]),
			})
			c++
			i++
		case searchDict(string(bytes[i : i+1])):
			res = append(res, Token{
				Type: tokenDict[string(bytes[i:i+1])],
				Str:  string(bytes[i : i+1]),
			})
		default:
			log.Panic("unknown character")
		}
	}

	return res
}
