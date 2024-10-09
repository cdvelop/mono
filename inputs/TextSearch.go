package inputs

func TextSearch(params ...any) textSearch {

	new := textSearch{
		attributes: attributes{
			htmlName:   "search",
			customName: "textSearch",
			// Pattern: `^[a-zA-ZÑñ0-9- ]{2,20}$`,
		},
		permitted: permitted{
			Letters:    true,
			Tilde:      false,
			Numbers:    true,
			Characters: []rune{'-', ' '},
			Minimum:    2,
			Maximum:    20,
		},
	}
	new.Set(&new.permitted, params)
	return new
}

type textSearch struct {
	attributes
	permitted
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
