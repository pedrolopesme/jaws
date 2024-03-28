package model

type Token struct {
	Algorithm string
	Valid     bool
	Audience  string
	Issuer    string
	CreatedAt string
	ExpiresAt string
	Header    string
	Claims    string
}

func NewToken(valid bool, audience, issuer, header, claims, created, expires string) *Token {
	return &Token{
		Valid:     valid,
		Audience:  audience,
		Issuer:    issuer,
		CreatedAt: created,
		ExpiresAt: expires,
		Header:    header,
		Claims:    claims,
	}
}
