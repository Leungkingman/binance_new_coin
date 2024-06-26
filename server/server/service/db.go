package service

import (
	"binanceNewCoin/server/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var gormDB *gorm.DB

func InitDatabase() {
	var err error
	//host := os.Getenv("DB_HOST")
	//port := os.Getenv("DB_PORT")
	//user := os.Getenv("DB_USER")
	//password := os.Getenv("DB_PASSWORD")
	//dbName := os.Getenv("DB_NAME")
	//charset := os.Getenv("DB_CHARSET")

	host := "8.218.81.24"
	port := "3306"
	user := "Ryan"
	password := "xxxxx"
	dbName := "binance_new_coin"
	charset := "utf8"

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=" + charset + "&parseTime=true&loc=Local"
	global.DB_ENGINE, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := global.DB_ENGINE.DB()
	db.SetConnMaxLifetime(time.Minute * 3)
	fmt.Printf("global.DB_ENGINE = %v \n", global.DB_ENGINE)
	fmt.Printf("db = %v \n", db)
	if err != nil {
		fmt.Printf("InitDatabase err = %v \n", err)
	}
}

func PrintDB() {
	db, err := global.DB_ENGINE.DB()
	if err != nil {
		fmt.Printf("PrintDB err = %v \n", err)
	} else {
		fmt.Printf("PrintDB db = %v \n", db)
	}
}