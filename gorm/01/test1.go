package main

import (
	"fmt"
	"learn-go/gorm/model"
	"time"
)

func main() {
	db := model.Gorm()
	db.AutoMigrate(&model.User{})

	//创建记录
	user := model.User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	result := db.Create(&user) // 通过数据的指针来创建
	fmt.Println(user.ID)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)

	//用指定的字段创建记录

	//创建记录并更新给出的字段
	user.ID = 0
	result = db.Select("Name", "Age", "CreatedAt").Create(&user)
	fmt.Println(user.ID)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)

	//创建记录并更新未给出的字段。
	user.ID = 0
	result = db.Omit("Name", "Age", "CreatedAt").Create(&user)
	fmt.Println(user.ID)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
}
