package model

type Token struct {
	Algorithm string
	Valid     bool
	Type      string
	Audience  string
	Issuer    string
	CreatedAt string
	ExpiresAt string
	Header    string
	Claims    string
}

func NewToken(valid bool, typ, audience, issuer, header, claims, created, expires string) *Token {
	return &Token{
		Valid:     valid,
		Type:      typ,
		Audience:  audience,
		Issuer:    issuer,
		CreatedAt: created,
		ExpiresAt: expires,
		Header:    header,
		Claims:    claims,
	}
}
