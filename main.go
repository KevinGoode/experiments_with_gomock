package main

import (
	"fmt"

	"./apis"
)

func main() {
	fmt.Println("Some dummy executable that instantiates some structs and links them together")
	num := apis.NewNumberStruct(42)
	str := apis.NewStringStruct("42nd Street")
	cont := apis.NewContainerStruct(str, num)
	fmt.Println(cont.WhoAmI())

}
