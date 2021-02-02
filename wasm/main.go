/**
 _*_ @Author: IronHuang _*_
 _*_ @blog:https://www.dvpos.com/ _*_
 _*_ @Date: 2021/2/1 下午9:20 _*_
**/

package main

import (
	"encoding/json"
	"strconv"
	"syscall/js"
)

type testS struct {
	A int
	B int
	C string
}

var done = make(chan struct{})

func main() {
	js.Global().Set("testWasm", js.FuncOf(testWasm))
	js.Global().Set("testDone", js.FuncOf(testDone))
	<-done
}

func testWasm(this js.Value, args []js.Value) interface{} {
	a, _ := strconv.Atoi(args[0].String())
	b, _ := strconv.Atoi(args[1].String())
	c := args[2].String()
	t := testS{
		A: a,
		B: b,
		C: c,
	}
	marshal, _ := json.Marshal(t)
	return js.ValueOf(string(marshal))
}

func testDone(this js.Value, args []js.Value) interface{} {
	done <- struct {
	}{}
	return "done"
}

func backValue(a int) string {
	itoa := strconv.Itoa(a)
	return itoa
}
