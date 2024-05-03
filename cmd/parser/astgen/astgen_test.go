package astgen

import (
	"MonarchServer/cmd/parser/tokenizer"
	"testing"
)

func TestString(t *testing.T) {
	ast := &AST{
		Nodes: []Node{
            &Prefix{
                Token: tokenizer.Token{Type: tokenizer.WORD, Literal: "TestSource"},
                Value: "TestSource",
            },
            &Command{
                Token: tokenizer.Token{Type: tokenizer.WORD, Literal: "PASS"},
                Value: "PASS",
                Params: []string{"supersecretpassword"},
            },
		},
	}

	if ast.String() != "TestSource PASS supersecretpassword" {
		t.Errorf("program.String() wrong. got=%q", ast.String())
	}
}
