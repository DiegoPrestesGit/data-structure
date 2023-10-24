package main

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/diegolikescode/adventofcode/y2015"
)

func main() {
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
	my := y2015.SharedType{}

	date := strings.Split(os.Args[1], "/")
	callIt := "Day" + string(date[1][:len(date[1])-1]) + "_" + string(date[1][len(date[1])-1])
	fmt.Println("CALLING THIS", callIt)

	inputTxtPath := "./y"+date[0]+"/"+"/input" + date[1][:len(date[1]) - 1] + ".txt"
	reflect.ValueOf(my).MethodByName(callIt).Call([]reflect.Value{reflect.ValueOf(inputTxtPath)})
}
