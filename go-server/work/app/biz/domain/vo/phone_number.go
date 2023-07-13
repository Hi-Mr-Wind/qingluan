package vo

import (
	"go-service/pkg/errno"
	"regexp"
	"strings"
)

type PhoneNumber struct {
	number string
}

func NewPhoneNumber(number string) PhoneNumber {
	return PhoneNumber{number: strings.TrimSpace(number)}
}

func (p PhoneNumber) GetNumber() string {
	return p.number
}

func (p PhoneNumber) CheckFormat() (bool, error) {
	if p.number == "" {
		return false, errno.NewSimpleBizError(errno.ErrPhoneEmpty, nil)
	}
	if !isPhoneValid(p.number) {
		return false, errno.NewSimpleBizError(errno.ErrInvalidPhoneFormat, nil)
	}
	return true, nil
}

func (p PhoneNumber) GetCountryCode() string {
	for _, reg := range phoneRegularExp {
		if reg.regexp.MatchString(p.number) {
			return reg.countryCode
		}
	}
	return ""
}

func (p PhoneNumber) IsChinaMainlandPhone() bool {
	return p.GetCountryCode() == "zh-CN"
}

func isPhoneValid(number string) bool {
	for _, reg := range phoneRegularExp {
		if reg.regexp.MatchString(number) {
			return true
		}
	}
	return false
}

type PhoneReg struct {
	countryCode string
	regexp      *regexp.Regexp
}

var (
	phoneRegularExp = []PhoneReg{
		{"zh-CN", regexp.MustCompile("^(\\+?0?86-?)?1\\d{10}$")},
		{"zh-TW", regexp.MustCompile("^(\\+?886-?|0)?9\\d{8}$")},
		{"ar-DZ", regexp.MustCompile("^(\\+?213|0)[567]\\d{8}$")},
		{"ar-SY", regexp.MustCompile("^(!?(\\+?963)|0)?9\\d{8}$")},
		{"ar-SA", regexp.MustCompile("^(!?(\\+?966)|0)?5\\d{8}$")},
		{"en-US", regexp.MustCompile("^(00)?(1)\\d{10,12}$")},
		{"cs-CZ", regexp.MustCompile("^(\\+?420)? ?[1-9][0-9]{2} ?[0-9]{3} ?[0-9]{3}$")},
		{"de-DE", regexp.MustCompile("^(\\+?49[ .\\-])?([(][0-9]{1,6}[)])?([0-9 .\\-/]{3,20})((x|ext|extension)[ ]?[0-9]{1,4})?$")},
		{"da-DK", regexp.MustCompile("^(\\+?45)?(\\d{8})$")},
		{"el-GR", regexp.MustCompile("^(\\+?30)?(69\\d{8})$")},
		{"en-AU", regexp.MustCompile("^(\\+?61|0)4\\d{8}$")},
		{"en-GB", regexp.MustCompile("^(\\+?44|0)7\\d{9}$")},
		{"en-HK", regexp.MustCompile("^(\\+?852-?)?[569]\\d{3}-?\\d{4}$")},
		{"en-IN", regexp.MustCompile("^(\\+?91|0)?[789]\\d{9}$")},
		{"en-NZ", regexp.MustCompile("^(\\+?64|0)2\\d{7,9}$")},
		{"en-ZA", regexp.MustCompile("^(\\+?27|0)\\d{9}$")},
		{"en-ZM", regexp.MustCompile("^(\\+?26)?09[567]\\d{7}$")},
		{"es-ES", regexp.MustCompile("^(\\+?34)?(6\\d|7[1234])\\d{7}$")},
		{"fi-FI", regexp.MustCompile("^(\\+?358|0)\\s?(4[01245]?|50)\\s?(\\d\\s?){4,8}\\d$")},
		{"fr-FR", regexp.MustCompile("^(\\+?33|0)[67]\\d{8}$")},
		{"he-IL", regexp.MustCompile("^(\\+972|0)([23489]|5[0248]|77)[1-9]\\d{6}$")},
		{"hu-HU", regexp.MustCompile("^(\\+?36)(20|30|70)\\d{7}$")},
		{"it-IT", regexp.MustCompile("^(\\+?39)?\\s?3\\d{2} ?\\d{6,7}$")},
		{"ja-JP", regexp.MustCompile("^(\\+?81|0)\\d{1,4}[ \\-]?\\d{1,4}[ \\-]?\\d{4}$")},
		{"ms-MY", regexp.MustCompile("^(\\+?6?01)(([145](-|\\s)?\\d{7,8})|([236789](\\s|-)?\\d{7}))$")},
		{"nb-NO", regexp.MustCompile("^(\\+?47)?[49]\\d{7}$")},
		{"nl-BE", regexp.MustCompile("^(\\+?32|0)4?\\d{8}$")},
		{"nn-NO", regexp.MustCompile("^(\\+?47)?[49]\\d{7}$")},
		{"pl-PL", regexp.MustCompile("^(\\+?48)? ?[5-8]\\d ?\\d{3} ?\\d{2} ?\\d{2}$")},
		{"pt-BR", regexp.MustCompile("^(\\+?55|0)-?[1-9]{2}-?[2-9]\\d{3,4}-?\\d{4}$")},
		{"pt-PT", regexp.MustCompile("^(\\+?351)?9[1236]\\d{7}$")},
		{"ru-RU", regexp.MustCompile("^(\\+?7|8)?9\\d{9}$")},
		{"sr-RS", regexp.MustCompile("^(\\+3816|06)[- \\d]{5,9}$")},
		{"tr-TR", regexp.MustCompile("^(\\+?90|0)?5\\d{9}$")},
		{"", regexp.MustCompile("^\\d{1,14}$")},
	}
)
