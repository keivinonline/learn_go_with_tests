package iteration

import "strings"

func Repeat(character string, repeatCount int) (result string) {
	// for i := 0; i < repeatCount; i++ {
	// 	result += character
	// }
	// return result

	return strings.Repeat(character, repeatCount)

}
