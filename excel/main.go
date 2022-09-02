package main

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
)

func compareStrSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (b == nil) != (a == nil) {
		return false
	}
	for key, value := range a {
		if value != b[key] {
			return false
		}
	}
	return true
}

func main() {
	skipHeader := false
	fixedHeader := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13"}
	file, err := excelize.OpenFile("test.xlsx")
	if err != nil {
		panic(err)
	}
	rows, err := file.Rows("sheet1")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			panic(err)
		}
		if skipHeader {
			if compareStrSlice(row, fixedHeader) {
				skipHeader = false
				continue
			} else {
				panic(errors.New("Excel格式错误"))
			}
		}
		if len(row) != len(fixedHeader) {
			continue
		}
		for _, value := range row {
			fmt.Print(value + ",")
		}
		fmt.Print("\n")
	}
}
