package zerowidth

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var zeroWidthChars = []rune{
	'\u200B', // SPACE
	'\u00A0', // NO-BREAK SPACE
	'\u2002', // EN SPACE
	'\u2003', // EM SPACE
	'\u2004', // THREE-PER-EM SPACE
	'\u2005', // FOUR-PER-EM SPACE
	'\u2006', // SIX-PER-EM SPACE
	'\u2007', // FIGURE SPACE
	'\u2008', // PUNCTUATION SPACE
	'\u2009', // THIN SPACE
	'\u200A', // HAIR SPACE
	'\u200B', // ZERO WIDTH SPACE
	'\u3000', // IDEOGRAPHIC SPACE
	'\uFEFF', // ZERO WIDTH NO-BREAK SPACE
	'\u0009', // CHARACTER TABULATION
	'\u3164', // 한글채움문자
	'\u2800', // Braille Pattern Blank
	'\u200D', // zero width joiner
	'\u115F', // 한글초성채움문자
	'\u1160', // 한글중성채움문자
	'\u00AD', // Soft Hyphen
	'\u202E', // RIGHT-TO-LEFT OVERRIDE
	'\uFE00', // Variation Selector-0
	'\uFE01', // Variation Selector-1
	'\uFE02', // Variation Selector-2
	'\uFE03', // Variation Selector-3
	'\uFE04', // Variation Selector-4
	'\uFE05', // Variation Selector-5
	'\uFE06', // Variation Selector-6
	'\uFE07', // Variation Selector-7
	'\uFE08', // Variation Selector-8
	'\uFE09', // Variation Selector-9
	'\uFE0A', // Variation Selector-A
	'\uFE0B', // Variation Selector-B
	'\uFE0C', // Variation Selector-C
	'\uFE0D', // Variation Selector-D
	'\uFE0E', // Variation Selector-E
	'\uFE0F', // Variation Selector-F
}

type ZeroWidth struct{}

type FindChar struct {
	Position int
	Char     rune
}

func NewZeroWidth() *ZeroWidth {
	return &ZeroWidth{}
}

func (z ZeroWidth) Find(text string) ([]FindChar, error) {
	var runePatterns []string
	for _, r := range zeroWidthChars {
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

func (z ZeroWidth) Remove(text string) (string, error) {
	findChars, err := z.Find(text)
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
