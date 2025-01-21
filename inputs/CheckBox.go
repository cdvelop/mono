package inputs

// eg: options=1:Admin,2:Editor,3:Visitante
func CheckBox(params ...any) *check {
	new := &check{
		input: input{
			attributes: attributes{
				htmlName: "checkbox",
			},
		},
	}
	new.Set(params)
	return new
}

type check struct {
	input
}
