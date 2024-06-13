package zerowidth

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// 크기가 없는 문자
var zeroWidthChars = []rune{
	0x000C, // FORM FEED (FF) \f
	0x00AD, // Soft Hyphen (SHY)
	0x034F, // Combining Grapheme Joiner (CGJ)
	0x200B, // Zero Width Space (ZWSP)
	0x200C, // Zero Width Non-Joiner (ZWNJ)
	0x200D, // Zero Width Joiner (ZWJ)
	0x200E, // Left-to-Right Mark (LRM)
	0x200F, // Right-to-Left Mark (RLM)
	0x202A, // Left-to-Right Embedding (LRE)
	0x202B, // Right-to-Left Embedding (RLE)
	0x202C, // Pop Directional Formatting (PDF)
	0x202D, // Left-to-Right Override (LRO)
	0x202E, // Right-to-Left Override (RLO)
	0x2060, // Word Joiner (WJ)
	0x2066, // Left-to-Right Isolate (LRI)
	0x2067, // Right-to-Left Isolate (RLI)
	0x2068, // First Strong Isolate (FSI)
	0x2069, // Pop Directional Isolate (PDI)
	0xFEFF, // Zero Width No-Break Space (BOM)
}

// 공백 문자
var spaceChars = []rune{
	0x0009, // Horizontal Tab (HT) \t
	0x0020, // (Space) - 일반 공백 문자
	0x00A0, // (No-Break Space) - 줄바꿈이 일어나지 않는 공백
	0x007f, // DELETE
	0x1680, // (Ogham Space Mark) - 오검 문자 사이의 공백
	0x2000, // (En Quad) - En 간격 공백
	0x2001, // (Em Quad) - Em 간격 공백
	0x2002, // (En Space) - En 간격 공백
	0x2003, // (Em Space) - Em 간격 공백
	0x2004, // (Three-Per-Em Space) - 1/3 Em 간격 공백
	0x2005, // (Four-Per-Em Space) - 1/4 Em 간격 공백
	0x2006, // (Six-Per-Em Space) - 1/6 Em 간격 공백
	0x2007, // (Figure Space) - 숫자 공백
	0x2008, // (Punctuation Space) - 구두점 공백
	0x2009, // (Thin Space) - 얇은 공백
	0x200A, // (Hair Space) - 매우 얇은 공백
	0x200B, // (Zero Width Space) - 너비가 없는 공백
	0x200C, // (Zero Width Non-Joiner) - 너비가 없는 비연결 공백
	0x200D, // (Zero Width Joiner) - 너비가 없는 연결 공백
	0x2028, // Line Separator (LS)
	0x2029, // Paragraph Separator (PS)
	0x202F, // (Narrow No-Break Space) - 좁은 비줄바꿈 공백
	0x205F, // (Medium Mathematical Space) - 중간 수학 공백
	0x2800, // Braille Pattern Blank
	0x3000, // (Ideographic Space) - 전각 공백
	0x115F, // 한글 초성 채움 문자
	0x1160, // 한글 중성 채움 문자
	0x3164, // 한글 채움 문자
}

// VARIATION SELECTOR
var varSelectors []rune

type ZeroWidth struct{}

type FindChar struct {
	Position int
	Char     rune
}

func NewZeroWidth() *ZeroWidth {
	return &ZeroWidth{}
}

func find(text string, chars []rune) ([]FindChar, error) {
	var runePatterns []string
	for _, r := range chars {
		runePatterns = append(runePatterns, regexp.QuoteMeta(string(r)))
	}
	pattern := strings.Join(runePatterns, "|")
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, nil
	}

	var findChars []FindChar
	matchIndexes := re.FindAllStringIndex(text, -1)
	for _, matchIndex := range matchIndexes {
		startRune := utf8.RuneCountInString(text[:matchIndex[0]])
		char := []rune(text[matchIndex[0]:matchIndex[1]])
		findChars = append(findChars, FindChar{startRune, char[0]})
	}
	return findChars, nil
}

func remove(text string, chars []rune) (string, error) {
	findChars, err := find(text, chars)
	if err != nil {
		return "", err
	}

	runes := []rune(text)
	var resultRunes []rune
	posIndex := 0
	for i, char := range runes {
		if posIndex < len(findChars) && i == findChars[posIndex].Position {
			posIndex++
		} else {
			resultRunes = append(resultRunes, char)
		}
	}

	return string(resultRunes), nil
}

func (_ ZeroWidth) Find(text string) ([]FindChar, error) {
	return find(text, zeroWidthChars)
}

func (_ ZeroWidth) Remove(text string) (string, error) {
	return remove(text, zeroWidthChars)
}

func (_ ZeroWidth) FindSpace(text string) ([]FindChar, error) {
	return find(text, spaceChars)
}

func (_ ZeroWidth) RemoveSpace(text string) (string, error) {
	return remove(text, spaceChars)
}

func (_ ZeroWidth) FindVarSelector(text string) ([]FindChar, error) {
	return find(text, varSelectors)
}

func (_ ZeroWidth) RemoveVarSelector(text string) (string, error) {
	return remove(text, varSelectors)
}

func init() {
	// Variation Selectors
	for i := 0xFE00; i <= 0xFE0F; i++ {
		varSelectors = append(varSelectors, rune(i))
	}

	// Variation Selectors Supplement
	for i := 0xE0100; i <= 0xE01EF; i++ {
		varSelectors = append(varSelectors, rune(i))
	}
}
