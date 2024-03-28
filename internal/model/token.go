package model

type Token struct {
	Algorithm string
	Valid     bool
	Audience  string
	Issuer    string
	Header    string
	Claims    string
}

func NewToken(valid bool, audience, issuer, header, claims string) *Token {
	return &Token{
		Valid:    valid,
		Audience: audience,
		Issuer:   issuer,
		Header:   header,
		Claims:   claims,
	}
}
