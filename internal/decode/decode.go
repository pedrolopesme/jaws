package decode

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/pedrolopeme/jaws/internal/model"
)

// Decode transforms a string JWT token into a valid Token structure
//
//	TODO add tests
func Decode(token string, key string) *model.Token {
	claims := jwt.MapClaims{}
	parsedToken, error := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})

	if (error) != nil {
		fmt.Print(error.Error())
	}

	// if !parsedToken.Valid {
	// 	return nil, fmt.Errorf("invalid token")
	// }

	// return model.NewToken(parsedToken.Header, claims), error
	return model.NewToken(parsedToken.Header, claims)
}
