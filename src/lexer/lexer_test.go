package lexer

import "testing"

func TestNextToken(t *testing.T) {

	tests := []string{
		"courses?filter=equals(displayName,'Brian O''Connor')",
		"students?filter=equals(displayName,null)",
		"teachers?filter=equals(displayName,lastName)",
	}

	for _, test := range tests {
		l := New(test)

		tok := l.NextToken()
		tok.Print()
		if tok.Type == Eof {
			break
		}
	}
}
