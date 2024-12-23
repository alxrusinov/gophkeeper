package auth

import "time"

// GetAccessTokenExp - return expiring of access token
func (a *Auth) GetAccessTokenExp() time.Duration {
	return accessExpire
}
