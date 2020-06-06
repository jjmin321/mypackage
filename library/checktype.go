package library

import "fmt"

// Checktype - 값의 타입을 확인해주는 라이브러리 , Ex) Checktype(num1)
func Checktype(a interface{}) {
	switch a.(type) {
	case bool:
		fmt.Println("This is bool,", a)
	case int, int8, int16, int32, int64:
		fmt.Println("This is int,", a)
	case float32, float64:
		fmt.Println("This is float,", a)
	case nil:
		fmt.Println("This is nil,", a)
	case string:
		fmt.Println("This is string,", a)
	default:
		fmt.Println("This is not bool, int, float, nil, string!,", a)
	}
}
