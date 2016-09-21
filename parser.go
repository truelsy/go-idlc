// Parser.go
package main

import (
	"fmt"
	"io"
	"strconv"
)

type Parser struct {
	s   *Scanner
	buf struct {
		tok Token
		lit string
		n   int
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) scan() (Token, string) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	p.buf.tok, p.buf.lit = p.s.Scan()
	return p.buf.tok, p.buf.lit
}

func (p *Parser) unscan() {
	p.buf.n = 1
}

func (p *Parser) scanIgnoreWhiteSpace() (Token, string) {
	tok, lit := p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return tok, lit
}

func (p *Parser) Parse() (*TokenStmt, error) {
	tokenStmt := &TokenStmt{}

	var isEof bool = false

	for {
		tokenElement := &TokenElement{}

		tok, lit := p.scanIgnoreWhiteSpace()
		if tok == MESSAGE {
			tokenElement.IsMsg = true
		} else if tok == STRUCT {
			// do nothing
		} else {
			return nil, fmt.Errorf("found %q, expected MESSAGE or STRUCT", lit)
		}

		// Parse Name
		tok, lit = p.scanIgnoreWhiteSpace()
		if tok != IDENT {
			return nil, fmt.Errorf("found %q, expected message name", lit)
		}
		tokenElement.Name = lit

		if tokenElement.IsMsg {
			if tok, lit := p.scanIgnoreWhiteSpace(); tok != COLON {
				return nil, fmt.Errorf("found %q, expected COLON", lit)
			}

			// Parse Message ID
			tok, lit = p.scanIgnoreWhiteSpace()
			if tok != DIGIT {
				return nil, fmt.Errorf("found %q, expected message id", lit)
			}
			tokenElement.Id, _ = strconv.Atoi(lit)
		}

		// Parse Message Member
		for {
			// Read a field.
			tok, lit := p.scanIgnoreWhiteSpace()
			//if tok == EOF {
			//	isEof = true
			//	break
			//}

			if tok == EBRACE {
				tok, _ = p.scanIgnoreWhiteSpace()
				if tok == EOF {
					isEof = true
				} else {
					p.unscan()
				}
				break
			}

			if tok == IDENT {

				//a := []rune(lit)
				//a[0] = unicode.ToUpper(a[0])
				//lit = string(a)

				// 변수명
				tokenElement.VarName = append(tokenElement.VarName, lit)

				tokNext, litNext := p.scanIgnoreWhiteSpace()
				if tokNext == IDENT {
					// 변수 타입
					tokenElement.VarType = append(tokenElement.VarType, litNext)
				} else {
					return nil, fmt.Errorf("found %q, expected ident", litNext)
				}

			}

			//			switch tok {
			//			case IDENT:
			//				tokenElement.VarName = append(tokenElement.VarName, lit)
			//			case VARTYPE:
			//				tokenElement.VarType = append(tokenElement.VarType, lit)
			//			}
		}

		tokenStmt.Elements = append(tokenStmt.Elements, tokenElement)

		if isEof {
			break
		}
	}

	return tokenStmt, nil
}
