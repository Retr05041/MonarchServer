package parser

import (
	"MonarchServer/cmd/parser/lexer"
	"testing"
)

// Test a "let"
func TestParser(t *testing.T) {
	input := `:TestSource PASS supersecret`
	l := lexer.New(input)
	p := New(l)

	program := p.Parse()

	if program == nil {
		t.Fatalf("Parse() returned nil")
	}
	if len(program.Nodes) != 4 {
		t.Fatalf("program.Statments does not contain 4 statments. got=%d", len(program.Nodes))
	}

	if program.Nodes[0].TokenLiteral() != ":" {
		t.Fatalf("program.Node[0] is not ':'. got=%s", program.Nodes[0].TokenLiteral())
	}
	if program.Nodes[1].TokenLiteral() != "TestSource" {
		t.Fatalf("program.Node[1] is not 'TestSource'. got=%s", program.Nodes[0].TokenLiteral())
	}
	if program.Nodes[2].TokenLiteral() != "PASS" {
		t.Fatalf("program.Node[2] is not 'PASS'. got=%s", program.Nodes[0].TokenLiteral())
	}
	if program.Nodes[3].TokenLiteral() != "supersecret" {
		t.Fatalf("program.Node[3] is not 'supersecret'. got=%s", program.Nodes[0].TokenLiteral())
	}
}
