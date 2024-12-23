package auth

// GetSigKey - returns sig aky  as []byte
func (a *Auth) GetSigKey() []byte {
	return []byte(sigKey)
}
