package models

type User struct {
	ID int `json:"id" xorm:"id"`
	Name string `json:"name" xorm:"name"`
	Email string `json:"email" xorm:"email"`
}

func (u *User) TableName() string {
	return "users"
}

func GetUser() *User {
	user := &User{}
	has, err := orm.Get(&user)
	// or has, err := engine.Id(27).Get(&user)

	if err != nil {
		println(err)
	}
	if !has {
		println("no data")
	}

	return user

}


func SaveUser(user *User) (int, error){
	_, err := orm.InsertOne(user)
	return user.ID, err
}
