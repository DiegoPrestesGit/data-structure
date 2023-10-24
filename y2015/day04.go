package y2015

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func (m SharedType) Day04_A(testFile string) {
	str := ReturnFileByLine(testFile)[0]
	fmt.Println(str)

	count := 0
	for true {
		data := []byte(str+strconv.Itoa(count))
		hash := md5.Sum(data)
		hashStr := fmt.Sprintf("%x", hash)

		if hashStr[:5] == "00000" {
			fmt.Println()
			fmt.Println(count)
			break
		}
		count++
	}
}

func (m SharedType) Day04_B(testFile string) {
	str := ReturnFileByLine(testFile)[0]
	fmt.Println(str)

	count := 0
	for true {
		data := []byte(str+strconv.Itoa(count))
		hash := md5.Sum(data)
		hashStr := fmt.Sprintf("%x", hash)

		if hashStr[:6] == "000000" {
			fmt.Println()
			fmt.Println(count)
			break
		}
		count++
	}
}

