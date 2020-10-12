package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	UserName string `gorm:"user_name" json:"username"`
	Password string `gorm:"password" json:"password"`
}

func CheckAuth(userName string, passWord string) bool {
	var (
		auth Auth
		ok   bool
	)
	err := db.Select("id").Where("user_name = ? and password = ?", userName, passWord).First(&auth).Error
	if err == nil {
		if auth.ID > 0 {
			ok = true
		}
	}
	return ok
}
