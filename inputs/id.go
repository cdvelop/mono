package inputs

func ID(params ...any) id {
	new := id{
		input: input{
			attributes: attributes{
				htmlName:   "hidden",
				customName: "id",
			},
			permitted: permitted{
				Letters:    false,
				Numbers:    true,
				Characters: []rune{'.'},
				Minimum:    1,  //
				Maximum:    39, //
			},
		},
	}
	new.Set(params)

	return new
}

type id struct {
	input
}

func (i id) GoodTestData() (out []string) {

	return []string{
		"56988765432",
		"1234567",
		"0",
		"123456789",
		"100",
		"5000",
		"423456789",
		"31",
		"523756789",
		"10000232326263727",
		"29",
		"923726789",
		"3234567",
		"823456789",
		"29",
	}
}

func (i id) WrongTestData() (out []string) {
	out = []string{"1-1", "-100", "h", "h1", "-1", " ", "", "#", "& ", "% &"}

	return
}
