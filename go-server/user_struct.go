package main

type User struct {
	Type string
	ID string
	Lastname string
	Firstname string
	Mail string
	Password string
}

func getEmptyUser() *User {
	project := User{
		Type: "User",
		ID: getUUID()}
	return &project
}
