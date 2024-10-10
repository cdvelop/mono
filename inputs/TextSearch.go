package inputs

func TextSearch(params ...any) textSearch {
	new := textSearch{
		input: input{
			attributes: attributes{
				htmlName:   "search",
				customName: "textSearch",
			},
			permitted: permitted{
				Letters:    true,
				Tilde:      false,
				Numbers:    true,
				Characters: []rune{'-', ' '},
				Minimum:    2,
				Maximum:    20,
			},
		},
	}
	new.Set(params)
	return new

}

type textSearch struct {
	input
}

func (s textSearch) GoodTestData() (out []string) {
	out = []string{
		"Ñuñez perez",
		"Maria Jose Diaz",
		"12038-0",
		"1990-07-21",
		"190-07-21",
		"lost",
	}
	return
}

func (s textSearch) WrongTestData() (out []string) {
	out = []string{
		"Dr. xxx 788",
		"peréz del jkjk",
		"los,true, vengadores",
		"0", " ", "", "#", "& ", "% &", "SELECT * FROM", "=",
	}

	return
}
