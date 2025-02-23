package variable

import "fmt"

func ConstantVariable() {
	const name = "const name"
	const email = "const.email@mail.com"

	// name = "updated name" << (error) unable to update value since use constant variable

	fmt.Println(name)
	//fmt.Println(email)

	//note: still able to run if any declared constant variable that not used
}

func MultipleConstantVariable() {
	const (
		name  = "multiple const name"
		email = "const.email@mail.com"
	)

	fmt.Println(name)
}
