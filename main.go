package main

import "fmt"
import (
	"ginserver/router"
)

func main(){
	fmt.Println("this is a gin project")
	router := router.InitRouter()
	router.Run(":8090")
}