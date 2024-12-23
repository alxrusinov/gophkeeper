package auth

import "time"

// Define some constants for our secret keys and token expiration durations.
const (
	accessSecret         = "signature_hmac_secret_shared_key"
	refreshSecret        = "signature_hmac_secret_shared_keyh"
	accessExpire         = 24 * time.Hour
	refreshExpire        = 24 * 7 * time.Hour
	accessEncriptionKey  = "signature_hmac_secret_shared_key"
	refreshEncriptionKey = "signature_hmac_secret_shared_key"
	sigKey               = "signature_hmac_secret_shared_key"
)
