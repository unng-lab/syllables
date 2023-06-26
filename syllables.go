package syllables

import (
	"strings"
)

type s struct {
	word   []rune
	pos    int
	curSyl []rune
	res    []string
}

func (s *s) first() bool {
	return s.pos == 0
}

func (s *s) last() bool {
	return s.pos == len(s.word)
}

func (s *s) next() string {
	return string(s.word[s.pos+1])
}

func (s *s) haveNNext() bool {
	return s.pos+1 < len(s.word)
}

func (s *s) nNext() string {
	return string(s.word[s.pos+1])
}

func (s *s) prev() string {
	return string(s.word[s.pos-1])
}

func newS(word string) *s {
	return &s{
		word:   []rune(word),
		curSyl: make([]rune, 0, 2),
	}
}

//Слоги делятся по восходящей звучности. Каждый звук имеет звучность от 1 до 4: 4 — гласные, 3 — сонорные,
//2 — шумные звонкие, 1 — шумные глухие. Конец слога должен быть с максимальной звучностью (цифра 4):
//ко-ле-со (14-34-14), ко-мпью-тер (14-31 4-143).
//
//Для сонорных есть правила.
//
//Сонорные между гласными относят к следующему слогу: ка-рман (14-3343). Допускается вариант: кар-ман (ошибки не будет).
//Если есть сочетание сонорного и шумного, то раздел ставится между ними: пар-та (143-14).
//Й всегда остается с предыдущим слогом: тай-на (143-34).

func Do(word string) []string {
	s := newS(word)
	for s.pos = range s.word {
		// если не буква тогда возвращаем, то что уже посчитали
		if !strings.ContainsRune(alphabet, s.word[s.pos]) {
			return s.res
		}
		// если это последняя буква, то если это гласная мы добавляем слог если любая другая, то мы цепляемся
		// к предыдущему слогу
		if s.pos == len(s.word)-1 {
			if !strings.ContainsRune(iVowels, s.word[s.pos]) {
				s.curSyl = append(s.curSyl, s.word[s.pos])
				s.res[len(s.res)-1] += string(s.curSyl)
			} else {
				s.curSyl = append(s.curSyl, s.word[s.pos])
				s.res = append(s.res, string(s.curSyl))
			}
			return s.res
		}

		if snrt(s.word[s.pos]) == 4 {
			s.curSyl = append(s.curSyl, s.word[s.pos])
			s.res = append(s.res, string(s.curSyl))
			s.curSyl = make([]rune, 0, 2)
		} else {
			if isiSonorous(s.word[s.pos]) &&
				(s.pos+2 < len(s.word)) &&
				isiVoiceless(s.word[s.pos+1]) {
				if len(s.curSyl) == 0 {
					s.res[len(s.res)-1] += string(s.word[s.pos])
				} else {
					panic("разобрать сонорные")
				}
			} else if isiEndConsonant(s.word[s.pos]) {
				if len(s.curSyl) == 0 {
					s.res[len(s.res)-1] += string(s.word[s.pos])
				} else {
					panic("разобрать й")
				}
			} else if isModifier(s.word[s.pos]) &&
				len(s.res) > 0 &&
				!(s.pos+2 <= len(s.word) &&
					s.word[s.pos-1] == 'т' &&
					s.word[s.pos] == 'ь' &&
					s.word[s.pos+1] == 'с' &&
					s.word[s.pos+2] == 'я') {
				if len(s.curSyl) == 1 {
					s.curSyl = append(s.curSyl, s.word[s.pos])
					s.res[len(s.res)-1] += string(s.curSyl)
					s.curSyl = make([]rune, 0, 2)
				} else {
					panic("разобрать модификатор")
				}
			} else {
				s.curSyl = append(s.curSyl, s.word[s.pos])
			}

		}
	}

	panic("разобраться")
}

// звучность
func snrt(r rune) int {
	if strings.ContainsRune(iVowels, r) {
		return 4
	}

	if strings.ContainsRune(iSonorous, r) {
		return 3
	}

	if strings.ContainsRune(iVoiced, r) {
		return 2
	}

	if strings.ContainsRune(iVoiceless, r) {
		return 1
	}

	if strings.ContainsRune(iEndConsonant, r) {
		return 0
	}

	return -1
}

func joinRunes(runes ...rune) string {
	var sb strings.Builder
	for _, r := range runes {
		sb.WriteRune(r)
	}
	return sb.String()
}
