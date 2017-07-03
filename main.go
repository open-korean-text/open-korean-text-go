package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func getNewFile(fileName string) (*os.File, error) {
	file, err := os.OpenFile(
		fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(0644))
	if err != nil {
		return nil, err
	}

	return file, nil
}

func getAdditionalWords(file *os.File) ([]string, error) {
	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}

		result = append(result, strings.TrimSpace(s))
	}

	// If there are no additional words, close this programe
	if len(result) == 0 {
		return nil, fmt.Errorf(file.Name(), " is empty.")
	}

	sort.Sort(sort.StringSlice(result))
	return result, nil
}

func getFile() (*os.File, error) {
	var fileName string
	fmt.Scan(&fileName)
	if strings.Index(fileName, ".txt") == -1 {
		return nil, fmt.Errorf("Extension of file must be 'txt'")
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return file, nil
}

func getInputFile() (*os.File, *os.File) {
	var file1, file2 *os.File

	fmt.Println("Current dictionary file.")
	for {
		if file, err := getFile(); err != nil {
			fmt.Println(err)
			fmt.Println("Input again.")
			continue
		} else {
			fmt.Println("Dictionary :", file.Name())
			fmt.Println()
			file1 = file
			break
		}
	}

	fmt.Println("Additional words.(txt file)")
	for {
		if file, err := getFile(); err != nil {
			fmt.Println(err)
			fmt.Println("Input again.")
			continue
		} else {
			fmt.Println("Additional words :", file.Name())
			fmt.Println()
			file2 = file
			break
		}
	}

	return file1, file2
}

func main() {
	// Get Input from user
	dicFile, newFile := getInputFile()
	defer dicFile.Close()
	defer newFile.Close()

	// Open a file that includes additional words and get an array of these words
	additionalWords, err := getAdditionalWords(newFile)
	if err != nil {
		log.Fatal(err)
	}

	// Open and create new dic file
	newDicFile, err := getNewFile("result.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newDicFile.Close()

	scanner := bufio.NewScanner(dicFile)
	for scanner.Scan() {
		s := scanner.Text()

		// If there are no additional words, it writes words in dic file
		if len(additionalWords) == 0 {
			if _, err := newDicFile.Write([]byte(s + "\r\n")); err != nil {
				log.Fatal(err)
			}

			continue
		}

		for len(additionalWords) > 0 {
			if word := additionalWords[0]; strings.Compare(word, s) == -1 {
				// Add a additional word and remove the word in array
				if _, err := newDicFile.Write([]byte(word + "\r\n")); err != nil {
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

		if _, err := newDicFile.Write([]byte(s + "\r\n")); err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("This work is done successfully! ^^")
}
