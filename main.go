package main

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"

	y2015 "github.com/diegolikescode/adventofcode/y2015/day01"
)

func main() {
	my := y2015.My2015{}
	if len(os.Args) != 2 {
		fmt.Println("plzz send dudes (I mean the date as argument like this: yyyy-dd)")
		return
	}

	var myRegex = regexp.MustCompile(`^[0-9]{4}/[0-9]{2}[A-B]{1}$`)
	didMatch := myRegex.MatchString(os.Args[1])
	if !didMatch {
		fmt.Println("plzz use yyyy/dd[A-B] format")
		return
	}

	date := strings.Split(os.Args[1], "/")

	callIt := "Day" + string(date[1][:len(date[1])-1]) + "_" + string(date[1][len(date[1])-1])

	fmt.Println("CALLING THIS", callIt)
	inputTxtPath := "./y"+date[0]+"/"+"day"+date[1][:len(date[1]) - 1]+"/input.txt"
	reflect.ValueOf(my).MethodByName(callIt).Call([]reflect.Value{reflect.ValueOf(inputTxtPath)})
}
