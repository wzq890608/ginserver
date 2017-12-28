package user

import (
	"github.com/gin-gonic/gin"
	"ginserver/config/db"
	"ginserver/model"
	"log"
)

func GetUser() []*model.User {
	var users []*model.User
	err := db.DB.Raw("select user_id,user_name,age,password from tab_user").Scan(&users).Error
	if err != nil {
		log.Fatalf( err.Error())
		//c.JSON(http.StatusExpectationFailed, err)
	}
	return users
	//c.JSON(http.StatusOK, users)
}

func GetFirstUser(c *gin.Context) model.User{
	var user model.User
	err := db.DB.First(&user).Error
	if err != nil {
		log.Fatalf( err.Error())
	}
	return user
	//c.JSON(http.StatusOK, user)
}
