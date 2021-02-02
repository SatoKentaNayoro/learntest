package main

import "fmt"

type TestRpc struct {}

func (t *TestRpc)Add(i int) int {
	i++
 	return i
}

func main() {
	//a := [2]string{"1","2"}
	//b := &a[0]
	//a[0]="3"
	//fmt.Println(*b)
	fmt.Println(byte(9))
}

