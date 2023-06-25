package syllables

import (
	"strings"
)

const alphabet = "абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"

const (
	vowels           = "аеёиоуыэюяАЕЁИОУЫЭЮЯ"
	nonPairConsonant = "мрлнМРЛН"
	endConsonant     = "йЙ"
	consonants       = "бвгджзчшщБВГДЖЗЧШЩ"
	voiceless        = "цкшхфпстЦКШХФПСТ"
	modifier         = `ьъЬЪ`
)

//Программа института

const (
	iVowels       = "аеёиоуыэюяАЕЁИОУЫЭЮЯ" //4 гласные
	iSonorous     = "мрлнМРЛН"             //3 сонорные
	iVoiced       = "бвгджзБВГДЖЗ"         //2 звонкие
	iVoiceless    = "пфктшсхцчщ"           //1 шумные
	iEndConsonant = "йЙ"
)

func isiVowels(r rune) bool {
	return strings.ContainsRune(iVowels, r)
}

func isiSonorous(r rune) bool {
	return strings.ContainsRune(iSonorous, r)
}

func isiVoiced(r rune) bool {
	return strings.ContainsRune(iVoiced, r)
}

func isiVoiceless(r rune) bool {
	return strings.ContainsRune(iVoiceless, r)
}

func isiEndConsonant(r rune) bool {
	return strings.ContainsRune(iEndConsonant, r)
}

func isVowel(r rune) bool {
	return strings.ContainsRune(vowels, r)
}

func isNonPairConsonant(r rune) bool {
	return strings.ContainsRune(nonPairConsonant, r)
}

func isEndConsonant(r rune) bool {
	return strings.ContainsRune(endConsonant, r)
}

func isConsonant(r rune) bool {
	return strings.ContainsRune(consonants, r)
}

func isVoiceless(r rune) bool {
	return strings.ContainsRune(voiceless, r)
}

func isModifier(r rune) bool {
	return strings.ContainsRune(modifier, r)
}

func nonLetter(r rune) bool {
	return !strings.ContainsRune(alphabet, r)
}
