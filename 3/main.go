package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(findFirstStringInBracket("Hello (new) (world)"))
	fmt.Println(findFirstStringInBracket("Hello ((new) (world))"))
}

func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		var regex, err = regexp.Compile(`([(].+?[)])`)
		if err != nil {
			return ""
		}

		var res string = regex.FindAllString(str, 1)[0]
		return res[1 : len(res)-1]
	}
	return ""
}
