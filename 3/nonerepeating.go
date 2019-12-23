package main

import "fmt"

func lengthOfNonRepeatingSubstr(s string) int {
	lastOccuered := make(map[rune]int)
	start := 0
	maxLength := 0
	for i,ch := range []rune(s) {
		if lastI,ok := lastOccuered[ch];ok && lastI >= start{
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccuered[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubstr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubstr("我是"))

}
