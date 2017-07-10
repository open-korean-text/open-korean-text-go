package dic

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var basePath string = os.Getenv("GOPATH") + "/src/github.com/open-korean-text/open-korean-text-go/dictionary/dic/"

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

func getDic(dicPath string) ([]string, error) {
	file, err := getFile(basePath + dicPath)
	if err != nil {
		return nil, err
	}

	dic, err := getWordList(file)
	if err != nil {
		return nil, err
	}

	return dic, nil
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
		wordList, err := getDic(v)
		if err != nil {
			return nil, err
		}

		dic = append(dic, wordList...)
	}

	return dic, nil
}

// GetEomiDic : To get eomi list
func GetEomiDic() ([]string, error) {
	return getDic("verb/eomi.txt")
}

// GetConjunctionDic : To get conjunction list
func GetConjunctionDic() ([]string, error) {
	return getDic("auxiliary/conjunctions.txt")
}

// GetAdverbDic : To get adverb list
func GetAdverbDic() ([]string, error) {
	return getDic("adverb/adverb.txt")
}

// GetTypoMap : get typo map
func GetTypoMap() map[string]string {
	file, err := os.Open(basePath + "typos/typos.txt")
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
