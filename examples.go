package golangexamples

import "github.com/ehteshamz/greetings"

func ConcatSlice(sliceToConcat []byte) string {
	var temp string

	for index := 0; index < len(sliceToConcat); index++ {
		if index < len(sliceToConcat)-1 {
			temp += string(sliceToConcat[index]) + "-"
		} else {
			temp += string(sliceToConcat[index])
		}
	}
	return temp
}

func Encrypt(sliceToEncrypt []byte, ceasarCount int) {
	for index := 0; index < len(sliceToEncrypt); index++ {
		sliceToEncrypt[index] += byte(ceasarCount)
	}
}

func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}
