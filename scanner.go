// Scanner.go
package main

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

func (s *Scanner) scanWhiteSpace() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

func (s *Scanner) scanIdent() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// Special Keyword
	switch strings.ToUpper(buf.String()) {
	case "MESSAGE":
		return MESSAGE, strings.ToUpper(buf.String())
	case "STRUCT":
		return STRUCT, buf.String()
	}

	return IDENT, buf.String()
}

func (s *Scanner) scanArray() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' && ch != '[' && ch != ']' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return IDENT, buf.String()
}

func (s *Scanner) scanDigit() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isDigit(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return DIGIT, buf.String()
}

func (s *Scanner) Scan() (Token, string) {
	ch := s.read()

	if isEof(ch) {
		return EOF, ""
	} else if isWhitespace(ch) {
		s.unread()
		return s.scanWhiteSpace()
	} else if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	} else if isDigit(ch) {
		s.unread()
		return s.scanDigit()
	} else if isOpenBrace(ch) {
		return OBRACE, string(ch)
	} else if isEndBrace(ch) {
		return EBRACE, string(ch)
	} else if isColon(ch) {
		return COLON, string(ch)
	} else if isOpenBracket(ch) {
		s.unread()
		return s.scanArray()
	}

	return ILLEGAL, string(ch)
}
