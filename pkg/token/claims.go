package token

import (
	"time"
)

// TokenClaims means a claim segment in a JWT.
type TokenClaims struct {
	ID           uint32 `json:"id"`
	QQOpenID     string `json:"qq_open_id"` // QQ 用户唯一标识符
	QQSessionKey string `json:"qq_session_key"`
	ExpiresAt    int64  `json:"expires_at"` // 过期时间（时间戳，10位）
}

// Valid checks whether the token is valid.
// Used by jwt-go package.
// Because of this method, the type TokenClaims implements the interface Claims of jwt-go.
func (c TokenClaims) Valid() error {
	now := time.Now().Unix()

	if !c.VerifyExpiresAt(now, false) {
		return ErrTokenExpired
	}

	return nil
}

// VerifyExpiresAt verifies the 'ExpiresAt' field.
// If required is false, this method will return true if the value matches or is unset.
func (c *TokenClaims) VerifyExpiresAt(now int64, required bool) bool {
	if c.ExpiresAt == 0 {
		return !required
	}
	return now <= c.ExpiresAt
}
