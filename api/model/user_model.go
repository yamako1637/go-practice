package model

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetUser() User {
	return User{
		Id:   1,
		Name: "name",
	}
}
