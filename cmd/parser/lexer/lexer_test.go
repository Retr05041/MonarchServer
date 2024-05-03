package lexer

import (
    "testing"
    "MonarchServer/cmd/parser/tokenizer"
)

func TestReadToken(t *testing.T) {
	// Get an arbitrary input from the user
    input := `:TestSource PASS mysupersecretpassword`


	// Create an anonymous struct (exact same as token struct)
	tests := []struct {
		expectedType    tokenizer.TokenType
		expectedLiteral string
	}{
		// Expected Type and Literal to test against from input
        {tokenizer.COLON, ":"},
		{tokenizer.PREFIX, "TestSource"},
		{tokenizer.CMD, "PASS"},
		{tokenizer.PARAMETER, "mysupersecretpassword"},
	}

	l := New(input)

	// Loop through and test each token i = iteration, tt = given test
	for i, tt := range tests {
		tok := l.ReadToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
