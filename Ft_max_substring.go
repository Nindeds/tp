package poulpe

import "fmt"

func Ft_max_substring(s string) int {
	var Highest rune
	var lasti int
	StringToRune := []rune(s)
	for i := 1; i < len(s)-1; i++ {
		fmt.Println("pd", string(StringToRune[i]))
		if StringToRune[i-1] >= StringToRune[i] && StringToRune[i-1] > Highest && StringToRune[i] > Highest {
			Highest = StringToRune[i-1]
			fmt.Println("highest", string(Highest))
			lasti += i
		} else {
			continue
		}
	}
	return lasti
}
