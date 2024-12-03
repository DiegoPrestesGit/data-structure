package y2015

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func (m SharedType) Day07_A(testFile string) {
	str := ReturnFileByLine(testFile)

	fmt.Println("original size", len(str))
	signals := make(map[string]uint16)
	recursiveStuff(str, signals)
}

func recursiveStuff(remainStr []string, signals map[string]uint16) {
	fmt.Println("REMAINING", len(remainStr))
	if len(remainStr) == 0 {
		fmt.Println("DONE!")
		fmt.Println(signals)
		return
	}

	for i, v := range remainStr {
		// toda vez que uma instrucao for executada, "recursiveStuff" deve recomecar de forma recursiva
		// e sem a instrucao que acabou de ser execuada, a condicao de parada é "remainStr" ter tamanho zero
		lineSplit := strings.Split(v, "->")
		operation := strings.Split(lineSplit[0][:len(lineSplit[0])-1], " ")
		variable := strings.Trim(lineSplit[1], " ")
		isStrReg := regexp.MustCompile("[a-zA-Z]")

		switch len(operation) {
		// atribuicao
		case 1:
			if isStrReg.MatchString(operation[0]) {
				if !checkKeyExists(signals, operation[0]) {
					continue
				} else {
					signals[variable] = signals[operation[0]]
					newStr := append(remainStr[:i], remainStr[i+1:]...)
					recursiveStuff(newStr, signals)
					fmt.Println(len(remainStr), "|", i, "|", v)
					return
				}
			} else {
				signals[variable] = ParseStr2uint16(operation[0])
				newStr := append(remainStr[:i], remainStr[i+1:]...)
				recursiveStuff(newStr, signals)
				fmt.Println(len(remainStr), "|", i, "|", v)
				return
			}
		// not operation
		case 2:
			if isStrReg.MatchString(operation[1]) && !checkKeyExists(signals, operation[1]) {
				continue
			}
			signals[variable] = ^signals[operation[1]]
			newStr := append(remainStr[:i], remainStr[i+1:]...)
			recursiveStuff(newStr, signals)
			fmt.Println(len(remainStr), "|", i, "|", v)
			return

		case 3:
			if operation[1] == "AND" {
				if (!checkKeyExists(signals, operation[0]) && isStrReg.MatchString(operation[0])) ||
					(!checkKeyExists(signals, operation[2]) && isStrReg.MatchString(operation[2])) {
					continue
				}

				var opVal1 uint16 = internalAssertion(signals, operation[0], *isStrReg)
				var opVal2 uint16 = internalAssertion(signals, operation[2], *isStrReg)
				signals[variable] = opVal1 & opVal2
				newStr := append(remainStr[:i], remainStr[i+1:]...)
				recursiveStuff(newStr, signals)
				fmt.Println(len(remainStr), "|", i, "|", v)
				return
			}
			if operation[1] == "OR" {
				if (!checkKeyExists(signals, operation[0]) && isStrReg.MatchString(operation[0])) ||
					(!checkKeyExists(signals, operation[2]) && isStrReg.MatchString(operation[2])) {
					continue
				}

				var opVal1 uint16 = signals[operation[0]]
				var opVal2 uint16 = signals[operation[2]]
				signals[variable] = opVal1 | opVal2
				newStr := append(remainStr[:i], remainStr[i+1:]...)
				recursiveStuff(newStr, signals)
				fmt.Println(len(remainStr), "|", i, "|", v)
				return
			}
			if operation[1] == "LSHIFT" {
				if !checkKeyExists(signals, operation[0]) && isStrReg.MatchString(operation[0]) {
					continue
				}

				var opVal1 uint16 = signals[operation[0]]
				var opVal2 uint16 = ParseStr2uint16(operation[2])
				signals[variable] = opVal1 << opVal2
				newStr := append(remainStr[:i], remainStr[i+1:]...)
				recursiveStuff(newStr, signals)
				fmt.Println(len(remainStr), "|", i, "|", v)
				return
			}
			if operation[1] == "RSHIFT" {
				if !checkKeyExists(signals, operation[0]) && isStrReg.MatchString(operation[0]) {
					continue
				}

				var opVal1 uint16 = signals[operation[0]]
				var opVal2 uint16 = ParseStr2uint16(operation[2])
				signals[variable] = opVal1 >> opVal2
				newStr := append(remainStr[:i], remainStr[i+1:]...)
				recursiveStuff(newStr, signals)
				fmt.Println(len(remainStr), "|", i, "|", v)
				return
			}
			break

		default:
			fmt.Println("THE DEFAULT CASE")
			break
		}
	}
	fmt.Println(signals)
}

func checkKeyExists(signals map[string]uint16, key string) bool {
	for k := range signals {
		if k == key {
			return true
		}
	}

	return false
}

func ParseStr2uint16(str string) uint16 {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic("AT THE DISCO (ParseStr2uint16)")
	}

	return uint16(val)
}

func internalAssertion(signals map[string]uint16, str string, regie regexp.Regexp) uint16 {
	if regie.MatchString(str) { // it is a signal variable
		return signals[str]
	} else { // it is a number
		return ParseStr2uint16(str)
	}
}

func verifyValidInteger(n string) bool {
	_, er := strconv.Atoi(n)
	if er != nil {
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
