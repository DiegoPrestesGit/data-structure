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

func Check2letters2times (str string) (bool) {
	for i:=1; i<len(str);i++ {
		var repeatedOne string = string(str[i-1])+string(str[i])
		for j := i+2; j<len(str);j++ {
			verifying := string(str[j-1])+string(str[j])
			if repeatedOne == verifying {
				fmt.Println(str, repeatedOne)
				return true
			}
		}
	}

	return false
}

func CheckSameLetterWith1inBetween (str string) (bool) {
	for i:=1; i < len(str)-1;i++ {
		if str[i-1] == str[i+1] {
			return true
		}
	}

	return false
}

func (m SharedType) Day05_B(testFile string) {
	str := ReturnFileByLine(testFile)
	niceOnes := 0

	for _, v := range str {
		if CheckSameLetterWith1inBetween(v) && Check2letters2times(v) {
			niceOnes++
		}
	}

	fmt.Println("nice ones", niceOnes)
}
