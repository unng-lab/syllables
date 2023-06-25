package syllables

import (
	"reflect"
	"strings"
	"testing"
)

const sep = "-"

var tests = []string{
	//"а-ист",
	//"а-у-ра",
	"бул-ка",
	"ве-сна",
	"вьё-тся",
	"и-зжить",
	"кла-ссный",
	"кор-ка",
	"конь-ки",
	"ло-па-та",
	"метр",
	"ми-ни-а-тю-ра",
	"мо-шка",
	"мы-ться",
	"о-ття-нуть",
	"подъ-езд",
	"поль-ка",
	"по-чта",
	"по-э-ма",
	"ра-ссчи-та-нный",
	"слой-ка",
	"со-лом-ка",
	"со-ци-аль-на-я",
	"то-чка",
	"у-е-зжать",
	"вдо-хно-ве-ни-е",
	"ко-ттедж",
	"о-ття-ну-ться",
}

func clear(str string) string {
	return strings.ReplaceAll(str, sep, "")
}

func convert(str string) []string {
	return strings.Split(str, sep)
}

func TestDo(t *testing.T) {
	for _, tt := range tests {
		t.Run(clear(tt), func(t *testing.T) {
			if got := Do(clear(tt)); !reflect.DeepEqual(got, convert(tt)) {
				t.Errorf("Do() = %v, want %v", got, convert(tt))
			}
		})
	}
}
