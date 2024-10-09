package inputs

func (r rut) GoodTestData() (out []string) {

	ok_rut := []string{
		"22537160-1",
		"5008452-3",
		"10493788-8",
		"21821424-k",
		"15890022-k",
		"7467499-2",
		"21129619-4",
		"24287548-6",
		"15093641-1",
		"10245390-5",
	}

	if r.dni_mode {
		out = []string{
			"15890022-k",
			"ax001223b",
			"A0C00A389",
			"B0004DF678",
		}
		out = append(out, ok_rut...)

		return

	} else {
		return ok_rut
	}

}

func (r rut) WrongTestData() (out []string) {

	out = []string{
		"7863697-k",
		"7863697-0",
		"14080717-0",
		"07863697-1",
		" - 100 ",
		"-100",
	}
	out = append(out, wrong_data...)

	return
}
