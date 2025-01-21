package inputs

import (
	"strings"
)

const tabulation = '	'
const white_space = ' '
const break_line = '\n'

var valid_letters = map[rune]bool{
	'a': true, 'b': true, 'c': true, 'd': true, 'e': true, 'f': true, 'g': true, 'h': true, 'i': true,
	'j': true, 'k': true, 'l': true, 'm': true, 'n': true, 'o': true, 'p': true, 'q': true, 'r': true,
	's': true, 't': true, 'u': true, 'v': true, 'w': true, 'x': true, 'y': true, 'z': true,
	'ñ': true,

	'A': true, 'B': true, 'C': true, 'D': true, 'E': true, 'F': true, 'G': true, 'H': true, 'I': true,
	'J': true, 'K': true, 'L': true, 'M': true, 'N': true, 'O': true, 'P': true, 'Q': true, 'R': true,
	'S': true, 'T': true, 'U': true, 'V': true, 'W': true, 'X': true, 'Y': true, 'Z': true,
	'Ñ': true,
}

var valid_tilde = map[rune]bool{
	'á': true, 'é': true, 'í': true, 'ó': true, 'ú': true,
}

var valid_number = map[rune]bool{
	'0': true, '1': true, '2': true, '3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true,
}

type permitted struct {
	Letters         bool
	Tilde           bool
	Numbers         bool
	BreakLine       bool     // line breaks allowed
	WhiteSpaces     bool     // white spaces allowed
	Tabulation      bool     // tabulation allowed
	TextNotAllowed  []string // text not allowed eg: "hola" not allowed
	Characters      []rune   // other special characters eg: '\','/','@'
	Minimum         int      // min characters eg 2 "lo" ok default 0 no defined
	Maximum         int      // max characters eg 1 "l" ok default 0 no defined}
	ExtraValidation extraValidation
	StartWith       *permitted // characters allowed at the beginning
}

type extraValidation interface {
	ExtraValidation(text string) error
}

func (h input) Validate(text string) error {

	switch h.htmlName {
	case "checkbox", "radio", "datalist", "select":
		return h.checkOptionKeys(text)
	}

	// Validar empty string
	if len(text) == 0 {
		if h.allowSkipCompleted {
			return nil
		}
		return Lang.Err(D.Field, D.Empty, D.NotAllowed)
	}

	if h.StartWith != nil {
		char := text[0:1]
		if err := h.StartWith.validate(char); err != nil {

			if char == " " {
				return Lang.Err(D.DoNotStartWith, D.WhiteSpace)
			}

			return Lang.Err(D.DoNotStartWith, char)
		}
	}

	if h.ExtraValidation != nil {
		if err := h.ExtraValidation.ExtraValidation(text); err != nil {
			return err
		}
	}

	return h.permitted.validate(text)
}

func (h permitted) validate(text string) (err error) {

	if h.Minimum != 0 {
		if len(text) < h.Minimum {
			return Lang.Err(D.MinSize, h.Minimum, D.Chars)
		}
	}

	if h.Maximum != 0 {
		if len(text) > h.Maximum {
			return Lang.Err(D.MaxSize, h.Maximum, D.Chars)
		}
	}

	if len(h.TextNotAllowed) != 0 {
		for _, notAllowed := range h.TextNotAllowed {
			if strings.Contains(text, notAllowed) {
				return Lang.Err(D.NotAllowed, ':', h.TextNotAllowed)
			}
		}
	}

	for _, char := range text {
		if char == tabulation && h.Tabulation {
			continue
		}

		if char == white_space && h.WhiteSpaces {
			continue
		}

		if char == break_line && h.BreakLine {
			continue
		}

		if h.Letters {
			if !valid_letters[char] {
				err = Lang.Err(char, D.NotLetter)
			} else {
				err = nil
				continue
			}
		}

		if h.Tilde {
			if !valid_tilde[char] {
				err = Lang.Err(char, D.TildeNotAllowed)
			} else {
				err = nil
				continue
			}
		}

		if h.Numbers {
			if !valid_number[char] {
				if char == ' ' {
					err = Lang.Err(D.WhiteSpace, D.NotAllowed)
				} else {
					err = Lang.Err(char, D.NotNumber)
				}
			} else {
				err = nil
				continue
			}
		}

		if len(h.Characters) != 0 {
			var found bool
			for _, c := range h.Characters {
				if c == char {
					found = true
					break
				}
			}

			if found {
				err = nil
				continue
			} else {
				if char == white_space {
					return Lang.Err(D.WhiteSpace, D.NotAllowed)
				} else if valid_tilde[char] {
					return Lang.Err(char, D.TildeNotAllowed)
				} else if char == tabulation {
					return Lang.Err(D.TabText, D.NotAllowed)
				} else if char == break_line {
					return Lang.Err(D.Newline, D.NotAllowed)
				}
				return Lang.Err(D.Char, char, D.NotAllowed)
			}
		}

		if err != nil {
			return err
		}
	}

	return err
}

// const carriage_return = '\r'

func (a attributes) checkOptionKeys(value string) error {

	dataInArray := strings.Split(value, ",")

	for _, keyIn := range dataInArray {

		if keyIn == "" {
			return Lang.Err(D.RequiredSelection, D.Field, a.Name)
		}

		var exist bool
		// fmt.Println("a.optionKeys", a.optionKeys)
		for _, opt := range a.options {
			if _, exist = opt[keyIn]; exist {
				break
			}
		}

		if !exist {
			return Lang.Err(D.Value, keyIn, D.NotAllowed, D.In, a.htmlName, a.Name)
		}

	}

	return nil

}

// Define un mapa de caracteres válidos
func (p permitted) MinMaxAllowedChars() (min, max int) {
	return p.Minimum, p.Maximum
}
