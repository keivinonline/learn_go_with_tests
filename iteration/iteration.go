package iteration

import "strings"

func Repeat(character string, repeatCount int) (result string) {
	//// 100-120+ ns/op
	// for i := 0; i < repeatCount; i++ {
	// 	result += character
	// }
	// return result

	//// 40-50 ns/op
	return strings.Repeat(character, repeatCount)

}
