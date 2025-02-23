package variable

import "fmt"

func Variable() {
	fmt.Println("Welcome to Variable")

	var indirectVarString string
	indirectVarString = "Indirect Variable"
	fmt.Println(indirectVarString)

	var directVarString string = "Direct Variable"
	fmt.Println(directVarString)

	var withoutDeclareDataType = "Automatically set variable as string without declare data type"
	fmt.Println(withoutDeclareDataType)

	withoutUsingVar := "Declare initial variable without using var and automatically read variable as string (depend on it value data type)"
	fmt.Println(withoutUsingVar)
}
