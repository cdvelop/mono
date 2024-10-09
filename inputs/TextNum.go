package inputs

func TextNumber(params ...any) textNumber {
	new := textNumber{
		attributes: attributes{
			htmlName:   "text",
			customName: "textNumber",
			// Pattern: `^[A-Za-z0-9_]{5,20}$`,
		},
		permitted: permitted{
			Letters:    true,
			Numbers:    true,
			Characters: []rune{'_'},
			Minimum:    5,
			Maximum:    20,
		},
	}
	new.Set(&new.permitted, params)

	return new
}

// texto, numero y guion bajo 5 a 15 caracteres
type textNumber struct {
	attributes
	permitted
}

func (t textNumber) GoodTestData() (out []string) {
	out = []string{
		"pc_caja",
		"pc_20",
		"info_40",
		"pc_50",
		"son_24_botellas",
		"los_cuatro",
		"son_2_cuadros",
	}
	return
}

func (t textNumber) WrongTestData() (out []string) {

	out = []string{
		"los cuatro",
		"tres",
		"et1_",
	}
	out = append(out, wrong_data...)

	return
}
