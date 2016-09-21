// token.go
package main

// Token represents a lexical token.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota

	EOF
	WS
	IDENT    // [a-z][A-Z]
	DIGIT    // [0-9]
	COLON    // :
	OBRACE   // {
	EBRACE   // }
	OBRACKET // [
	EBRACKET // ]

	// Special Keyeord
	MESSAGE
	STRUCT
)

var eof = rune(0)
