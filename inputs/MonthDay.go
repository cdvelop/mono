package inputs

// options: "hidden": campo oculto para el usuario
func MonthDay(params ...any) monthDay {
	new := monthDay{
		input: input{
			attributes: attributes{
				htmlName:   "text",
				customName: "monthDay",
				// Pattern:    `^[0-9]{2,2}$`,
			},
			permitted: permitted{
				Numbers:    true,
				Characters: []rune{},
				Minimum:    2,
				Maximum:    2,
			},
		},
	}
	new.Set(params)

	return new
}

// formato fecha: DD-MM
type monthDay struct {
	input
}

func (m monthDay) GoodTestData() (out []string) {

	out = []string{
		"01",
		"30",
		"03",
		"22",
		"31",
		"29",
		"10",
		"12",
		"05",
	}

	return
}

func (m monthDay) WrongTestData() (out []string) {
	out = []string{
		"1-1",
		"21/12",
	}

	out = append(out, wrong_data...)

	return
}
