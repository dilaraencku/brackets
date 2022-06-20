package main

import (
	"TGACase"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	var cevap string


	for cevap != "E"  || cevap !=  "e" {

		fmt.Print("Enter text : ")
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\n", "", -1)
		TGACase.IsBalanced(text)

		fmt.Print("if you want to exit press E, if you don't want any key is enough ")
		cevap, _ = reader.ReadString('\n')

	}




}
