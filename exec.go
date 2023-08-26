package main

import (
	"log"
	"strconv"
)

type Result struct {
	Type  int
	Value int
}

const (
	OPER = iota
	NUMBER
)

func exec(root *Node) Result {
	if len(root.Child) == 0 {
		if len(root.Str) != 1 {
			log.Panic("did not parse right")
		}

		switch root.Str[0].Type {
		case ADD:
			return Result{
				Type:  OPER,
				Value: ADD,
			}
		case SUB:
			return Result{
				Type:  OPER,
				Value: SUB,
			}
		case MUL:
			return Result{
				Type:  OPER,
				Value: MUL,
			}
		case DIV:
			return Result{
				Type:  OPER,
				Value: DIV,
			}
		case OB:
			return Result{
				Type:  OPER,
				Value: OB,
			}
		case CB:
			return Result{
				Type:  OPER,
				Value: CB,
			}
		case NUM:
			n, err := strconv.ParseInt(root.Str[0].Str, 10, 32)
			if err != nil {
				log.Panic(err)
			}

			return Result{
				Type:  NUMBER,
				Value: int(n),
			}
		default:
			log.Panicf("unexpected %s", root.Str[0].Str)
		}
	}

	var ret Result = exec(root.Child[0])
	var mode int

	if ret.Type == OPER && ret.Value == OB {
		return exec(root.Child[1])
	}

	for _, ch := range root.Child[1:] {
		if res := exec(ch); res.Type == OPER {
			mode = res.Value
		} else {
			switch mode {
			case ADD:
				ret.Value += res.Value
			case SUB:
				ret.Value -= res.Value
			case MUL:
				ret.Value *= res.Value
			case DIV:
				ret.Value /= res.Value
			}
		}
	}

	return ret
}
