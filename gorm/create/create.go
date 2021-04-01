package main

import (
	"fmt"
	"learn-go/gorm/model"
	"time"
)

func main() {
	db := model.Gorm()
	db.AutoMigrate(&model.User{})

	// 创建记录
	user := model.User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	result := db.Create(&user) // 通过数据的指针来创建
	fmt.Println(user.ID)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)

	// 用指定的字段创建记录

	// 创建记录并更新给出的字段
	user.ID = 0
	result = db.Select("Name", "Age", "CreatedAt").Create(&user)
	fmt.Println(user.ID)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)

	// 创建记录并更新未给出的字段。
	user.ID = 0
	result = db.Omit("Name", "Age", "CreatedAt").Create(&user)
	fmt.Println(user.ID)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)

	// 批量插入
	var users = []model.User{{Name: "jinzhu1", Birthday: time.Now()}, {Name: "jinzhu2", Birthday: time.Now()}, {Name: "jinzhu3", Birthday: time.Now()}}
	result = db.Create(&users)
	for _, user := range users {
		fmt.Println(user.ID)
		fmt.Println(result.Error)
		fmt.Println(result.RowsAffected)
	}
	users = []model.User{{Name: "jinzhu1", Birthday: time.Now()}, {Name: "jinzhu2", Birthday: time.Now()}, {Name: "jinzhu3", Birthday: time.Now()}}
	result = db.CreateInBatches(users, 100)
	fmt.Println(result.RowsAffected)

	// 根据map创建
	db.Model(&model.User{}).Create(map[string]interface{}{
		"Name": "jinzhu", "Age": 18,
	})

	// batch insert from `[]map[string]interface{}{}`
	db.Model(&model.User{}).Create([]map[string]interface{}{
		{"Name": "jinzhu_1", "Age": 18},
		{"Name": "jinzhu_2", "Age": 20},
	})
}
