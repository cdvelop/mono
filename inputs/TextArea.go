package inputs

// options:
// title="permitido letras números - , :"
// cols="2" default 1
// rows="8" default 3
func TextArea(params ...any) textArea {
	new := textArea{
		input: input{
			attributes: attributes{
				Rows: `rows="3"`,
				Cols: `cols="1"`,
				// PlaceHolder: `placeHolder="` + info + `"`,
				// Pattern: `^[A-Za-zÑñáéíóú 0-9:$%.,+-/\\()|\n/g]{2,1000}`,
				Oninput: `oninput="TexAreaOninput(this)"`,
				// Onchange: `onchange="` + DefaultValidateFunction + `"`,
			},
			permitted: permitted{
				Letters:     true,
				Tilde:       true,
				Numbers:     true,
				BreakLine:   true,
				WhiteSpaces: true,
				Tabulation:  true,
				Characters:  []rune{'$', '%', '+', '#', '-', '.', ',', ':', '(', ')'},
				Minimum:     2,
				Maximum:     1000,
			},
		},
	}
	new.Set(params)

	return new
}

type textArea struct {
	input
}

func (t textArea) InputName(customName, htmlName *string) {
	if customName != nil {
		*customName = "TextArea"
	}
	if htmlName != nil {
		*htmlName = "textarea"
	}
}

func (t textArea) ResetParameters() any {

	return &struct {
		ResetJsFuncName    string
		Enable             bool
		NotSendQueryObject bool
		Params             map[string]any
	}{
		ResetJsFuncName: "ResetTextArea",
		Enable:          true,
	}
}
