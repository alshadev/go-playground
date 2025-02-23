package datatype

func GetExampleString() string {
	return "This is from GetString() function"
}

func GetExampleLengthString() int {
	return len(GetExampleString())
}

func GetCharOfStringByIndex(text string, index int) string {
	//should parse to string manually since it will return byte if only use 'text[index]'
	return string(text[index])
}
