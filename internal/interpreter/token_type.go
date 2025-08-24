package interpreter

type TokenType string

const (
    DelimiterO TokenType = "DELIMITER_OPEN"
    DelimiterC TokenType = "DELIMITER_CLOSE"
    BackendDirective TokenType = "BACKENDS"
    BackendUrl TokenType = "BACKEND_URL"
    Illegal TokenType = "ILLEGAL"
    EOF TokenType = "EOF"
)


