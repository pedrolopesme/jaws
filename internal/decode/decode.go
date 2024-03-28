package decode

import (
	"encoding/json"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/pedrolopeme/jaws/internal/model"
)

// Decode transforms a string JWT token into a valid Token structure
func Decode(token string, key string) *model.Token {
	claims := jwt.MapClaims{}
	parsedToken, _ := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})

	return model.NewToken(
		parsedToken.Valid,
		getClaimValues(claims, "aud"),
		getClaimValues(claims, "iss"),
		decodeClaims(parsedToken.Header),
		decodeClaims(claims))
}

// decodeClaims recursively decodes a jwt.MapClaims object into a JSON string.
func decodeClaims(claims jwt.MapClaims) string {
	parsedClaims, err := json.Marshal(claims)
	if err != nil {
		return ""
	}
	return string(parsedClaims)
}

// getClaimValues retrieves the value of a specific claim from a jwt.MapClaims object.
func getClaimValues(claims jwt.MapClaims, claim string) string {
	return claims[claim].(string)
}
