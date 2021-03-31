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
}
