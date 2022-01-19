package basic

import (
	"fmt"
	"reflect"
)

func RangeSource(i interface{}, param *[]string) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	if !v.IsValid() {
		return
	}
	if t.Kind() == reflect.Ptr {
		if v.IsNil() {
			return
		}
		t = t.Elem()
		v = v.Elem()
		in := v.Interface()
		RangeSource(in, param)
		return
	}
	if t.Kind() == reflect.Struct {
		RangeStructFiled(i, param)
		return
	}

	if t.Kind() == reflect.Slice {
		RangeStructSlice(i, param)
		return
	}

	if t.Kind() == reflect.Map {
		RangeStructMap(i, param)
		return
	}
}

func RangeStructFiled(i interface{}, param *[]string) {
	tx := reflect.TypeOf(i)
	var v reflect.Value
	var value reflect.Value
	v = reflect.ValueOf(i)
	fieldNum := tx.NumField()
	for i := 0; i < fieldNum; i++ {
		if !tx.Field(i).IsExported() {
			continue
		}
		k := tx.Field(i).Name
		if v.Kind() == reflect.Struct {
			value = v.FieldByName(k)
		}
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				continue
			}
			value = v.Elem().FieldByName(k)
		}
		if value.Kind() != reflect.Struct && value.Kind() != reflect.Ptr {
			if value.Kind() == reflect.String {
				*param = append(*param, fmt.Sprintf("%s", value))
			}
		}
		if value.Kind() == reflect.Struct || value.Kind() == reflect.Ptr || value.Kind() == reflect.Slice || value.Kind() == reflect.Map {
			if value.IsValid() {
				if value.Kind() == reflect.Ptr {
					continue
				}
				RangeSource(value.Interface(), param)
			}
		}
	}
}

func RangeStructSlice(i interface{}, param *[]string) {
	v := reflect.ValueOf(i)
	for sl := 0; sl < v.Len(); sl++ {
		if v.Index(sl).Kind() == reflect.String {
			value := v.Index(sl)
			*param = append(*param, fmt.Sprintf("%s", value))
		}
		if v.Index(sl).Kind() == reflect.Interface {
			value := v.Index(sl)
			switch value.Interface().(type) {
			case string:
				*param = append(*param, fmt.Sprintf("%s", value))
				break
			default:
				RangeSource(value.Interface(), param)
				break
			}
		}

		if v.Index(sl).Kind() == reflect.Struct || v.Index(sl).Kind() == reflect.Ptr || v.Index(sl).Kind() == reflect.Slice || v.Index(sl).Kind() == reflect.Map {
			if v.Index(sl).Kind() == reflect.Ptr && v.Index(sl).IsNil() {
				continue
			}
			RangeSource(v.Index(sl).Interface(), param)
		}
	}
}

func RangeStructMap(i interface{}, param *[]string) {
	v := reflect.ValueOf(i)
	for _, key := range v.MapKeys() {
		value := v.MapIndex(key)
		if key.Kind() == reflect.String {
			*param = append(*param, fmt.Sprintf("%s", value))
		}

		if value.Kind() == reflect.String {
			*param = append(*param, fmt.Sprintf("%s", value))
		}

		if key.Kind() == reflect.Struct || key.Kind() == reflect.Ptr || key.Kind() == reflect.Slice || key.Kind() == reflect.Map {
			if value.Kind() == reflect.Ptr && value.IsNil() {
				continue
			}
			RangeSource(key.Interface(), param)
		}

		if value.Kind() == reflect.Struct || value.Kind() == reflect.Ptr || value.Kind() == reflect.Slice || value.Kind() == reflect.Map {
			if value.Kind() == reflect.Ptr && value.IsNil() {
				continue
			}
			RangeSource(value.Interface(), param)
		}
	}
}
