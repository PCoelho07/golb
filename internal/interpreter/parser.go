package interpreter

import (
	"fmt"
	"slices"

	"github.com/PCoelho07/golb/internal/loadbalancer"
)

type Parser struct {
	tokens []Token
	pos    int
}

func NewParser(t []Token) *Parser {
	return &Parser{
		tokens: t,
	}
}

func (p *Parser) Parse() (*loadbalancer.LoadBalancerConfig, error) {
    allowedRootTokens := []TokenType{BackendDirective, EOF}
    lbConfig := &loadbalancer.LoadBalancerConfig{}

    for {
        t := p.current()

        if !slices.Contains(allowedRootTokens, t.Type) {
            return nil, fmt.Errorf("unexpected %s at line %d", t.Lexeme, t.Line)
        }

        if t.Type == BackendDirective {
            backends, err := p.ParseBackendList()
            if err != nil {
                return nil, err
            }

            lbConfig.SetBackendUrls(backends)
            continue
        }

        if t.Type == EOF {
            break 
        }

        p.pos++
    }
    

    _, err := p.consume(EOF)
    if err != nil {
        return nil, err
    }

    return lbConfig, nil
}

func (p *Parser) current() Token {
	if p.pos >= len(p.tokens) {
		return NewToken(Illegal, "EOF", p.pos)
	}

	return p.tokens[p.pos]
}

func (p *Parser) consume(tt TokenType) (Token, error) {
	t := p.current()
	if tt != t.Type {
		return t, fmt.Errorf("unexpected token %s at line %d, expected %s", t.Lexeme, t.Line, tt)
	}

	p.pos++

	return t, nil
}

func (p *Parser) ParseBackendList() ([]string, error) {
	_, err := p.consume(BackendDirective)
	if err != nil {
		return nil, err
	}

	_, err = p.consume(DelimiterO)
	if err != nil {
		return nil, err
	}

    var urls []string
    for {
        tk := p.current()

        if tk.Type == DelimiterC {
            break
        }

        if tk.Type != BackendUrl {
            return nil, fmt.Errorf("unexpected token %s at line %d, expected URL", tk.Lexeme, tk.Line)
        }

        urls = append(urls, tk.Lexeme)
        p.pos++
    }

    _, err = p.consume(DelimiterC)
	if err != nil {
		return nil, err
	}

    return urls, nil
}
