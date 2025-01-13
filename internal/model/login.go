package model

// Login - model of login structure
type Login struct {
	// Username - name of user
	Username string `json:"username" bson:"username"`
	// Password - password of user
	Password string `json:"password" bson:"password"`
}
