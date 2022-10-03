package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func (e *Decoder) scan(st any, record []string) error {
	// 1) Провалидировать входные данные:
	//    1.1 нам передали структуру
	//    1.2 нам передали указатель на нее
	// 2) Все поля структуры:
	//    2.1 узнать его тип
	//    2.2 посмотреть теги и понять из какого поля парсим
	//    2.3 распарсить значени
	//    2.4 записать значение
	t := reflect.TypeOf(st)
	if k := t.Kind(); k != reflect.Pointer {
		return fmt.Errorf("bad kind %v", k)
	}

	t = t.Elem()
	if k := t.Kind(); k != reflect.Struct {
		return fmt.Errorf("bad kind %v", k)
	}

	v := reflect.ValueOf(st).Elem()

	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)
		ft := fv.Type()

		name := t.Field(i).Name
		if csvTag := t.Field(i).Tag.Get("csv"); csvTag != "" {
			name = csvTag
		}
		idx, ok := e.names[name]
		if !ok {
			return fmt.Errorf("not found in map")
		}
		csvVal := record[idx]

		switch ft.Kind() {
		case reflect.String:
			fv.SetString(csvVal)
		case reflect.Int:
			x, err := strconv.ParseInt(csvVal, 10, 64)
			if err != nil {
				return err
			}
			fv.SetInt(x)
		case reflect.Float64:
			x, err := strconv.ParseFloat(csvVal, 64)
			if err != nil {
				return err
			}
			fv.SetFloat(x)
		default:
			return fmt.Errorf("unsupported type")
		}
	}

	return nil
}
