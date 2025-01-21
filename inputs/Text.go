package inputs

import "strings"

func Text(params ...any) *text {
	new := &text{
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
	out = []string{"peréz del rozal", " estos son \\n los podria"}
	out = append(out, wrong_data...)
	return
}

func TextNumberCode(params ...any) *textNumCode {
	new := &textNumCode{
		input: input{
			attributes: attributes{
				htmlName:   "tel",
				customName: "textNumberCode",
			},
			permitted: permitted{
				Letters:    true,
				Numbers:    true,
				Characters: []rune{'_', '-'},
				Minimum:    2,
				Maximum:    15,
				StartWith: &permitted{
					Letters: true,
					Numbers: true,
				},
			},
		},
	}
	new.Set(params)

	return new
}

// texto y numero para código ej: V234
type textNumCode struct {
	input
}

func (t textNumCode) GoodTestData() (out []string) {
	out = []string{"et1", "12f", "GH03", "JJ10", "Wednesday", "los567", "677GH", "son_24_botellas"}
	return
}

func (t textNumCode) WrongTestData() (out []string) {
	out = []string{"los cuatro", "son 2 cuadros"}
	out = append(out, wrong_data...)
	return
}

// parámetros opcionales:
// "hidden" si se vera oculto o no.
func TextOnly(params ...any) *textOnly {
	new := &textOnly{
		input: input{
			attributes: attributes{
				htmlName:   "text",
				customName: "textOnly",
			},
			permitted: permitted{
				Letters:    true,
				Minimum:    3,
				Maximum:    50,
				Characters: []rune{' '},
			},
		},
	}
	new.Set(params)

	return new
}

type textOnly struct {
	input
}

func (t textOnly) GoodTestData() (out []string) {
	out = []string{"Ñuñez perez", "juli", "luz", "hola que tal", "Wednesday", "lost"}
	return
}

func (t textOnly) WrongTestData() (out []string) {
	out = []string{"Dr. xxx 788", "peréz. del jkjk", "los,true, vengadores"}
	out = append(out, wrong_data...)
	return
}

func TextNumber(params ...any) *textNumber {
	new := &textNumber{
		input: input{
			attributes: attributes{
				htmlName:   "text",
				customName: "textNumber",
				// Pattern:    `^[A-Za-z0-9_]{5,20}`,
			},
			permitted: permitted{
				Letters:    true,
				Numbers:    true,
				Characters: []rune{'_'},
				Minimum:    5,
				Maximum:    20,
			},
		},
	}
	new.Set(params)

	return new
}

// texto, numero y guion bajo 5 a 15 caracteres
type textNumber struct {
	input
}

func (t textNumber) GoodTestData() (out []string) {
	out = []string{"pc_caja", "pc_20", "info_40", "pc_50", "son_24_botellas", "los_cuatro", "son_2_cuadros"}
	return
}

func (t textNumber) WrongTestData() (out []string) {
	out = []string{"los cuatro", "tres", "et1_"}
	out = append(out, wrong_data...)
	return
}
