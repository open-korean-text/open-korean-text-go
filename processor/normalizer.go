package processor

import (
	"log"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/open-korean-text/open-korean-text-go/dictionary"
	"github.com/open-korean-text/open-korean-text-go/hangul"
	"github.com/open-korean-text/open-korean-text-go/util"
)

const (
	extendedKoreanRegex    = "([ㄱ-ㅣ가-힣]+)"
	koreanToNormalizeRegex = "([가-힣]+)(ㅋ+|ㅎ+|[ㅠㅜ]+)"
	codaNException         = "[은|는|운|인|텐|근|른|픈|닌|든|던]"
)

func correctTypo(input string) string {
	temp := input
	typoMap := dic.GetTypoMap()

	for k, v := range typoMap {
		re := regexp.MustCompile(k)
		temp = re.ReplaceAllString(temp, v)
	}

	return temp
}

func normalizeCodaN(input string) string {
	inputLen := utf8.RuneCountInString(input)
	if inputLen < 2 {
		return input
	}

	lastTwo := util.Substr(input, inputLen-2, inputLen)
	last := util.GetCharStr(input, inputLen-1)

	lastTwoHead := util.GetCharStr(lastTwo, 0)

	nounDic, err := dic.GetNounDic()
	if err != nil {
		log.Fatal(err)
	}

	adverbDic, err := dic.GetAdverbDic()
	if err != nil {
		log.Fatal(err)
	}

	conjunctionDic, err := dic.GetConjunctionDic()
	if err != nil {
		log.Fatal(err)
	}

	lastTwoHeadR, _ := utf8.DecodeRuneInString(lastTwoHead)
	if isWordInDic(input, nounDic) ||
		isWordInDic(input, adverbDic) ||
		isWordInDic(input, conjunctionDic) ||
		isWordInDic(lastTwo, nounDic) ||
		lastTwoHeadR < '가' ||
		lastTwoHeadR > '힣' {
		return input
	}

	if re := regexp.MustCompile(codaNException); re.FindString(lastTwoHead) != "" {
		return input
	}

	hc := hangul.DecomposeHangul(lastTwoHead)
	newHead := util.Substr(input, 0, inputLen-2) +
		hangul.ComposeHangul(string(hc.Onset), string(hc.Vowel), string(' '))

	lastR, _ := utf8.DecodeRuneInString(last)
	if hc.Coda == 'ㄴ' &&
		(lastR == '데' || lastR == '가' || lastR == '지') &&
		isWordInDic(newHead, nounDic) {
		return newHead + "인" + last
	}

	return input
}

func normalizeRepeating(input string) string {
	result := input
	re := regexp.MustCompile("(..)")
	mGroup := re.FindAllString(result, -1)

	for _, v := range mGroup {

		pattern := "(" + regexp.QuoteMeta(v) + "){3,}"
		re2 := regexp.MustCompile(pattern)
		mGroup2 := re2.FindAllStringSubmatch(result, -1)

		for _, v2 := range mGroup2 {
			re3 := regexp.MustCompile(v2[0])
			result = re3.ReplaceAllString(result, strings.Repeat(v2[1], 2))
		}
	}

	return result
}

func removeRepeatingChar(input string) string {
	arr := strings.Split(input, "")

	result := ""
	buffer := ""
	bufferIdx := 0
	for _, v := range arr {
		if buffer != v {
			buffer = v
			bufferIdx = 1
			result += v
			continue
		}

		bufferIdx++

		if bufferIdx <= 3 {
			result += v
		}
	}

	return result
}

func isWordInDic(word string, dictionary []string) bool {
	for _, v := range dictionary {
		if word == v {
			return true
		}
	}

	return false
}

func secondToLastDecomposed(init string) (*hangul.HangulChar, bool) {
	var hc *hangul.HangulChar
	ok := false
	initLen := utf8.RuneCountInString(init)
	if initLen > 0 {
		c := util.GetCharStr(init, initLen-1)
		hc = hangul.DecomposeHangul(c)
		if hc.Coda == ' ' {
			ok = hangul.CheckHangulChar(hc)
		}
	}

	return hc, ok
}

func normalizeEmotionAttachedChunk(chunk, target string) string {
	chunkLen := utf8.RuneCountInString(chunk)
	init := util.Substr(chunk, 0, chunkLen-1)
	initLen := chunkLen - 1

	hc, ok := secondToLastDecomposed(init)

	c := util.GetCharStr(chunk, chunkLen-1)
	hc2 := hangul.DecomposeHangul(c)

	if hc2.Coda == 'ㅋ' || hc2.Coda == 'ㅎ' {
		return init + hangul.ComposeHangul(hc2.Onset, hc2.Vowel, ' ')
	}

	if ok &&
		string(hc2.Vowel) == string(target[0]) &&
		hangul.CheckCharInCodaMap(hc2.Onset) {
		return util.Substr(init, 0, initLen-1) + hangul.ComposeHangul(hc.Onset, hc.Vowel, hc2.Onset)
	}

	return chunk
}

func normalizeEnding(input string) string {
	nounDic, err := dic.GetNounDic()
	if err != nil {
		log.Fatal(err)
	}

	eomiDic, err := dic.GetEomiDic()
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(koreanToNormalizeRegex)
	result := re.ReplaceAllStringFunc(
		input,
		func(m string) string {
			mGroup := re.FindStringSubmatch(m)
			chunk := mGroup[1]
			chunkLen := utf8.RuneCountInString(chunk)
			chunkRight1 := util.Substr(chunk, chunkLen-1, chunkLen)
			chunkRight2 := util.Substr(chunk, chunkLen-2, chunkLen)
			target := mGroup[2]

			var normalizedChunk string
			flag := isWordInDic(chunk, nounDic) ||
				isWordInDic(chunkRight1, eomiDic) ||
				isWordInDic(chunkRight2, eomiDic)
			if flag {
				normalizedChunk = chunk
			} else {
				normalizedChunk = normalizeEmotionAttachedChunk(chunk, target)
			}

			return normalizedChunk + target
		},
	)

	return result
}

// Normalize : normalize Korean text
func Normalize(input string) string {
	re := regexp.MustCompile(extendedKoreanRegex)
	endingNormalized := re.ReplaceAllStringFunc(
		input,
		func(m string) string {
			return normalizeEnding(m)
		},
	)

	exclamationNormalized := removeRepeatingChar(endingNormalized)
	repeatingNormalized := normalizeRepeating(exclamationNormalized)
	codaNNormalized := normalizeCodaN(repeatingNormalized)
	typoCorrected := correctTypo(codaNNormalized)

	return typoCorrected
}
