package models

type User struct {
	ID       string   `bson:"_id,omitempty" json:"id"`
	Username string   `bson:"username" json:"username"`
	Password string   `bson:"password" json:"password"`
	Wallet   []string `bson:"wallet" json:"wallet"`
}
