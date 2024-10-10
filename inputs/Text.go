package inputs

import "strings"

func Text(params ...any) text {
	new := text{
		input: input{
			attributes: attributes{
				htmlName: "text",
			},
			permitted: permitted{
				Letters:    true,
				Tilde:      false,
				Numbers:    true,
				Characters: []rune{' ', '.', ',', '(', ')'},
				Minimum:    2,
				Maximum:    100,
			},
		},
	}
	new.Set(params)

	return new
}

// texto,punto,coma, paréntesis o números permitidos
type text struct {
	input
}

// options: first_name,last_name, phrase
func (t text) GoodTestData() (out []string) {

	first_name := []string{"Maria", "Juan", "Marcela", "Luz", "Carmen", "Jose", "Octavio"}

	last_name := []string{"Soto", "Del Rosario", "Del Carmen", "Ñuñez", "Perez", "Cadiz", "Caro"}

	phrase := []string{"Dr. Maria Jose Diaz Cadiz", "son 4 (4 bidones)", "pc dental (1)", "equipo (4)"}

	placeholder := strings.ToLower(t.PlaceHolder)

	switch {
	case strings.Contains(placeholder, "nombre y apellido"):

		return permutation(first_name, last_name)
	case strings.Contains(placeholder, "nombre"):
		return first_name

	case strings.Contains(placeholder, "apellido"):
		return last_name

	default:
		out = append(out, phrase...)
		out = append(out, first_name...)
		out = append(out, last_name...)
	}

	return
}

func (text) WrongTestData() (out []string) {

	out = []string{
		"peréz del rozal",
		" estos son \\n los podria",
	}

	out = append(out, wrong_data...)

	return
}
