package TGACase

import (
	"fmt"
	"reflect"
)

func IsBalanced(myString string) string {

	openingBrackets := []string{"[", "(", "{"}
	closingBrackets := []string{"]", ")", "}"}

	var controlArray []string
	var closingBracketsFlag int

	for i := 0; i <= len(myString)-1; i++ {

		responseForOpeningBrackets, _ := isCharExist(openingBrackets, string(myString[i]))
		responseForClosingBrackets, _ := isCharExist(closingBrackets, string(myString[i]))


		if responseForOpeningBrackets == true {

			closingBracketsFlag = 0

			controlArray = append(controlArray, string(myString[i]))

		} else if responseForClosingBrackets == true {

			closingBracketsFlag++

			myChar := string(myString[i])

			if myChar == ")" {

				if controlArray[len(controlArray)-closingBracketsFlag] == "(" {

				} else {

					return "NO"

				}

			} else if myChar == "}" {

				if controlArray[len(controlArray)-closingBracketsFlag] == "{" {

				} else {

					return "NO"

				}

			} else if myChar == "]" {

				if controlArray[len(controlArray)-closingBracketsFlag] == "[" {

				} else {

					return "NO"

				}

			}

		} else {

			fmt.Println("sdadsadsa")
			return "NO"
		}

	}

	return "YES"

}

func isCharExist(brackets []string, char string) (bool, int) {
	s := reflect.ValueOf(brackets)

	if s.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}

	for i := range brackets {
		if string(brackets[i]) == char {
			return true, i
		}
	}

	return false, 0
}
