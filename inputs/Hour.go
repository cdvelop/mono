package inputs

import "errors"

// formato 08:00
// options: min="08:00", max="17:00"
func Hour(params ...any) *hour {
	new := &hour{
		input: input{
			attributes: attributes{
				htmlName:   "time",
				customName: "Hour",
				Title:      `title="formato hora: HH:MM"`,
				// Pattern: `^([0-9]|0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]`,
			},
			permitted: permitted{
				Numbers:    true,
				Characters: []rune{':'},
				Minimum:    5,
				Maximum:    5,
			},
		},
	}
	new.Set(params)

	return new
}

type hour struct {
	input
}

func (h hour) Validate(value string) error {

	if len(value) >= 2 && value[0] == '2' && value[1] == '4' {
		return errors.New("la hora 24 no existe")
	}

	return h.permitted.Validate(value)

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
