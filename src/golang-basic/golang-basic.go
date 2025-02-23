package main

import (
	"fmt"
	datatype "golang-basic/data-type"
)

func main() {
	datatype.DataType()
	fmt.Println("Length of '"+datatype.GetExampleString()+"' is", datatype.GetExampleLengthString())
	fmt.Println(datatype.GetCharOfStringByIndex("Hello, World!", 2))
}
