package interpreter

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

type Lexer struct {
	TokenList []Token
	Filename  string
}

func NewLexer(flnm string) *Lexer {
	return &Lexer{
		Filename: flnm,
	}
}

func (l *Lexer) Tokenize() error {
	f, err := os.Open(l.Filename)
	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	lCount := 1
	var tokenList []Token

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		var token Token
        r, _ := regexp.Compile("https?:\\/\\/[^\\s}]+|localhost|[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}(?:\\.[a-zA-Z]{2,})*")
		for _, w := range words {
			switch {
			case w == "{":
				token = NewToken(DelimiterO, w, lCount)
			case w == "}":
				token = NewToken(DelimiterC, w, lCount)
			case w == "backends":
				token = NewToken(BackendDirective, w, lCount)
            case r.MatchString(w):
                token = NewToken(BackendUrl, w, lCount)
			default:
                token = NewToken(Illegal, w, lCount)
			}
			tokenList = append(tokenList, token)
		}
		lCount++
	}

    tokenList = append(tokenList, NewToken(EOF, "EOF", lCount))
    l.TokenList = tokenList

	return nil
}
