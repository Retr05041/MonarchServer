package lexer

import "MonarchServer/cmd/parser/tokenizer"

type Lexer struct {
	input         string
	currCharIndex int
	nextCharIndex int
	char          byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.slideRight()
	return l
}

func (l *Lexer) slideRight() {
	if l.nextCharIndex >= len(l.input) {
		l.char = 0 // 0 = ASCII Code "NUL"
	} else {
		l.char = l.input[l.nextCharIndex]
	}

	l.currCharIndex = l.nextCharIndex
	l.nextCharIndex += 1
}

func (l *Lexer) ReadToken() tokenizer.Token {
	// Create an uninitialised token
	var currTok tokenizer.Token

    l.skipWhitespace()

	// Depending on what the current ch from the lexers input, create a new token
	switch l.char {
	case ':':
        currTok = newToken(tokenizer.COLON, l.char) 
    case '!':
        currTok = newToken(tokenizer.BANG, l.char)
    case '@':
        currTok = newToken(tokenizer.AT, l.char)
    case '*':
        currTok = newToken(tokenizer.STAR, l.char)
    case '-':
        currTok = newToken(tokenizer.DASH, l.char)
    case '=':
        currTok = newToken(tokenizer.EQUALS, l.char)
	case 0:
		currTok.Literal = ""
		currTok.Type = tokenizer.EOF
	default:
        // At this point if it ain't a special char, its just gonna get filtered out
		if isLetter(l.char) {
			currTok.Literal = l.readWords()
			currTok.Type = tokenizer.WORD
			return currTok
		} else if isDigit(l.char) {
			currTok.Literal = l.readNumber()
			currTok.Type = tokenizer.WORD
			return currTok
		} else {
			// If it's outside of ASCII (No unicode support)
			currTok = newToken(tokenizer.ILLEGAL, l.char)
		}
	}

	l.slideRight()
	return currTok
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.slideRight()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readWords() string {
	index := l.currCharIndex
	for isLetter(l.char) {
		l.slideRight()
	}
	return l.input[index:l.currCharIndex]

}

func (l *Lexer) readNumber() string {
	index := l.currCharIndex
	for isDigit(l.char) {
		l.slideRight()
	}
	return l.input[index:l.currCharIndex]
}

func newToken(tokenType tokenizer.TokenType, ch byte) tokenizer.Token {
	return tokenizer.Token{Type: tokenType, Literal: string(ch)}
}
