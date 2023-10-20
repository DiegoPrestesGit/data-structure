package day01

import (
	"bufio"
	"fmt"
	"os"
)

func main () {
	readIt , err := os.ReadDir(".")
	fmt.Println(readIt)
	myFile, err := os.Open("./input.txt");
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(myFile)
	fileScanner.Split(bufio.ScanWords)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	myFile.Close()
}

