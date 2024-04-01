package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	type args struct {
		valid     bool
		typ       string
		audience  string
		issuer    string
		header    string
		claims    string
		created   string
		expiresAt string
	}
	tests := []struct {
		name string
		args args
		want *Token
	}{
		{
			name: "valid token",
			args: args{
				valid:     true,
				typ:       "JWT",
				audience:  "my-audience",
				issuer:    "my-issuer",
				header:    "my-header",
				claims:    "my-claims",
				created:   "1234567890",
				expiresAt: "1234567890",
			},
			want: &Token{
				Valid:     true,
				Type:      "JWT",
				Audience:  "my-audience",
				Issuer:    "my-issuer",
				CreatedAt: "1234567890",
				ExpiresAt: "1234567890",
				Header:    "my-header",
				Claims:    "my-claims",
			},
		},
		{
			name: "invalid token",
			args: args{
				valid:     false,
				typ:       "JWT",
				audience:  "my-audience",
				issuer:    "my-issuer",
				header:    "my-header",
				claims:    "my-claims",
				created:   "1234567890",
				expiresAt: "1234567890",
			},
			want: &Token{
				Valid:     false,
				Type:      "JWT",
				Audience:  "my-audience",
				Issuer:    "my-issuer",
				CreatedAt: "1234567890",
				ExpiresAt: "1234567890",
				Header:    "my-header",
				Claims:    "my-claims",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewToken(tt.args.valid, tt.args.typ, tt.args.audience, tt.args.issuer, tt.args.header, tt.args.claims, tt.args.created, tt.args.expiresAt)
			assert.Equal(t, tt.want, got)
		})
	}
}
