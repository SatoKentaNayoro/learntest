package main

import (
	"fmt"
	"path"
)

func main() {
	//decodeString, _ := hex.DecodeString("63bdb5c463f46574f02bab0e7204354d39cc8ae726f46991b489de6a420368b5")
	//toString := base64.StdEncoding.EncodeToString(decodeString)
	//fmt.Println(toString)
	pathDir := path.Dir("/Users/iron_huang/IdeaProjects/tests")
	fmt.Println(pathDir)
}
