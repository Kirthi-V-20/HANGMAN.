package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func getSecretWord(wordFileName string) string {
	allowedWords := []string{}
	wordFile, err := os.Open(wordFileName)
	if err != nil {
		fmt.Errorf("Failed to open the file: %v", err)
	}
	defer wordFile.Close()
	scanner := bufio.NewScanner(wordFile)
	for scanner.Scan() {
		word := scanner.Text()

		if word == strings.ToLower(word) {
			allowedWords = append(allowedWords, word)
		}
	}
	randNum := rand.Intn(len(allowedWords))
	return allowedWords[randNum]

}
func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}
