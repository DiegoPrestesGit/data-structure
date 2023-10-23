package y2015

import (
	"fmt"
)

func (m My2015) Day01_B(testFile string) {
	str := readFile(testFile)

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

