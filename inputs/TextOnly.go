package inputs

// parámetros opcionales:
// "hidden" si se vera oculto o no.
func TextOnly(params ...any) textOnly {
	new := textOnly{
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

	out = []string{
		"Ñuñez perez",
		"juli",
		"luz",
		"hola que tal",
		"Wednesday",
		"lost",
	}

	return
}

func (t textOnly) WrongTestData() (out []string) {

	out = []string{
		"Dr. xxx 788",
		"peréz. del jkjk",
		"los,true, vengadores",
	}
	out = append(out, wrong_data...)

	return
}
