package inputs

import (
	"errors"
	"strings"
)

type permitted struct {
	Letters     bool
	Tilde       bool
	Numbers     bool
	BreakLine   bool   // line breaks allowed
	WhiteSpaces bool   // white spaces allowed
	Tabulation  bool   // tabulation allowed
	Characters  []rune // other special characters eg: '\','/','@'
	Minimum     int    // min characters eg 2 "lo" ok default 0 no defined
	Maximum     int    // max characters eg 1 "l" ok default 0 no defined}
}

const tabulation = '	'
const white_space = ' '
const break_line = '\n'

// const carriage_return = '\r'

func (p permitted) Validate(text string) error {

	var err error

	if p.Minimum != 0 {
		if len(text) < p.Minimum {
			return errors.New(Lang.TNum("min_size", p.Minimum))
		}
	}

	if p.Maximum != 0 {
		if len(text) > p.Maximum {
			return errors.New(Lang.TNum("max_size", p.Maximum))
		}
	}

	for _, char := range text {
		if char == tabulation && p.Tabulation {
			continue
		}

		if char == white_space && p.WhiteSpaces {
			continue
		}

		if char == break_line && p.BreakLine {
			continue
		}

		if p.Letters {
			if !valid_letters[char] {
				err = errors.New(Lang.TChar("not_letter", string(char)))
			} else {
				err = nil
				continue
			}
		}

		if p.Tilde {
			if !valid_tilde[char] {
				err = errors.New(Lang.TChar("unsupported_tilde", string(char)))
			} else {
				err = nil
				continue
			}
		}

		if p.Numbers {
			if !valid_number[char] {
				if char == ' ' {
					err = errors.New(Lang.T("white_spaces"))
				} else {
					err = errors.New(Lang.TChar("not_number", string(char)))
				}
			} else {
				err = nil
				continue
			}
		}

		if len(p.Characters) != 0 {
			var found bool
			for _, c := range p.Characters {
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
					return errors.New(Lang.T("white_spaces"))
				} else if valid_tilde[char] {
					return errors.New(Lang.TChar("tilde_not_allowed", string(char)))
				} else if char == tabulation {
					return errors.New(Lang.T("tab_not_allowed"))
				} else if char == break_line {
					return errors.New(Lang.T("newline_not_allowed"))
				}
				return errors.New(Lang.TChar("not_allowed", string(char)))
			}
		}

		if err != nil {
			return err
		}
	}

	return err
}

func (a attributes) checkOptionKeys(value string) error {

	dataInArray := strings.Split(value, ",")

	for _, keyIn := range dataInArray {

		if keyIn == "" {
			return errors.New("selección requerida campo " + a.Name)
		}

		var exist bool
		// fmt.Println("a.optionKeys", a.optionKeys)
		for _, opt := range a.options {
			if _, exist = opt[keyIn]; exist {
				break
			}
		}

		if !exist {
			return errors.New("valor " + keyIn + " no permitido en " + a.htmlName + " " + a.Name)
		}

	}

	return nil

}

func (a attributes) GoodTestData() (out []string) {
	for _, opt := range a.options {
		for k := range opt {
			out = append(out, k)
		}
	}
	return
}

func (a attributes) WrongTestData() (out []string) {
	for _, wd := range wrong_data {
		found := false
		for _, opt := range a.options {
			if _, exists := opt[wd]; exists {
				found = true
				break
			}
		}
		if !found {
			out = append(out, wd)
		}
	}
	return
}

// Define un mapa de caracteres válidos
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

func (p permitted) MinMaxAllowedChars() (min, max int) {
	return p.Minimum, p.Maximum
}
