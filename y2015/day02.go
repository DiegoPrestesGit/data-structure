package y2015

import (
	"fmt"
	"strconv"
	"strings"
)

func StrParseInt(str string) (int) {
	val, err := strconv.Atoi(str); if err != nil {
		panic(err)
	}
	return val

}

func (m SharedType) Day02_A(testFile string) {
	fmt.Println("RUNNING Day02_A")
	var fileStr []string = ReturnFileByLine(testFile)

	fullSum := 0
	for _, str := range fileStr {
		// str{3} length x width x height
		// calculate area by: (2*l*w) + (2*w*h) + (2*h*l) + (the size of the smallest size)
		str := strings.Split(str, "x")
		l := StrParseInt(str[0])
		w := StrParseInt(str[1])
		h := StrParseInt(str[2])
		sum := (2*l*w) + (2*w*h) + (2*h*l) + min((l*w), (w*h), (h*l))
		fullSum += sum
	}
		fmt.Println(fullSum)
}

func (m SharedType) Day02_B(testFile string) {
	fmt.Println("RUNNING Day02_A")
	var fileStr []string = ReturnFileByLine(testFile)

	fullSum := 0
	for _, str := range fileStr {
		// str{3} length x width x height
		// calculate area by: (2*l*w) + (2*w*h) + (2*h*l) + (the size of the smallest size)
		str := strings.Split(str, "x")
		l := StrParseInt(str[0])
		w := StrParseInt(str[1])
		h := StrParseInt(str[2])
		sum := (2*l*w) + (2*w*h) + (2*h*l) + min((l*w), (w*h), (h*l))
		fullSum += sum
	}
		fmt.Println(fullSum)
}
