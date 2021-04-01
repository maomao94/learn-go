package main

import (
	"fmt"
	"learn-go/gorm/model"

	"gorm.io/gorm"
)

func main() {
	db := model.Gorm()

	// Email 的 ID 是 `10`
	user := model.User{Model: gorm.Model{ID: 166}}

	db.Delete(&user)
	// DELETE from emails where id = 10;

	// 带额外条件的删除
	db.Where("name = ?", "jinzhu").Delete(&user)
	// DELETE from emails where id = 10 AND name = "jinzhu";

	fmt.Println("complete")
}
