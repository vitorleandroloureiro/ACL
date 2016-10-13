package models

type User struct {
	Name     string
	Username string
	Email    string
	Roles    []Role
}

func (user User) Save() bool {

	Dic.Get("db").(DB).Get("asdasd");

	//bucket.Upsert("u:" + user.Email, user, 0)
	//
	//var inUser User
	//bucket.Get("u:" + user.Email, &inUser)
	//fmt.Printf("User: %v\n", inUser)
	//
	return true;
}