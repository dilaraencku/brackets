package TGACase

import (
	"fmt"
	"reflect"
)

func IsBalanced(myString string) string {

	acilisParantezleri := []string{"[", "(", "{"}
	kapanisParantezleri := []string{"]", ")", "}"}

	fmt.Println("myString :", myString)

	var controlArray []string
	var kapanisParanteziSayaci int

	for i := 0; i <= len(myString)-1; i++ {

		responseForAcilisParantezleri, _ := isCharExist(acilisParantezleri, string(myString[i]))
		responseForKapanisParantezleri, _ := isCharExist(kapanisParantezleri, string(myString[i]))


		if responseForAcilisParantezleri == true {

			kapanisParanteziSayaci = 0

			controlArray = append(controlArray, string(myString[i]))

		} else if responseForKapanisParantezleri == true {

			kapanisParanteziSayaci++

			myChar := string(myString[i])

			if myChar == ")" {

				if controlArray[len(controlArray)-kapanisParanteziSayaci] == "(" {

				} else {

					fmt.Println("NO")
					return "NO"

				}

			} else if myChar == "}" {

				if controlArray[len(controlArray)-kapanisParanteziSayaci] == "{" {

				} else {

					fmt.Println("NO")
					return "NO"

				}

			} else if myChar == "]" {

				if controlArray[len(controlArray)-kapanisParanteziSayaci] == "[" {

				} else {

					fmt.Println("NO")
					return "NO"

				}

			}

		} else {
			return ("PLEASE ENTER VALID STRING, EXPECTED STRING SHOULD CONTAINS ONLY (,[,{,),},] characters")
		}

	}

	fmt.Println("YES")
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
