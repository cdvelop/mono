package inputs

// format date: html DD-MM-YYYY - validation YYYY-MM-DD
func Date(params ...any) *date {
	new := &date{
		input: input{
			attributes: attributes{
				htmlName: "date",
				Title:    `title="` + Lang.T(D.Format, D.Date, ':') + `: DD-MM-YYYY"`,
			},
		},
	}
	new.Set(params)
	return new
}

type date struct {
	input
}

// validación con datos de entrada
func (d date) Validate(value string) error {
	return d.CheckDateExists(value)
}

func (d date) GoodTestData() (out []string) {
	out = []string{"2002-01-03", "1998-02-01", "1999-03-08", "2022-04-21", "1999-05-30", "2020-09-29", "1991-10-02", "2000-11-12", "1993-12-15"}
	return
}

func (d date) WrongTestData() (out []string) {
	out = []string{"21/12/1998", "0000-00-00", "31-01"}
	out = append(out, wrong_data...)
	return
}

// options: "hidden": campo oculto para el usuario
func MonthDay(params ...any) *monthDay {
	new := &monthDay{
		input: input{
			attributes: attributes{
				htmlName:   "text",
				customName: "MonthDay",
			},
			permitted: permitted{
				Numbers:         true,
				Minimum:         2,
				Maximum:         2,
				ExtraValidation: &monthDay{},
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

func (m monthDay) ExtraValidation(text string) error {
	_, err := validateDay(text)
	return err
}

func (m monthDay) GoodTestData() (out []string) {
	out = []string{"01", "30", "03", "22", "31", "29", "10", "12", "05"}
	return
}

func (m monthDay) WrongTestData() (out []string) {
	out = []string{"1-1", "21/12"}

	out = append(out, wrong_data...)

	return
}
