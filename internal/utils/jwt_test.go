package utils_test

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo-list-task/internal/utils"
)

func TestGenerateJWT(t *testing.T) {
	token, err := utils.GenerateJWT()

	assert.NotEmpty(t, token)
	assert.NoError(t, err)
}

func generateValidJWT(secret string, expires time.Duration) (string, error) {
	claims := utils.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expires)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func TestValidateJWT_Success(t *testing.T) {
	secret := "secret"
	tokenString, err := generateValidJWT(secret, time.Hour)

	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	claims, err := utils.ValidateJWT(tokenString)

	assert.NoError(t, err)
	assert.NotNil(t, claims)

	castedClaims, ok := claims.(*utils.Claims)
	assert.True(t, ok)
	assert.WithinDuration(t, time.Now().Add(time.Hour), castedClaims.ExpiresAt.Time, time.Minute)
}

func TestValidateJWT_InvalidSignature(t *testing.T) {
	tokenString, err := generateValidJWT("wrong_secret", time.Hour)

	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	claims, err := utils.ValidateJWT(tokenString)

	assert.Error(t, err)
	assert.Nil(t, claims)
	assert.Contains(t, err.Error(), "signature is invalid")
}

func TestValidateJWT_ExpiredToken(t *testing.T) {
	secret := "secret"
	tokenString, err := generateValidJWT(secret, -time.Hour)

	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	claims, err := utils.ValidateJWT(tokenString)

	assert.Error(t, err)
	assert.Nil(t, claims)
	assert.Contains(t, err.Error(), "token is expired")
}
