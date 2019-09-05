package models

type Auth struct {
	ID int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role string `json:"role"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username : username, Password : password}).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}

func AddAuth(username, passowrd string) bool {
	auth := Auth{Username: username, Password: passowrd, Role: "normal"}
	if err := db.Create(&auth).Error; err != nil {
		return false
	}
	return true
}

func GetAllUsers(pageSize, pageCurrent int) (users []Auth) {
	db.Order("ID desc").Find(&users)
	return
}

func GetUsersTotal() (count int) {
	db.Model(&Auth{}).Count(&count)
	return
}