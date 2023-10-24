package y2015

import (
	"fmt"
	"strconv"
	"strings"
)

func (m SharedType) Day06_A (testFile string) {
	str := ReturnFileByLine(testFile)

	size := 1_000
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	for _, line := range str {
		// destruct the instructions
		splitedLine := strings.Split(line, " ")
		_, err := strconv.Atoi(string(splitedLine[1][0]))
		// the end will always be at the end of the instruction, so we get him here
		endArr := strings.Split(splitedLine[len(splitedLine)-1], ",")
		e := []int{StrParseInt(string(endArr[0])), StrParseInt(string(endArr[1]))}
		if err != nil {
			begin := strings.Split(splitedLine[2], ",")
			var b []int = []int{StrParseInt(begin[0]), StrParseInt(begin[1])}
			if splitedLine[1] == "on" {
				matrix = lightsOn(matrix, b, e)
			} else {
				matrix = lightsOff(matrix, b, e)
			}
		} else {
			b := []int{
				StrParseInt(string(strings.Split(splitedLine[1], ",")[0])),
				StrParseInt(string(strings.Split(splitedLine[1], ",")[1])),
			}
			matrix = toggle(matrix, b, e)
		}
	}
	counted := count(matrix)
	fmt.Println(counted)
}

func lightsOn (matrix [][]int, begin []int, end []int) ([][]int) {
	for i:=begin[0]; i <= end[0]; i++ {
		for j:=begin[1]; j <= end[1]; j++ {
			matrix[i][j] = 1
		}
	}
	return matrix
}

func lightsOff (matrix [][]int, begin []int, end []int) ([][]int) {
	for i:=begin[0]; i <= end[0]; i++ {
		for j:=begin[1]; j <= end[1]; j++ {
			matrix[i][j] = 0
		}
	}
	return matrix
}

func toggle (matrix [][]int, begin []int, end []int) ([][]int) {
	for i:=begin[0]; i <= end[0]; i++ {
		for j:=begin[1]; j <= end[1]; j++ {
			if matrix[i][j] == 1 {
				matrix[i][j] = 0
			} else {
				matrix[i][j] = 1
			}
		}
	}
	return matrix
}

func count (matrix [][]int) (int) {
	count := 0
	for rowIdx, row := range matrix {
		for colIdx := range row {
			if matrix[rowIdx][colIdx] == 1 {
				count++
			}
		}
	}
	return count
}


func (m SharedType) Day06_B (testFile string) {
	str := ReturnFileByLine(testFile)

	size := 1_000
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	for _, line := range str {
		// destruct the instructions
		splitedLine := strings.Split(line, " ")
		_, err := strconv.Atoi(string(splitedLine[1][0]))
		// the end will always be at the end of the instruction, so we get him here
		endArr := strings.Split(splitedLine[len(splitedLine)-1], ",")
		e := []int{StrParseInt(string(endArr[0])), StrParseInt(string(endArr[1]))}
		if err != nil {
			begin := strings.Split(splitedLine[2], ",")
			var b []int = []int{StrParseInt(begin[0]), StrParseInt(begin[1])}
			if splitedLine[1] == "on" {
				matrix = lightsOn_B(matrix, b, e)
			} else {
				matrix = lightsOff_B(matrix, b, e)
			}
		} else {
			b := []int{
				StrParseInt(string(strings.Split(splitedLine[1], ",")[0])),
				StrParseInt(string(strings.Split(splitedLine[1], ",")[1])),
			}
			matrix = toggle_B(matrix, b, e)
		}
	}
	counted := countBrightness(matrix)
	fmt.Println(counted)
}

func lightsOn_B (matrix [][]int, begin []int, end []int) ([][]int) {
	for i:=begin[0]; i <= end[0]; i++ {
		for j:=begin[1]; j <= end[1]; j++ {
			matrix[i][j] += 1
		}
	}
	return matrix
}

func lightsOff_B (matrix [][]int, begin []int, end []int) ([][]int) {
	for i:=begin[0]; i <= end[0]; i++ {
		for j:=begin[1]; j <= end[1]; j++ {
			if matrix[i][j] > 0 {
				matrix[i][j] -= 1
			}
		}
	}
	return matrix
}

func toggle_B (matrix [][]int, begin []int, end []int) ([][]int) {
	for i:=begin[0]; i <= end[0]; i++ {
		for j:=begin[1]; j <= end[1]; j++ {
			matrix[i][j] += 2
		}
	}
	return matrix
}

func countBrightness (matrix [][]int) (int) {
	count := 0
	for rowIdx, row := range matrix {
		for colIdx := range row {
			count += matrix[rowIdx][colIdx]
		}
	}
	return count
}

