package y2015

import (
	"fmt"
	"regexp"
)

func CheckConsecutiveCharacters(str string) (bool) {
	for i:= 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			return true
		}
	}
	return false
}

func (m SharedType) Day05_A(testFile string) {
	str := ReturnFileByLine(testFile)
	niceOnes := 0

	for _, v := range str {
		threeVows := `.*[aeiou].*[aeiou].*[aeiou].*`
		negativeCheckPattern := `(ab|cd|pq|xy)`

		matchVows := regexp.MustCompile(threeVows).MatchString(v)
		matchNegative := regexp.MustCompile(negativeCheckPattern).MatchString(v)

		if matchVows && !matchNegative && CheckConsecutiveCharacters(v) {
			niceOnes++
		}
	}

	fmt.Println()
	fmt.Println("nice ones", niceOnes)
}

