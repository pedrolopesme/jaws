package model

type Token struct {
	Header map[string]interface{}
	Claims map[string]interface{}
}

func NewToken(header map[string]interface{}, claims map[string]interface{}) *Token {
	return &Token{
		Header: header,
		Claims: claims,
	}
}
