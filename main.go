package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
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

	for {
		fmt.Scan(&fileName)
		if strings.Index(fileName, ".txt") == -1 {
			fmt.Println("Extension of file must be 'txt'")
			fmt.Println("Input again.")
			continue
		} else {
			break
		}
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return file, nil
}

func getInputFile() (*os.File, *os.File, error) {
	fmt.Println("Current dictionary file.")
	file1, err := getFile()
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("Additional words.(txt file)")
	file2, err := getFile()
	if err != nil {
		return nil, nil, err
	}

	return file1, file2, nil
}

func createNewDic() {
	// Get Input from user
	dicFile, newFile, err := getInputFile()
	if err != nil {
		log.Fatal(err)
	}
	defer dicFile.Close()
	defer newFile.Close()

	// Open a file that includes additional words and get an array of these words
	additionalWords, err := getAdditionalWords(newFile)
	if err != nil {
		log.Fatal(err)
	}

	// Open and create new dic file
	newDicFile, err := getNewFile("new_" + filepath.Base(dicFile.Name()))
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

func getSelection() (int, error) {
	fmt.Println("Select menu number.")
	fmt.Println("1. Create new dictionary file")
	fmt.Println("2. Exit")

	var number string
	fmt.Scan(&number)

	return strconv.Atoi(number)
}

func main() {
	for {
		// Select menu
		number, err := getSelection()
		if err != nil {
			log.Fatal(err)
		}

		switch number {
		case 1:
			createNewDic()
		case 2:
			fmt.Println("Have a nice day! Bye~")
			return
		}

		fmt.Println()
	}
}
