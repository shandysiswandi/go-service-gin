package token

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTToken is
type JWTToken struct {
	accessKey  string
	refreshKey string
}

// NewJWTToken is
func NewJWTToken(ak, rk string) Token {
	return &JWTToken{ak, rk}
}

// NewAccessToken is
func (j *JWTToken) NewAccessToken(pd PayloadData, exp time.Duration) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, NewPayload(pd, exp)).SignedString([]byte(j.accessKey))
}

// NewRefreshToken is
func (j *JWTToken) NewRefreshToken(pd PayloadData, exp time.Duration) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, NewPayload(pd, exp)).SignedString([]byte(j.refreshKey))
}

// VerifyAccessToken is
func (j *JWTToken) VerifyAccessToken(t string) (*Payload, error) {
	return j.verifyToken(t, j.accessKey)
}

// VerifyRefreshToken is
func (j *JWTToken) VerifyRefreshToken(t string) (*Payload, error) {
	return j.verifyToken(t, j.refreshKey)
}

func (JWTToken) verifyToken(t, k string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(k), nil
	}

	jwtToken, err := jwt.ParseWithClaims(t, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
