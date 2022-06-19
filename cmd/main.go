package cmd

import (
	"TGACase"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)


	fmt.Print("Merhaba, lütfen kontrol edilmesini istediğiniz parantez kümesini giriniz : ")
	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\n", "", -1)


	TGACase.BracketsSymmetry("((})")

}
