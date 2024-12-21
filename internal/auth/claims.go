package auth

import (
	"github.com/kataras/iris/v12/middleware/jwt"
)

// Claims - is a custom JWT claims struct.
type Claims struct {
	jwt.Claims
	UserID int64 `json:"user_id"`
}
