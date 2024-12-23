package auth

import (
	"github.com/kataras/iris/v12/middleware/jwt"
)

// Claims - is a custom JWT claims struct.
type Claims struct {
	jwt.Claims
	UserID string `json:"user_id"`
}
