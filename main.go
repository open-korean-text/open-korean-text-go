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

func getFileName() (string, error) {
	var file string
	fmt.Scan(&file)
	if strings.Index(file, ".txt") == -1 {
		return "", fmt.Errorf("Extension of file must be 'txt'.")
	}

	return file, nil
}

func getInput() (string, string, error) {
	var file1, file2 string

	fmt.Println("Current dictionary file.")
	for {
		file, err := getFileName()
		if err != nil {
			fmt.Println(err, "Input again.")
			continue
		} else {
			fmt.Println("Dictionary :", file)
			fmt.Println()
			file1 = file
			break
		}
	}

	fmt.Println("Additional words.(txt file)")
	for {
		file, err := getFileName()
		if err != nil {
			fmt.Println(err, "Input again.")
			continue
		} else {
			fmt.Println("Additional words :", file)
			fmt.Println()
			file2 = file
			break
		}
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
