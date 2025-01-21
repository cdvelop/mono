package inputs

// eg: options=1:Admin,2:Editor,3:Visitante
func DataList(params ...any) *input {
	new := &input{
		attributes: attributes{
			htmlName: "datalist",
		},
	}
	new.Set(params)

	return new
}
