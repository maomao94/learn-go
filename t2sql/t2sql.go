package main

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"learn-go/t2sql/model"
	"strings"
)

func main() {
	ctx := context.Background()
	datasource := "developer:KvB4dSql@tcp(evcard-st-lan.mysql.rds.aliyuncs.com:3306)/vlms_assets?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	t := model.NewTAssetsVehicleModel(sqlx.NewMysql(datasource))
	build := t.RowBuilder()
	query, err := t.FindAll(ctx, build, "")
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range query {
		b := squirrel.Insert("aaa").Columns("id", "name").Values(v.Id, v.Vin)
		sql, args, _ := b.ToSql()
		sql = strings.Replace(sql, "?", "%d", 1)
		sql = strings.Replace(sql, "?", "'%s'", 1)
		query := fmt.Sprintf(sql, args...)
		fmt.Print(query + "\n")
	}
}
