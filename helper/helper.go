package helper

import "reflect"

func GetStructValueByKeyName(s interface{}, name string) interface{} {
	value := reflect.ValueOf(s)
	if value.Kind() != reflect.Struct {
		return nil
	}
	field := value.FieldByName(name)
	if field.IsValid() {
		return field.Interface()
	}
	return nil
}
