package jwt

import (
	"fmt"
	customErr "github.com/devarsh/vrpl/error"
	"github.com/devarsh/vrpl/jwt/model"
	"github.com/dgrijalva/jwt-go"
)

type JwtGen struct {
	signingMethod jwt.SigningMethod
	key           []byte
}

func (jg *JwtGen) GenerateJwtToken(payload *model.TokenPayload) (string, error) {
	token := jwt.NewWithClaims(jg.signingMethod, payload)
	ss, err := token.SignedString(jg.key)
	if err != nil {
		return "", customErr.JwtError(err)
	}
	return ss, nil
}

func (jg *JwtGen) VerifyJwtToken(tokenString string) (*model.TokenPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.TokenPayload{}, func(token *jwt.Token) (interface{}, error) {
		return jg.key, nil
	})
	if token.Valid {
		if customToken, ok := token.Claims.(*model.TokenPayload); ok {
			return customToken, nil
		}
		return nil, customErr.JwtError(fmt.Errorf("Error Parsing Custom Claims"))
	}
	return nil, customErr.JwtError(err)
}
