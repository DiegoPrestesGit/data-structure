package y2015

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func (m SharedType) Day07_A(testFile string) {
	str := ReturnFileByLine(testFile)

	signals := make(map[string]uint16)
	gateCount := make(map[string]int)
	for i, v := range str {
		lineSplit := strings.Split(v, "->")
		operation := strings.Split(lineSplit[0][:len(lineSplit[0])-1], " ")
		variable := strings.Trim(lineSplit[1], " ")

		isStrReg := regexp.MustCompile("[a-zA-Z]")
		// fmt.Println(signals)
		switch(len(operation)) {
		case 1: //atribution
			gateCount["attribution"]++
			if isStrReg.MatchString(operation[0]) {
				fmt.Println("YES IT IS STRING", operation[0], signals[operation[0]])
				signals[variable] = signals[operation[0]]
			} else {
				fmt.Println("NO ITS NOT", operation[0])
				signals[variable] = ParseStr2uint16(operation[0])
			}
			break

		case 2: // NOT operation
			gateCount["not"]++
			signals[variable] = ^signals[operation[1]]
			fmt.Println("NOT", signals[operation[1]], ^signals[operation[1]], signals[variable])
		break

		case 3:
			if verifyValidInteger(operation[0]) || verifyValidInteger(operation[2]) {
				fmt.Println(i, "WE GOT SOME NUMBERS OVER HEEEEERE", operation[0], operation[2])
				if !verifyValidInteger(operation[0]) && signals[operation[0]] != 0 {
					fmt.Println("VERY SPECIAL, indeed", signals[operation[0]])
				}
			}

			if operation[1] == "AND" {
				gateCount["and"]++
				var opVal1 uint16 = internalAssertion(signals, operation[0], *isStrReg)
				var opVal2 uint16 = internalAssertion(signals, operation[1], *isStrReg)
				signals[variable] = opVal1 & opVal2
				fmt.Println(i, "AND", opVal1, opVal2, signals[variable])
			}
			if operation[1] == "OR" {
				gateCount["or"]++
				var opVal1 uint16 = signals[operation[0]]
				var opVal2 uint16 = signals[operation[2]]
				signals[variable] = opVal1 | opVal2
				fmt.Println(i, "OR", opVal1, opVal2, signals[variable])
			}
			if operation[1] == "LSHIFT" {
				gateCount["lshift"]++
				var opVal1 uint16 = signals[operation[0]]
				var opVal2 uint16 = ParseStr2uint16(operation[2])
				signals[variable] = opVal1 << opVal2
				fmt.Println(i, "LSHIFT", opVal1, opVal2, signals[variable])
			}
			if operation[1] == "RSHIFT" {
				gateCount["rshift"]++
				var opVal1 uint16 = signals[operation[0]]
				var opVal2 uint16 = ParseStr2uint16(operation[2])
				signals[variable] = opVal1 >> opVal2
				fmt.Println(i, "RSHIFT", opVal1, opVal2, signals[variable])
			}
		break

		default:
			fmt.Println("THE DEFAULT CASE")
		}
	}
	fmt.Println(gateCount)
	fmt.Println("value of A:", signals["a"])
}

func ParseStr2uint16 (str string) (uint16) {
	val, err := strconv.ParseInt(str, 0, 16); if err != nil {
		panic("AT THE DISCO (ParseStr2uint16)")
	}

	return uint16(val)
}

func internalAssertion (signals map[string]uint16, str string, regie regexp.Regexp) (uint16) {
	if regie.MatchString(str) { // it is a signal variable
		return signals[str]
	} else { // it is a number
		return ParseStr2uint16(str)
	}
}

func verifyValidInteger (n string) (bool) {
	_, er := strconv.Atoi(n); if er != nil {
		return false
	}
	return true
}


/*
func Integer2binary(n int) (string) {
	if n == 0 {
		return ""
	}

	return Integer2binary(n/2) + strconv.Itoa(n%2)
}

func Binary2integer(bin string, pos float64) (int) {
	if len(bin) == 0 {
		return 0
	}

	return Binary2integer(string(bin[:len(bin)-1]), pos+1) +
	(StrParseInt(string(bin[len(bin)-1])) * int(math.Pow(2, pos)))
}
*/
