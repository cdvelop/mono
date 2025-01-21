package inputs

// ej: options=m:male,f:female
func Radio(params ...any) *radio {
	new := &radio{
		input: input{
			attributes: attributes{
				htmlName: "radio",
				Onchange: `onchange="RadioChange(this);"`,
			},
		},
	}
	new.Set(params)
	return new
}

// ej: {"f": "Femenino", "m": "Masculino"}.
func RadioGender(params ...any) *radio {
	options := append(params, "name=genre", `options=f:`+Lang.T(D.Female)+`,m:`+Lang.T(D.Male)+``)
	return Radio(options...)
}

type radio struct {
	input
}
