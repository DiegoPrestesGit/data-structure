package y2015

import (
	"bufio"
	"fmt"
	"os"
)

type My2015 struct{}

func readFile(testFile string) (string) {
	myFile, err := os.Open(testFile)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(myFile)
	fileScanner.Split(bufio.ScanWords)

	str := ""
	for fileScanner.Scan() {
		str += fileScanner.Text()
	}

	myFile.Close()
	return str
}

func (m My2015) Day01_A(testFile string) {
	fmt.Println("RUNNING Day01_A")
	str := readFile(testFile)

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
