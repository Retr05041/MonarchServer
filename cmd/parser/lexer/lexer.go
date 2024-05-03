package lexer

import "MonarchServer/cmd/parser/tokenizer"

type Lexer struct {
	input         string
	currCharIndex int
	nextCharIndex int
	char          byte
	srcFlag       bool
	cmdFlag       bool
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
		l.srcFlag = true
        l.slideRight()
        return currTok
	case 0:
		currTok.Literal = ""
		currTok.Type = tokenizer.EOF
	default:
		// At this point if it ain't a special char, its just gonna get filtered out
		if isLetterOrDigit(l.char) {
			currTok.Literal = l.readWords()
            if l.srcFlag {
                currTok.Type = tokenizer.PREFIX
                l.srcFlag = false
                l.cmdFlag = true
            } else if l.cmdFlag || l.currCharIndex == 0 {
				currTok.Type = tokenizer.CMD
                l.cmdFlag = false
			} else {
				currTok.Type = tokenizer.PARAMETER
			}
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

func isLetterOrDigit(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || '0' <= ch && ch <= '9'
}

func (l *Lexer) readWords() string {
	index := l.currCharIndex
	for isLetterOrDigit(l.char) {
		l.slideRight()
	}
	return l.input[index:l.currCharIndex]

}

func newToken(tokenType tokenizer.TokenType, ch byte) tokenizer.Token {
	return tokenizer.Token{Type: tokenType, Literal: string(ch)}
}
