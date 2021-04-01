package main

import (
	"errors"
	"fmt"
	"learn-go/gorm/model"

	"github.com/golang-module/carbon"

	"gorm.io/gorm"
)

func main() {
	db := model.Gorm()

	var user model.User
	var users []model.User
	// 检索单个对象
	// 获取第一条记录（主键升序）
	db.First(&user)
	// SELECT * FROM users ORDER BY id LIMIT 1;

	// 获取一条记录，没有指定排序字段
	db.Take(&user)
	// SELECT * FROM users LIMIT 1;

	// 获取最后一条记录（主键降序）
	db.Last(&user)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	result := db.First(&user)
	//result.RowsAffected // 返回找到的记录数
	//result.Error        // returns error

	// 检查 ErrRecordNotFound 错误
	errors.Is(result.Error, gorm.ErrRecordNotFound)

	// 用主键检索
	db.First(&user, 10)
	// SELECT * FROM users WHERE id = 10;

	db.First(&user, "10")
	// SELECT * FROM users WHERE id = 10;

	db.Find(&users, []int{1, 2, 3})
	// SELECT * FROM users WHERE id IN (1,2,3);

	// 检索全部对象
	// 获取全部记录
	result = db.Find(&users)
	// SELECT * FROM users;
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)

	// 条件
	// String 条件

	// 获取第一条匹配的记录
	db.Where("name = ?", "jinzhu").First(&user)
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

	// 获取全部匹配的记录
	db.Where("name <> ?", "jinzhu").Find(&users)
	// SELECT * FROM users WHERE name <> 'jinzhu';

	// IN
	db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

	// LIKE
	db.Where("name LIKE ?", "%jin%").Find(&users)
	// SELECT * FROM users WHERE name LIKE '%jin%';

	// AND
	db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

	// Time
	// 三周前
	lastWeek := carbon.Parse(carbon.Now().ToDateString()).SubWeeks(1).ToDateTimeString() // 2020-02-08 13:14:15
	db.Where("updated_at > ?", lastWeek).Find(&users)
	// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

	// BETWEEN
	db.Where("created_at BETWEEN ? AND ?", lastWeek, carbon.Parse(carbon.Now().ToDateString()).ToDateTimeString()).Find(&users)
	// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

	// Struct & Map 条件
	// Struct
	db.Where(&model.User{Name: "jinzhu", Age: 20}).First(&user)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;

	// Map
	db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	// 主键切片条件
	db.Where([]int64{20, 21, 22}).Find(&users)
	// SELECT * FROM users WHERE id IN (20, 21, 22);

	// 指定结构体查询字段
	db.Where(&model.User{Name: "jinzhu", Age: 0}, "name", "Age").Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;

	db.Where(&model.User{Name: "jinzhu", Age: 0}, "Age").Find(&users)
	// SELECT * FROM users WHERE age = 0;

	// 内联条件
}
