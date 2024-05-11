package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rarrazaan/be-player-performance-app/internal/config"
	"github.com/rarrazaan/be-player-performance-app/internal/utils/errors"
)

var (
	JWTSigningMethod = jwt.SigningMethodHS256
)

type (
	AccessJWTClaim struct {
		jwt.RegisteredClaims
		UserID    string   `json:"user_id"`
		UserName  string   `json:"user_name"`
		UserEmail string   `json:"user_email"`
	}
	SignAccessTokenPayload struct {
		UserID    string
		UserName  string
		UserEmail string
	}
	IJWT interface {
		GenerateAccessToken(payload SignAccessTokenPayload) (*string, error)
	}
	jWebToken struct {
		config config.Config
	}
)

func NewJWT(config config.Config) IJWT {
	return &jWebToken{
		config: config,
	}
}

func (c AccessJWTClaim) Valid() error {
	now := time.Now()
	if !c.VerifyExpiresAt(now, true) {
		return errors.ErrAccessTokenExpired
	}

	return nil
}

func (j *jWebToken) GenerateAccessToken(payload SignAccessTokenPayload) (*string, error) {
	expiresAt := time.Now().Add(time.Hour * time.Duration(j.config.Jwt.AccessTokenExpiration))
	now := time.Now()

	registeredClaims := jwt.RegisteredClaims{
		Issuer:    j.config.ServiceName,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}

	claims := AccessJWTClaim{
		RegisteredClaims: registeredClaims,
		UserID:           payload.UserID,
		UserName:         payload.UserName,
		UserEmail:        payload.UserEmail,
	}

	accessToken := jwt.NewWithClaims(JWTSigningMethod, claims)
	t, err := accessToken.SignedString([]byte(j.config.Jwt.JWTSecret))
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func ValidateAccessToken(generateToken string, config config.Config) (*jwt.Token, error) {
	computeFunction := func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.ErrInvalidToken
		}

		return []byte(config.Jwt.JWTSecret), nil
	}

	token, err := jwt.ParseWithClaims(generateToken, new(AccessJWTClaim), computeFunction)
	if err != nil {
		if e, ok := err.(*jwt.ValidationError); ok {
			if e, ok := e.Inner.(*errors.CustomError); ok {
				return nil, e
			}

			return nil, err
		}
	}

	return token, nil
}

func ParseAccessTokenClaim(accessToken string, config config.Config) (*AccessJWTClaim, error) {
	token, _ := ValidateAccessToken(accessToken, config)
	if t, ok := token.Claims.(*AccessJWTClaim); ok {
		return t, nil
	}
	return nil, errors.ErrInvalidToken
}
