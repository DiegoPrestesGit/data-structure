package y2015

import (
	"fmt"
)

func (m SharedType) Day01_A(testFile string) {
	str := ReturnFileContentByChar(testFile)

	floor := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			floor += 1
		} else {
			floor -= 1
		}
	}
	fmt.Println("FLOOR =>", floor)
}

func (m SharedType) Day01_B(testFile string) {
	str := ReturnFileContentByChar(testFile)

	floor := 0
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			floor += 1
		} else {
			floor -= 1
		}
		count += 1
		if floor == -1 {
			break
		}
	}
	fmt.Println("FLOOR =>", floor)
	fmt.Println("COUNTING =>", count)
}
