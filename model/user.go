package model


type User struct {
	UserId    int    `gorm:"primary_key" gorm:"column:user_id" json:"userId"`
	UserName string `gorm:"column:user_name" json:"userName"`
	Age      int `json:"age"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return "tab_user"
}

