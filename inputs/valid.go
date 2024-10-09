package inputs

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type permitted struct {
	Letters     bool
	Tilde       bool
	Numbers     bool
	BreakLine   bool   // saltos de linea permitidos
	WhiteSpaces bool   // permitidos espacios en blanco
	Tabulation  bool   // permitido tabular
	Characters  []rune // otros caracteres especiales ej: '\','/','@'
	Minimum     int    //caracteres min ej 2 "lo" ok default 0 no defined
	Maximum     int    //caracteres max ej 1 "l" ok default 0 no defined
}

const tabulation = '	'
const white_space = ' '
const break_line = '\n'

// const carriage_return = '\r'
const errorWhiteSpace = "espacios en blanco no permitidos"

func (p permitted) Validate(text string) error {
	var err error

	if p.Minimum != 0 {
		if len(text) < p.Minimum {
			return errors.New("tamaño mínimo " + strconv.Itoa(p.Minimum) + " caracteres")
		}
	}

	if p.Maximum != 0 {
		if len(text) > p.Maximum {
			return errors.New("tamaño máximo " + strconv.Itoa(p.Maximum) + " caracteres")
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
				err = errors.New(string(char) + " no es una letra")
			} else {
				err = nil
				continue
			}
		}

		if p.Tilde {
			if !valid_tilde[char] {
				err = errors.New("tilde " + string(char) + " no soportada")
			} else {
				err = nil
				continue
			}
		}

		if p.Numbers {
			if !valid_number[char] {
				if char == ' ' {
					err = errors.New(errorWhiteSpace)
				} else {
					err = errors.New(string(char) + " no es un numero")
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
					return errors.New("espacios en blanco no permitidos")
				} else if valid_tilde[char] {
					return errors.New(string(char) + " con tilde no permitida")
				} else if char == tabulation {
					return errors.New("tabulation de texto no permitida")
				} else if char == break_line {
					return errors.New("salto de linea no permitido")
				}
				return errors.New("carácter " + string(char) + " no permitido")
			}
		}

		if err != nil {
			return err
		}
	}

	return err
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

func (p permitted) ValidateInput(value string) error {
	return p.Validate(value)
}

// Método para generar dinámicamente el título
func (p *permitted) setDynamicTitle(attr *attributes) {
	var parts []string
	parts = append(parts, "texto")

	// Lógica de validación para letras
	if p.Letters {
		parts = append(parts, "letras")
	}

	// Lógica de validación para números
	if p.Numbers {
		parts = append(parts, "números")
	}

	// Lógica de validación para caracteres permitidos
	if len(p.Characters) > 0 {
		var chars []string
		for _, char := range p.Characters {
			if char == ' ' {
				chars = append(chars, "␣") // Reemplaza el espacio con el carácter visible '␣'
			} else {
				chars = append(chars, string(char))
			}
		}
		parts = append(parts, fmt.Sprintf("caracteres permitidos: %v ", strings.Join(chars, " ")))
	}

	if p.Minimum != 0 {
		parts = append(parts, fmt.Sprintf("min. %d", p.Minimum))
	}

	if p.Maximum != 0 {
		parts = append(parts, fmt.Sprintf("max. %d", p.Maximum))
	}

	// Generar el valor final para el atributo Title
	attr.Title = strings.Join(parts, " ")
}
