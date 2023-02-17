package tools

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/gookit/color"
	"log"
	"time"
)

type JwtStruct struct {
	Name    string `json:"name"`
	Context string `json:"context"`
	jwt.StandardClaims
}

var SignKey = []byte("1SAfaf49ga")

func JwtEncoder(name string, context string, expDuration int64) (string, error) {
	a := JwtStruct{
		Name:    name,
		Context: context,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + expDuration,
			Issuer:    "ACM_LiRuinian",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a)
	tokenString, err := token.SignedString(SignKey)
	if err != nil {
		log.Println(color.FgRed.Render(err.Error()))
		return "", err
	}
	return tokenString, nil
}

func JwtDecoder(tokenString string) (string, string, error) {
	s, err := ParseToken(tokenString)
	if err != nil {
		return "", "", err
	}
	return s.Name, s.Context, nil

}

func ParseToken(tokenString string) (*JwtStruct, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtStruct{}, func(token *jwt.Token) (i interface{}, err error) {
		return SignKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JwtStruct); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
