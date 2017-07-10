package hangul

import "unicode/utf8"

// HangulChar : hangul character includes onset, vowel, coda
type HangulChar struct {
	Onset rune
	Vowel rune
	Coda  rune
}

type doubleCoda struct {
	first  rune
	second rune
}

const (
	hangulBase = 0xAC00
	onsetBase  = 21 * 28
	vowelBase  = 28
)

var onsetList = [...]rune{
	'ㄱ', 'ㄲ', 'ㄴ', 'ㄷ', 'ㄸ', 'ㄹ', 'ㅁ', 'ㅂ', 'ㅃ',
	'ㅅ', 'ㅆ', 'ㅇ', 'ㅈ', 'ㅉ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ',
}

var vowelList = [...]rune{
	'ㅏ', 'ㅐ', 'ㅑ', 'ㅒ', 'ㅓ', 'ㅔ',
	'ㅕ', 'ㅖ', 'ㅗ', 'ㅘ', 'ㅙ', 'ㅚ',
	'ㅛ', 'ㅜ', 'ㅝ', 'ㅞ', 'ㅟ', 'ㅠ',
	'ㅡ', 'ㅢ', 'ㅣ',
}

var codaList = [...]rune{
	' ', 'ㄱ', 'ㄲ', 'ㄳ', 'ㄴ', 'ㄵ', 'ㄶ', 'ㄷ',
	'ㄹ', 'ㄺ', 'ㄻ', 'ㄼ', 'ㄽ', 'ㄾ', 'ㄿ', 'ㅀ',
	'ㅁ', 'ㅂ', 'ㅄ', 'ㅅ', 'ㅆ', 'ㅇ', 'ㅈ', 'ㅊ',
	'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ',
}

var onsetMap, vowelMap, codaMap map[rune]int

var doubleCodas = map[rune]doubleCoda{
	'ㄳ': doubleCoda{'ㄱ', 'ㅅ'},
	'ㄵ': doubleCoda{'ㄴ', 'ㅈ'},
	'ㄶ': doubleCoda{'ㄴ', 'ㅎ'},
	'ㄺ': doubleCoda{'ㄹ', 'ㄱ'},
	'ㄻ': doubleCoda{'ㄹ', 'ㅁ'},
	'ㄼ': doubleCoda{'ㄹ', 'ㅂ'},
	'ㄽ': doubleCoda{'ㄹ', 'ㅅ'},
	'ㄾ': doubleCoda{'ㄹ', 'ㅌ'},
	'ㄿ': doubleCoda{'ㄹ', 'ㅍ'},
	'ㅀ': doubleCoda{'ㄹ', 'ㅎ'},
	'ㅄ': doubleCoda{'ㅂ', 'ㅅ'},
}

func init() {
	onsetMap = make(map[rune]int)
	vowelMap = make(map[rune]int)
	codaMap = make(map[rune]int)

	for i, v := range onsetList {
		onsetMap[v] = i
	}

	for i, v := range vowelList {
		vowelMap[v] = i
	}

	for i, v := range codaList {
		codaMap[v] = i
	}
}

// DecomposeHangul : decompose hangul char
func DecomposeHangul(c string) *HangulChar {
	r, _ := utf8.DecodeRuneInString(c)

	_, ok1 := onsetMap[r]
	_, ok2 := vowelMap[r]
	_, ok3 := codaMap[r]

	if ok1 || ok2 || ok3 {
		return nil
	}

	u := r - hangulBase
	return &HangulChar{
		onsetList[u/onsetBase],
		vowelList[(u%onsetBase)/vowelBase],
		codaList[u%vowelBase],
	}
}

// ComposeHangul : compose hangul char
func ComposeHangul(onset, vowel, coda string) string {
	onsetR, _ := utf8.DecodeRuneInString(onset)
	vowelR, _ := utf8.DecodeRuneInString(vowel)
	codaR, _ := utf8.DecodeRuneInString(coda)

	if onsetR == ' ' || vowelR == ' ' {
		return ""
	}

	return string(hangulBase +
		(onsetMap[onsetR] * onsetBase) +
		(vowelMap[onsetR] * vowelBase) +
		codaMap[codaR])
}

// ComposeHangulChar : compose hangul char
func ComposeHangulChar(hc *HangulChar) string {
	return ComposeHangul(
		string(hc.Onset),
		string(hc.Vowel),
		string(hc.Coda))
}

// CheckHangulChar : Is HangulChar empty?
func CheckHangulChar(hc *HangulChar) bool {
	return utf8.ValidRune(hc.Onset) &&
		utf8.ValidRune(hc.Vowel) &&
		utf8.ValidRune(hc.Coda)
}

// CheckCharInCodaMap : Is char in codaMap?
func CheckCharInCodaMap(coda rune) bool {
	_, ok := codaMap[coda]
	return ok
}
