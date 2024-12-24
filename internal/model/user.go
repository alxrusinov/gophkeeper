package model

type User struct {
	ID       string `bson:"id"`
	Username string `bson:"username"`
}
