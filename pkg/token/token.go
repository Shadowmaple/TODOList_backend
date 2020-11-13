package token

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	// ErrMissingHeader means the `Authorization` header was empty.
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")
)

// Context is the context of the JSON web token.
type Context struct {
	ID uint32
}

// ParseRequest gets the token from the header and
// pass it to the Parse function to parses the token.
func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	if len(header) == 0 {
		return nil, ErrMissingHeader
	}

	payload, err := ResolveToken(header)
	if err != nil {
		return nil, err
	}

	ctx := &Context{
		ID: payload.ID,
	}

	return ctx, nil
}
