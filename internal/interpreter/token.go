package interpreter

type Token struct {
    Type TokenType
    Lexeme string
    Line int
}

func NewToken(t TokenType, lx string, ln int) Token {
    return Token{
        Type: t,
        Lexeme: lx,
        Line: ln,
    }
}
