package y2015

import (
	"fmt"
	"reflect"
)

func AddHouseNotVisited(housesVisited [][]int, position []int) [][]int {
	for _, v := range housesVisited {
		if reflect.DeepEqual(v, position) {
			return housesVisited
		}
	}

	housesVisited = append(housesVisited, position)
	return housesVisited
}

func (m SharedType) Day03_A(testFile string) {
	str := ReturnFileContentByChar(testFile)

	housesVisited := [][]int{{0, 0}}
	counter := []int{0, 0}
	for _, v := range str {
		switch v {
		case 'v':
			counter = []int{counter[0] - 1, counter[1]}
			housesVisited = AddHouseNotVisited(housesVisited, counter)
			break
		case '^':
			counter = []int{counter[0] + 1, counter[1]}
			housesVisited = AddHouseNotVisited(housesVisited, counter)
			break
		case '<':
			counter = []int{counter[0], counter[1] - 1}
			housesVisited = AddHouseNotVisited(housesVisited, counter)
			break
		case '>':
			counter = []int{counter[0], counter[1] + 1}
			housesVisited = AddHouseNotVisited(housesVisited, counter)
			break
		}
	}
	fmt.Println(len(housesVisited))
}

func (m SharedType) Day03_B(testFile string) {
	str := ReturnFileContentByChar(testFile)

	housesVisited := [][]int{{0, 0}}
	santaCounter := []int{0, 0}
	robotCounter := []int{0, 0}

	for i, v := range str {
		switch v {
		case 'v':
			if i%2 == 0 {
				santaCounter = []int{santaCounter[0] - 1, santaCounter[1]}
				housesVisited = AddHouseNotVisited(housesVisited, santaCounter)
			} else {
				robotCounter = []int{robotCounter[0] - 1, robotCounter[1]}
				housesVisited = AddHouseNotVisited(housesVisited, robotCounter)
			}
			break
		case '^':
			if i%2 == 0 {
				santaCounter = []int{santaCounter[0] + 1, santaCounter[1]}
				housesVisited = AddHouseNotVisited(housesVisited, santaCounter)
			} else {
				robotCounter = []int{robotCounter[0] + 1, robotCounter[1]}
				housesVisited = AddHouseNotVisited(housesVisited, robotCounter)
			}
			break
		case '<':
			if i%2 == 0 {
				santaCounter = []int{santaCounter[0], santaCounter[1] - 1}
				housesVisited = AddHouseNotVisited(housesVisited, santaCounter)
			} else {
				robotCounter = []int{robotCounter[0], robotCounter[1] - 1}
				housesVisited = AddHouseNotVisited(housesVisited, robotCounter)
			}
			break
		case '>':
			if i%2 == 0 {
				santaCounter = []int{santaCounter[0], santaCounter[1] + 1}
				housesVisited = AddHouseNotVisited(housesVisited, santaCounter)
			} else {
				robotCounter = []int{robotCounter[0], robotCounter[1] + 1}
				housesVisited = AddHouseNotVisited(housesVisited, robotCounter)
			}
			break
		}
	}
	fmt.Println(len(housesVisited))
}
