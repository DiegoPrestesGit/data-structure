package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("plzz send dudes (I mean the date as argument like this: yyyy-dd)")
		return
	}

	var myRegex = regexp.MustCompile(`^[0-9]{4}/[0-9]{2}$`)
	didMatch := myRegex.MatchString(os.Args[1])

	if !didMatch {
		fmt.Println("plzz use yyyy/dd format")
		return
	}

	date := strings.Split(os.Args[1], "/")
	path := "y"+date[0] + "/day" + date[1] + "/A.go"

	if _, err := os.Stat(path); err != nil {
		fmt.Println("day not found")
	}

	cmd := exec.Command("go", "run", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running the code", err)
	}
}
