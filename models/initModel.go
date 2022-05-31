package models

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var PDO *gorm.DB

func init() {

}

func Conn() *gorm.DB {
	user := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.pwd")
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	dbname := viper.GetString("mysql.dbname")

	dsn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	PDO, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		}})
	if err != nil {
		fmt.Println(err)
	}

	sqlDB, err := PDO.DB()
	if err != nil {
		fmt.Println(err)
	}

	sqlDB.SetMaxOpenConns(100)

	return PDO
}
