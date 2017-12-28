package db

import (
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"ginserver/config"
	"log"
)

func NewDatabase(cnf *config.DatabaseConfig) (*gorm.DB, error) {
	args := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.DatabaseName,
	)
	db, err := gorm.Open(cnf.Type, args)
	//defer db.Close()
	db.LogMode(cnf.IsDevelopment)
	if err != nil {
		return db, err
	}
	return db, nil

}

var DB *gorm.DB

func init() {
	var err error
	conf := config.LoadConfig("config")
	DB, err = NewDatabase(&conf.Database)
	if err!=nil{
		log.Fatalf(err.Error())
	}
}

