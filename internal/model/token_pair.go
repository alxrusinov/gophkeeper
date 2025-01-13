package model

// TokenPair - has access and refresh tokens apir
type TokenPair struct {
	AccessToken  string
	RefreshToken string
	Exp          int64
}
