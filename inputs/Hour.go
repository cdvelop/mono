package inputs

// formato 08:00
// options: min="08:00", max="17:00"
func Hour(params ...any) *hour {
	new := &hour{
		input: input{
			attributes: attributes{
				htmlName:   "time",
				customName: "Hour",
				Title:      `title="` + Lang.T(D.Format, D.Hour, ':') + ` HH:MM"`,
			},
			permitted: permitted{
				Numbers:        true,
				Characters:     []rune{':'},
				Minimum:        5,
				Maximum:        5,
				TextNotAllowed: []string{"24:"},
			},
		},
	}
	new.Set(params)

	return new
}

type hour struct {
	input
}

func (h hour) GoodTestData() (out []string) {
	out = []string{
		"23:59",
		"00:00",
		"12:00",
		"13:17",
		"21:53",
		"00:40",
		"08:30",
		"12:00",
		"15:01",
	}

	return
}

func (h hour) WrongTestData() (out []string) {
	out = []string{
		"24:00",
		"13-34",
	}
	out = append(out, wrong_data...)
	return
}
