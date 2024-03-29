package slice

import "reflect"

func Flatten[T any](data any) []T {
TYPE_SWITCH:
	switch v := data.(type) {
	case T:
		return []T{v}
	case *T:
		return []T{*v}
	case []T:
		return v
	case *[]T:
		return *v
	case [][]T:
		var result []T
		for i := range v {
			result = append(result, v[i]...)
		}
		return result
	case *[][]T:
		var result []T
		for i := range *v {
			result = append(result, (*v)[i]...)
		}
		return result
	default:
		rv := reflect.ValueOf(data)
		switch rv.Kind() {
		case reflect.Pointer:
			data = rv.Elem().Interface()
			goto TYPE_SWITCH
		case reflect.Slice:
			var result []T
			for i := 0; i < rv.Len(); i++ {
				elem := rv.Index(i).Interface()
				result = append(result, Flatten[T](elem)...)
			}
			return result
		case reflect.Array:
			var result []T
			for i := 0; i < rv.Len(); i++ {
				elem := rv.Index(i).Interface()
				result = append(result, Flatten[T](elem)...)
			}
			return result
		case reflect.Struct:
			var result []T
			for i := 0; i < rv.NumField(); i++ {
				elem := rv.Field(i).Interface()
				result = append(result, Flatten[T](elem)...)
			}
			return result
		case reflect.Map:
			var result []T
			for _, key := range rv.MapKeys() {
				elem := rv.MapIndex(key).Interface()
				result = append(result, Flatten[T](key.Interface())...)
				result = append(result, Flatten[T](elem)...)
			}
			return result
		}
	}
	return nil
}
