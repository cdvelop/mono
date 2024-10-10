package inputs

func Date(params ...any) date {
	new := date{
		input: input{
			attributes: attributes{
				htmlName: "date",
				Title:    `title="formato fecha: DD-MM-YYYY"`,
			},
			// Pattern: `[0-9]{4}-(0[1-9]|1[012])-(0[1-9]|1[0-9]|2[0-9]|3[01])`,
		},
	}
	new.Set(params)
	return new
}

// formato fecha: DD-MM-YYYY
type date struct {
	input
}

// validación con datos de entrada
func (d date) Validate(value string) error {
	return d.CheckDateExists(value)
}

func (d date) GoodTestData() (out []string) {
	out = []string{
		"2002-01-03",
		"1998-02-01",
		"1999-03-08",
		"2022-04-21",
		"1999-05-30",
		"2020-09-29",
		"1991-10-02",
		"2000-11-12",
		"1993-12-15",
	}
	return
}

func (d date) WrongTestData() (out []string) {

	out = []string{
		"21/12/1998",
		"0000-00-00",
		"31-01",
	}
	out = append(out, wrong_data...)

	return
}
