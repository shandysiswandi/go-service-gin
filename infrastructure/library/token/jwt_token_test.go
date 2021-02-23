package token_test

import (
	"go-service-gin/infrastructure/library/token"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewJWTToken(t *testing.T) {
	tok := token.NewJWTToken("a", "b")

	pd := token.PayloadData{
		ID:    "id",
		Email: "email",
	}
	exp := time.Hour

	/**************
	* Access Token
	***************/
	acc, err := tok.NewAccessToken(pd, exp)
	assert.NoError(t, err)
	assert.NotEmpty(t, acc)

	accPyaload, err := tok.VerifyAccessToken(acc)
	assert.NoError(t, err)
	assert.NotEmpty(t, accPyaload)

	assert.Equal(t, "id", accPyaload.Data.ID)
	assert.Equal(t, "email", accPyaload.Data.Email)
	assert.NotZero(t, accPyaload.ID)
	assert.NotEqual(t, uuid.New(), accPyaload.ID)
	assert.WithinDuration(t, time.Now(), accPyaload.IssuedAt, time.Second)
	assert.WithinDuration(t, time.Now().Add(exp), accPyaload.ExpiredAt, time.Second)

	/**************
	* Refresh Token
	***************/
	exp = 2 * time.Hour
	ref, err := tok.NewRefreshToken(pd, exp)
	assert.NoError(t, err)
	assert.NotEmpty(t, acc)

	refPyaload, err := tok.VerifyRefreshToken(ref)
	assert.NoError(t, err)
	assert.NotEmpty(t, refPyaload)

	assert.Equal(t, "id", refPyaload.Data.ID)
	assert.Equal(t, "email", refPyaload.Data.Email)
	assert.NotZero(t, refPyaload.ID)
	assert.NotEqual(t, uuid.New(), refPyaload.ID)
	assert.WithinDuration(t, time.Now(), refPyaload.IssuedAt, time.Second)
	assert.WithinDuration(t, time.Now().Add(exp), refPyaload.ExpiredAt, time.Second)
}

func TestNewJWTToken_Expired(t *testing.T) {
	tok := token.NewJWTToken("a", "b")

	pd := token.PayloadData{
		ID:    "id",
		Email: "email",
	}
	exp := -time.Hour

	/**************
	* Access Token
	***************/
	acc, err := tok.NewAccessToken(pd, exp)
	assert.NoError(t, err)
	assert.NotEmpty(t, acc)

	accPyaload, err := tok.VerifyAccessToken(acc)
	assert.Error(t, err)
	assert.EqualError(t, err, token.ErrExpiredToken.Error())
	assert.Nil(t, accPyaload)

	/**************
	* Refresh Token
	***************/
	exp = -2 * time.Hour
	ref, err := tok.NewRefreshToken(pd, exp)
	assert.NoError(t, err)
	assert.NotEmpty(t, acc)

	refPyaload, err := tok.VerifyRefreshToken(ref)
	assert.Error(t, err)
	assert.EqualError(t, err, token.ErrExpiredToken.Error())
	assert.Nil(t, refPyaload)
}

func TestNewJWTToken_Invalid(t *testing.T) {
	p := token.NewPayload(token.PayloadData{ID: "id", Email: "email"}, 1)
	assert.NotNil(t, p)

	jwtToken, err := jwt.NewWithClaims(jwt.SigningMethodNone, p).
		SignedString(jwt.UnsafeAllowNoneSignatureType)
	assert.NoError(t, err)

	tok := token.NewJWTToken("a", "b")
	/**************
	* Access Token
	***************/
	accPyaload, err := tok.VerifyAccessToken(jwtToken)
	assert.Error(t, err)
	assert.EqualError(t, err, token.ErrInvalidToken.Error())
	assert.Nil(t, accPyaload)

	/**************
	* Refresh Token
	***************/
	refPyaload, err := tok.VerifyRefreshToken(jwtToken)
	assert.Error(t, err)
	assert.EqualError(t, err, token.ErrInvalidToken.Error())
	assert.Nil(t, refPyaload)
}
