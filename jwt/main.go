package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GetToken() (string, error) {
	mySigningKey := []byte("AllYourBase")

	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 500,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
	return ss, nil
}

func ParseToken() {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE2NTMzNTg0NDUsImlzcyI6InRlc3QifQ.AYWWKof7YLqUViDRYaMAKYpcBBYeWhhISSJXxtbfDA8"
	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// sample token is expired.  override time so it parses as valid

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Foo, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}

}

//刷新token值
func FreshToken(tokenString string) (string, error) {

	jwt.TimeFunc = func() time.Time {

		return time.Unix(0, 0)
	}
	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte("AllYourBase"), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return GetToken()
	}
	return "", err
}

func main() {
	GetToken()
	// ParseToken()
}
