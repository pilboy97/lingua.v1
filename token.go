package main

const (
	EOF = iota
	NUM
	NAME
	TRUE
	FALSE
	LAND
	LOR
	ASS
	ADD
	SUB
	INC
	DEC
	MUL
	DIV
	MOD
	LSH
	RSH
	BAND
	BOR
	BXOR
	EQ
	NEQ
	GR
	GRE
	LE
	LEQ
	OB
	CB
	OSB
	CSB
	OBL
	CBL
	COLON
	SCOLON
	COMMA
)

var tokenDict = map[string]int{
	"true":  TRUE,
	"false": FALSE,
	"and":   LAND,
	"or":    LOR,
	"=":     ASS,
	"+":     ADD,
	"-":     SUB,
	"++":    INC,
	"--":    DEC,
	"*":     MUL,
	"/":     DIV,
	"%":     MOD,
	"<<":    LSH,
	">>":    RSH,
	"&":     BAND,
	"|":     BOR,
	"^":     BXOR,
	"==":    EQ,
	"!=":    NEQ,
	">":     GR,
	">=":    GRE,
	"<":     LE,
	"<=":    LEQ,
	"(":     OB,
	")":     CB,
	"[":     OSB,
	"]":     CSB,
	"{":     OBL,
	"}":     CBL,
	":":     COLON,
	";":     SCOLON,
	",":     COMMA,
}

type Token struct {
	Type int
	Str  string
}
