package main

import ("net/url")
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

func getUserByMail(userMAil string) *User {
	db := getDB();

	result := ViewResponse{}
	parameters := url.Values{}
	parameters.Set("key", "\""+userMAil+"\"")

	db.GetView("user_mail","get_users_by_mail", &result, &parameters)

	user := User{}
	user.ID = ""

	for _, value := range result.Rows {
		m, _ := value.(map[string]interface{})
		u, _ := m["value"].(map[string]interface{})
		
		id, _ := u["ID"].(string)
		user.ID = id
		lastname, _ := u["Lastname"].(string)
		user.Lastname = lastname
		firstname, _ := u["Firstname"].(string)
		user.Firstname = firstname
		mail, _ := u["Mail"].(string)
		user.Mail = mail
		password, _ := u["Password"].(string)
		user.Password = password
	}
	return &user
}
