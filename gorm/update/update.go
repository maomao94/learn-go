package main

import (
	"fmt"
	"learn-go/gorm/model"
	"time"

	"gorm.io/gorm"
)

func main() {
	db := model.Gorm()

	// 更新

	// 保存所有字段
	var user model.User
	db.First(&user)
	user.Name = "jinzhu 2"
	user.Age = 100
	db.Save(&user)
	// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;

	// 更新单个列
	// 条件更新
	db.Model(&model.User{}).Where("active = ?", true).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;

	// User 的 ID 是 `111`
	db.Model(&user).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

	// 根据条件和 model 的值进行更新
	db.Model(&user).Where("active = ?", true).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;

	// 更新多列
	// 根据 `struct` 更新属性，只会更新非零值的字段
	db.Model(&user).Updates(model.User{Name: "hello", Age: 18})
	// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

	// 根据 `map` 更新属性
	db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18})
	// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

	// 更新选定字段
	// 使用 Map 进行 Select
	// User's ID is `111`:
	db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18})
	// UPDATE users SET name='hello' WHERE id=111;

	db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18})
	// UPDATE users SET age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

	// 使用 Struct 进行 Select（会 select 零值的字段）
	db.Model(&user).Select("Name", "Age").Updates(model.User{Name: "new_name", Age: 0})
	// UPDATE users SET name='new_name', age=0 WHERE id=111;

	// Select 所有字段（查询包括零值字段的所有字段）
	db.Model(&user).Select("*").Updates(model.User{Name: "jinzhu", Birthday: time.Now(), Age: 0, Model: gorm.Model{
		CreatedAt: time.Now(),
		ID:        user.ID,
	}})

	// Select 除 Role 外的所有字段（包括零值字段的所有字段）
	db.Model(&user).Select("*").Omit("Birthday").Updates(model.User{Name: "jinzhu", Birthday: time.Now(), Age: 0, Model: gorm.Model{
		CreatedAt: time.Now(),
		ID:        user.ID,
	}})
	fmt.Println("complete")
}
