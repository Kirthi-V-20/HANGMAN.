package main

import (
	"fmt"
)

func getSecretWord(wordFile string) string {
	return "hii"
}

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}
