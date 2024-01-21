package domain

type User struct {
	Nickname    string   `bson:"nickname"`
	Password    string   `bson:"password"`
	Gender      string   `bson:"gender"`
	PhoneNumber string   `bson:"phoneNumber"`
	Interests   []string `bson:"interests"`
}
