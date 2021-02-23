package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// all errors token
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Token is
type Token interface {
	NewAccessToken(PayloadData, time.Duration) (string, error)
	NewRefreshToken(PayloadData, time.Duration) (string, error)
	VerifyAccessToken(string) (*Payload, error)
	VerifyRefreshToken(string) (*Payload, error)
}

// PayloadData is
type PayloadData struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// Payload is
type Payload struct {
	ID        uuid.UUID   `json:"id"`
	Data      PayloadData `json:"data"`
	IssuedAt  time.Time   `json:"issued_at"`
	ExpiredAt time.Time   `json:"expired_at"`
}

// NewPayload is
func NewPayload(pd PayloadData, exp time.Duration) *Payload {
	p := &Payload{
		ID:        uuid.New(),
		IssuedAt:  time.Now(),
		Data:      pd,
		ExpiredAt: time.Now().Add(exp),
	}

	return p
}

// Valid is
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
