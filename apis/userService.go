package apis

import (
	"github.com/gin-gonic/gin"
	"ginserver/model"
	"ginserver/dao/user"
	"net/http"
	"sync"
	"ginserver/mytest"
)

func GetAllUser(c *gin.Context)  {
	var users []*model.User
	users = user.GetUser()
	c.JSON(http.StatusOK, users)
}

func TestThread(c *gin.Context)  {
	var wg sync.WaitGroup
	wg.Add(2)
	go mytest.TestThread1(wg)
	go mytest.TestThread2(wg)
	wg.Wait()
	c.String(http.StatusOK,"ok")
}