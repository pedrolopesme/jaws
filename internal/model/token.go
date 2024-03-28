package model

type Token struct {
	Header string
	Claims string
}

func NewToken(header string, claims string) *Token {
	return &Token{
		Header: header,
		Claims: claims,
	}
}
