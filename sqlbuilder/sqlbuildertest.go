package main

import (
	"errors"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"reflect"
)

type Example interface {
	Equal(field string, value interface{}) error
	In(field string, value []string) error
}

type SqlExample struct {
	builder *sqlbuilder.SelectBuilder
	andExpr []string
}

func (s *SqlExample) In(field string, value []string) error {
	if value == nil || len(value) == 0 {
		return nil
	}
	s.andExpr = append(s.andExpr, s.builder.In(field, value))
	return nil
}

func (s *SqlExample) Equal(field string, value interface{}) error {
	val := reflect.ValueOf(value) // 获取reflect.Type类型
	kd := val.Kind()              // 获取到st对应的类别
	if kd == reflect.Struct {
		return errors.New("is struct")
	}
	if s.isBlank(val) {
		return errors.New("is empty")
	}
	s.andExpr = append(s.andExpr, s.builder.Equal(field, value))
	return nil
}

func (s SqlExample) isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

func NewSqlExample(builder *sqlbuilder.SelectBuilder) SqlExample {
	return SqlExample{
		builder: builder,
		andExpr: make([]string, 0),
	}
}

func main() {
	sql := sqlbuilder.Select("id", "name").From("demo.user").
		Where("status = 1").Limit(10).
		String()
	fmt.Println(sql)
	// Output:
	// SELECT id, name FROM demo.user WHERE status = 1 LIMIT 10
	fmt.Println("===================")

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "name", sb.As("COUNT(*)", "c"))
	sb.From("user")
	sb.Where(sb.In("status", 1, 2, 5),
		sb.Between("create_time", "2020-01-02", "2020-03-03"),
		sb.Equal("task_id", 000001)).OrderBy("update_time")
	sql, args := sb.Build()
	fmt.Println(sql)
	fmt.Println(args)
	// Output:
	// SELECT id, name, COUNT(*) AS c FROM user WHERE status IN (?, ?, ?) AND create_time BETWEEN ? AND ? AND task_id = ? ORDER BY update_time
	// [1 2 5 2020-01-02 2020-03-03 1]
	fmt.Println("===================")

	exampleList := []string{}
	statusList := [3]int{1, 2, 5}
	taskId := 9996
	sb1 := sqlbuilder.NewSelectBuilder()
	if len(statusList) > 0 {
		exampleList = append(exampleList, sb1.In("status", statusList))
	}
	if taskId > 0 {
		exampleList = append(exampleList, sb1.Equal("task_id", taskId))
	}
	sb1.Select("id", "name", "status")
	sb1.From("user")
	sb1.Where(exampleList...).OrderBy("update_time")
	sql, args = sb1.Build()
	fmt.Println(sql)
	fmt.Println(args)
	fmt.Println("===================")

	emptyList := [3]string{"1", "2", "3"}
	sb2 := sqlbuilder.NewSelectBuilder()
	example := NewSqlExample(sb2)
	example.Equal("status", "1")
	example.Equal("task_id", taskId)
	example.In("task_status", emptyList[:])
	sb2.Select("id", "name", "status")
	sb2.From("user")
	sb2.Where(example.andExpr...).OrderBy("update_time")
	sql, args = sb2.Build()
	fmt.Println(sql)
	fmt.Println(args)
}
