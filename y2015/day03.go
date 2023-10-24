package y2015

import (
	"fmt"
	"reflect"
)

func AddHouseNotVisited (housesVisited [][]int, position []int) ([][]int) {
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
		switch(v) {
		case 'v':
			counter = []int{counter[0]-1, counter[1]}
			housesVisited = AddHouseNotVisited(housesVisited, counter)
			break
		case '^':
			counter = []int{counter[0]+1, counter[1]}
			housesVisited = AddHouseNotVisited(housesVisited, counter)
			break
		case '<':
			counter = []int{counter[0], counter[1]-1}
			housesVisited = AddHouseNotVisited(housesVisited, counter)
			break
		case '>':
			counter = []int{counter[0], counter[1]+1}
			housesVisited = AddHouseNotVisited(housesVisited, counter)
			break
		}
	}
	fmt.Println(len(housesVisited))
}

func (m SharedType) Day03_B(testFile string) {

}

