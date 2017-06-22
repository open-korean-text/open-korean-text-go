package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func getAdditionalWords(file string) ([]string, error) {
	additionalFile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer additionalFile.Close()

	var result []string
	scanner := bufio.NewScanner(additionalFile)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}

		result = append(result, strings.TrimSpace(s))
	}

	sort.Sort(sort.StringSlice(result))
	return result, nil
}

func getInput() (string, string, error) {
	fmt.Println("Please input like this '[dictionary.txt] [additionalWords.txt]'.")

	var file1, file2 string
	fmt.Scan(&file1, &file2)
	fmt.Println("dictionary :", file1)
	fmt.Println("additional words :", file2)

	if strings.Index(file1, ".txt") == -1 {
		return "", "", fmt.Errorf("Extension of dictionary file is 'txt'")
	}

	if strings.Index(file2, ".txt") == -1 {
		return "", "", fmt.Errorf("Extension of additional words file is 'txt'")
	}

	return file1, file2, nil
}

func main() {
	// Get Input from user
	file1, file2, err := getInput()
	if err != nil {
		log.Fatal(err)
	}

	// Open dic file
	dicFile, err := os.Open(file1)
	if err != nil {
		log.Fatal(err)
	}
	defer dicFile.Close()

	// Open a file that includes additional words and get an array of these words
	additionalWords, err := getAdditionalWords(file2)
	if err != nil {
		log.Fatal(err)
	}

	// Open and create new dic file
	newDicFile, err := os.OpenFile(
		"result.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(0644))
	if err != nil {
		log.Fatal(err)
	}
	defer newDicFile.Close()

	scanner := bufio.NewScanner(dicFile)
	for scanner.Scan() {
		s := scanner.Text()

		// If there are no additional words, it writes words in dic file
		if len(additionalWords) == 0 {
			_, err = newDicFile.Write([]byte(s + "\r\n"))
			if err != nil {
				log.Fatal(err)
			}

			continue
		}

		for len(additionalWords) > 0 {
			word := additionalWords[0]
			if strings.Compare(word, s) == -1 {
				// Add a additional word and remove the word in array
				_, err = newDicFile.Write([]byte(word + "\r\n"))
				if err != nil {
					log.Fatal(err)
				}
				additionalWords = additionalWords[1:]
			} else if strings.Compare(word, s) == 0 {
				// If each words are same, it removes a additional word in array and gets out of the loop
				additionalWords = additionalWords[1:]
				break
			} else {
				// If it's not time to add a additional word, it gets out of the loop
				break
			}
		}

		_, err = newDicFile.Write([]byte(s + "\r\n"))
		if err != nil {
			log.Fatal(err)
		}
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("This work is done successfully! ^^")
}
