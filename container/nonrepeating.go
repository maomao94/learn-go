package container

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func lengthOfNonRepeatingSubStr2(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

var lastOccurred = make([]int, 0xffff) //65535

func lengthOfNonRepeatingSubStr3(s string) int {
	//lastOccurred := make([]int, 0xffff) //65535
	for i := range lastOccurred {
		lastOccurred[i] = -1
	}
	//lastOccurred[0x8BFE] = 6
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func nonrepeating() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr("abdef"))
	fmt.Println(lengthOfNonRepeatingSubStr2("一二三四五六"))
	fmt.Println(lengthOfNonRepeatingSubStr2("一二三二一"))
	fmt.Println(lengthOfNonRepeatingSubStr2("就看看可能八年级八点半逼逼u"))
}
