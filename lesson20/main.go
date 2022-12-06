package main

import (
	"database/sql"
	_ "encoding/json"
	_ "fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// grom demo2

// 定义模型
type User struct {
	gorm.Model   // 内嵌gorm.Model
	Name         string
	Age          sql.NullInt64 `gorm:"column:user_age"` //零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        //设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` //设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  //设置 num 为自增类型
	Address      string  `gorm:"index:addr""`     // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               //忽略本字段
}

// 使用`AnimalID`作为主键
type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}

// 唯一指定表名
func (Animal) TableName() string {
	return "ismi"
}

func main() {
	// 修改默认表名规则
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "SMS_" + defaultTableName
	}
	//连接mysql数据库
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SingularTable(true) // 禁用复数

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})

	// 使用user结构体创建一个名字叫 xiaohongmao 的表
	//db.Table("xiaohongmao").CreateTable(&User{})

}
