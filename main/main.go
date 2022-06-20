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

	isContinue := true

	for isContinue {

		fmt.Print("Enter text : ")
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\n", "", -1)

		if len(text) > 0 {
			TGACase.IsBalanced(text)
		}

	}

}
