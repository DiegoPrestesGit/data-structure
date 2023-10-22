package y2015

import (
	"bufio"
	"fmt"
	"os"
)

type My2015 struct{}

func (m My2015) Day01_A() {
	fmt.Println("I GOT YOU HOMIE")
	myFile, err := os.Open("./y2015/day01/input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(myFile)
	fileScanner.Split(bufio.ScanWords)

	str := ""
	for fileScanner.Scan() {
		str += fileScanner.Text()
	}
	fmt.Println(len(str))

	myFile.Close()
}
