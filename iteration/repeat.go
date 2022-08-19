package iteration

import "strings"

const repeatCount = 5

func Repeat(char string) string {
	repeated := strings.Repeat(char, repeatCount)
	return repeated
}
