package auth

import "time"

// Define some constants for our secret keys and token expiration durations.
const (
	accessSecret         = "top-secret-access"
	refreshSecret        = "top-secret-refresh"
	accessExpire         = 24 * time.Hour
	refreshExpire        = 24 * 7 * time.Hour
	accessEncriptionKey  = "some access sault"
	refreshEncriptionKey = "some refresh sault"
)
