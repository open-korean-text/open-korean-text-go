package dic

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func getWordList(file *os.File) ([]string, error) {
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

func getFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return file, nil
}

// GetNounDic : To get noun list
func GetNounDic() ([]string, error) {
	fileList := [...]string{
		"noun/nouns.txt", "noun/entities.txt", "noun/spam.txt",
		"noun/names.txt", "noun/twitter.txt", "noun/lol.txt",
		"noun/slangs.txt", "noun/company_names.txt",
		"noun/foreign.txt", "noun/geolocations.txt", "noun/profane.txt",
		"substantives/given_names.txt", "noun/kpop.txt", "noun/bible.txt",
		"noun/pokemon.txt", "noun/congress.txt", "noun/wikipedia_title_nouns.txt",
	}

	var dic []string

	for _, v := range fileList {
		file, err := getFile("dic/" + v)
		if err != nil {
			return nil, err
		}

		wordList, err := getWordList(file)
		if err != nil {
			return nil, err
		}

		dic = append(dic, wordList...)
	}

	return dic, nil
}

// GetEomiDic : To get eomi list
func GetEomiDic() ([]string, error) {
	file, err := getFile("dic/verb/eomi.txt")
	if err != nil {
		return nil, err
	}

	dic, err := getWordList(file)
	if err != nil {
		return nil, err
	}

	return dic, nil
}

// GetConjunctionDic : To get conjunction list
func GetConjunctionDic() ([]string, error) {
	file, err := getFile("dic/auxiliary/conjunctions.txt")
	if err != nil {
		return nil, err
	}

	dic, err := getWordList(file)
	if err != nil {
		return nil, err
	}

	return dic, nil
}

// GetAdverbDic : To get adverb list
func GetAdverbDic() ([]string, error) {
	file, err := getFile("dic/adverb/adverb.txt")
	if err != nil {
		return nil, err
	}

	dic, err := getWordList(file)
	if err != nil {
		return nil, err
	}

	return dic, nil
}

// GetTypoMap : get typo map
func GetTypoMap() map[string]string {
	file, err := os.Open("dic/typos/typos.txt")
	if err != nil {
		log.Fatal(err)
	}

	typoMap := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		sGroup := strings.Split(s, " ")
		typoMap[sGroup[0]] = sGroup[1]
	}

	return typoMap
}
