package inputs

//value="valor a mostrar"
func Info(params ...any) *info {
	new := &info{
		input: input{
			attributes: attributes{
				htmlName:   "text",
				customName: "info",
			},
		},
	}
	return new
}

// input de carácter informativo
type info struct {
	input
}

func (i info) Render(tabIndex *int) string {
	return i.Value
}

func (i info) Validate(string) error {
	return nil
}
