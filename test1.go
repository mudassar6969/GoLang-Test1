package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

func findWordsCount(str string) int {

	if len(str) == 0 {
		return 0
	}
	count := 1
	for _, r := range str {

		if unicode.IsUpper(r) {
			count++
		}
	}

	return count
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func filesExercise(fileName, textToWrite string) {

	file, fileCreateError := os.Create(fileName)
	checkError(fileCreateError)

	defer file.Close()

	_, writeError := file.WriteString(textToWrite)
	checkError(writeError)

	readFile(fileName)
}

func readFile(fileName string) {
	data, readError := ioutil.ReadFile(fileName)
	checkError(readError)
	fmt.Print(string(data))
}

type TokenData struct {
	Href string
	Text string
}

func parseFile(fileName string) []TokenData {

	data, readError := ioutil.ReadFile(fileName)
	checkError(readError)

	var tokensData []TokenData

	reader := strings.NewReader(string(data))

	tokenizer := html.NewTokenizer(reader)

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			break
		}

		if tokenType == html.StartTagToken {
			startTagB, _ := tokenizer.TagName()

			if string(startTagB) == "a" {

				_, tagValueB, _ := tokenizer.TagAttr()
				tokenData := TokenData{}

				tokenData.Href = string(tagValueB)
				tokenizer.Next()
				tagB, _ := tokenizer.TagName()

				for string(startTagB) != string(tagB) {
					if tokenType != html.CommentToken {
						tokenData.Href += tokenizer.Token().Data
					}
					tokenType = tokenizer.Next()
					tagB, _ = tokenizer.TagName()
				}

				tokensData = append(tokensData, tokenData)
			}
		}
	}

	return tokensData

}

func main() {
	// wordsCount := findWordsCount("saveChangesInTheEditor")
	// fmt.Println("Words Count: ", wordsCount)

	//filesExercise("sample.txt", "sample text to write")

	fmt.Println(parseFile("ex4.html"))
}
