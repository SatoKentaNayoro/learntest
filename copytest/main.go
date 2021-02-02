package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type testType struct {
	Aaa int64
	B []string
}

func main() {
	bbb := new(testType)
	ccc := testType{
		Aaa: 0,
		B: []string{}}
	marshal, err := json.Marshal(ccc)
	err = json.Unmarshal(marshal, bbb)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bbb)

	if err := equalWithInitialization(bbb, testType{});err == nil {
		fmt.Println(333)
	}else {
		fmt.Println(err)
	}
}


// compare value with an initialization model of given value's type
func equalWithInitialization(value,model interface{}) error {
	var (
		fv = reflect.Value{}
		fm = reflect.Value{}
	)

	if reflect.TypeOf(value).Kind() == reflect.Ptr {
		fv = reflect.ValueOf(value).Elem()
	}else {
		fv = reflect.ValueOf(value)
	}

	if reflect.TypeOf(model).Kind() == reflect.Ptr {
		fm = reflect.ValueOf(model).Elem()
	}else {
		fm = reflect.ValueOf(model)
	}
	fmt.Println(fmt.Sprintf("%v",fv))
	sprintfV := fmt.Sprintf("%v",fv)
	fmt.Println(fmt.Sprintf("%v",fm))
	sprintfM := fmt.Sprintf("%v",fm)
	if sprintfM == sprintfV {
		return nil
	}
	return errors.New("not equal")
}