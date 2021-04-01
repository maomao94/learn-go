package main

import (
	"fmt"
	"learn-go/gorm/model"

	"gorm.io/gorm"
)

func main() {
	db := model.Gorm()

	// 删除一条记录
	// Email 的 ID 是 `10`
	user := model.User{Model: gorm.Model{ID: 166}}

	db.Delete(&user)
	// DELETE from emails where id = 10;

	// 带额外条件的删除
	db.Where("name = ?", "jinzhu").Delete(&user)
	// DELETE from emails where id = 10 AND name = "jinzhu";

	// 根据主键删除
	db.Delete(&model.User{}, 10)
	// DELETE FROM users WHERE id = 10;

	db.Scopes().Delete(&model.User{}, "10")
	// DELETE FROM users WHERE id = 10;

	// Unscoped 取消逻辑删除
	db.Unscoped().Delete(&model.User{}, []int{1, 2, 3})
	// DELETE FROM users WHERE id IN (1,2,3);

	fmt.Println("complete")
}
