package variable

import "fmt"

func MultipleVariable() {
	var (
		name  string = "name"
		email string = "email@mail.com"
	)

	fmt.Println(name)
	fmt.Println(email)
}
