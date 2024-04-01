package decode

import (
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/pedrolopeme/jaws/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	type args struct {
		token string
		key   string
	}
	tests := []struct {
		name string
		args args
		want *model.Token
	}{
		{
			name: "valid token",
			args: args{
				token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
				key:   "your-256-bit-secret",
			},
			want: &model.Token{Algorithm: "", Valid: true, Type: "", Audience: "", Issuer: "", CreatedAt: "2018-01-17 23:30:22Z", ExpiresAt: "", Header: "{\"alg\":\"HS256\",\"typ\":\"JWT\"}", Claims: "{\"iat\":1516239022,\"name\":\"John Doe\",\"sub\":\"1234567890\"}"},
		},
		{
			name: "invalid token",
			args: args{
				token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJleHAiOjE1MTYyMzkwODJ9.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
				key:   "wrong-secret",
			},
			want: &model.Token{Algorithm: "", Valid: false, Type: "", Audience: "", Issuer: "", CreatedAt: "2018-01-17 23:30:22Z", ExpiresAt: "2018-01-17 23:31:22Z", Header: "{\"alg\":\"HS256\",\"typ\":\"JWT\"}", Claims: "{\"exp\":1516239082,\"iat\":1516239022,\"name\":\"John Doe\",\"sub\":\"1234567890\"}"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Decode(tt.args.token, tt.args.key)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDecodeClaims(t *testing.T) {
	type args struct {
		claims jwt.MapClaims
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid claims",
			args: args{
				claims: jwt.MapClaims{
					"sub":  "1234567890",
					"name": "John Doe",
					"iat":  1516239022,
					"exp":  1516239082,
				},
			},
			want: "{\"exp\":1516239082,\"iat\":1516239022,\"name\":\"John Doe\",\"sub\":\"1234567890\"}",
		},
		{
			name: "empty claims",
			args: args{
				claims: jwt.MapClaims{},
			},
			want: "{}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := decodeClaims(tt.args.claims)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetClaimStringValue(t *testing.T) {
	type args struct {
		claims jwt.MapClaims
		claim  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid claim",
			args: args{
				claims: jwt.MapClaims{
					"sub":  "1234567890",
					"name": "John Doe",
				},
				claim: "sub",
			},
			want: "1234567890",
		},
		{
			name: "invalid claim",
			args: args{
				claims: jwt.MapClaims{
					"sub":  "1234567890",
					"name": "John Doe",
				},
				claim: "invalid",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getClaimStringValue(tt.args.claims, tt.args.claim)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetClaimDateValue(t *testing.T) {
	type args struct {
		claims jwt.MapClaims
		claim  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid claim",
			args: args{
				claims: jwt.MapClaims{
					"sub":  "1234567890",
					"name": "John Doe",
					"iat":  float64(1516239022),
					"exp":  float64(1516239082),
				},
				claim: "iat",
			},
			want: "2018-01-17 23:30:22Z",
		},
		{
			name: "invalid claim",
			args: args{
				claims: jwt.MapClaims{
					"sub":  "1234567890",
					"name": "John Doe",
					"iat":  float64(1516239022),
					"exp":  float64(1516239082),
				},
				claim: "invalid",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getClaimDateValue(tt.args.claims, tt.args.claim)
			assert.Equal(t, tt.want, got)
		})
	}
}
