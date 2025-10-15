package main

import (
	"bufio"
	"fmt"
	"os"
)

func getSecretWord(wordFileName string) string {
	word, err := os.Open(wordFileName)
	if err != nil {
		fmt.Errorf("Failed to open the file: %v", err)
	}
	defer word.Close()
	scanner := bufio.NewScanner(word)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return "hii"
}
func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}
