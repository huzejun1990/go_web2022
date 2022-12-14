// @Author huzejun 2022/12/14 18:09:00
package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func InitMySQL() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}
