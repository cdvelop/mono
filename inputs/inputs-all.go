package inputs

var badTestData = []string{"", " ", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "ñ", "Ñ", "á", "é", "í", "ó", "ú", "Á", "É", "Í", "Ó", "Ú", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "+", "=", "{", "}", "[", "]", "|", "\\", ":", ";", "\"", "'", "<", ">", ",", ".", "?", "/", "`", "~"}

// eg: options=1:Admin,2:Editor,3:Visitante
func DataList(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "datalist"},
	}
	new.Set(params)

	return new
}

// eg: options=1:Admin,2:Editor,3:Visitante
func CheckBox(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "checkbox"},
	}
	new.Set(params)
	return new
}

// formato fecha: DD-MM-YYYY
func DateAge(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "date", customName: "DateAge", Title: `title="formato fecha: DD-MM-YYYY"`, Onchange: `onchange="DateAgeChange(this)"`},
	}
	new.Set(params)
	return new
}

// formato dia DD como palabra ej. Lunes 24 Diciembre
func DayWord(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "text", customName: "DayWord", DataSet: []map[string]string{{"spanish": ""}}},
	}
	new.Set(params)
	return new
}

// ej: options=m:male,f:female
func Radio(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "radio", Onchange: `onchange="RadioChange(this);"`},
	}
	new.Set(params)
	return new
}

// ej: {"f": "Femenino", "m": "Masculino"}.
func RadioGender(params ...any) *input {
	options := append(params, "name=genre", `options=f:`+Lang.T(D.Female)+`,m:`+Lang.T(D.Male)+``)
	return Radio(options...)
}

func Text(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "text"},
		permitted:  permitted{Letters: true, Tilde: false, Numbers: true, Characters: []rune{' ', '.', ',', '(', ')'}, Minimum: 2, Maximum: 100},
		testData:   testData{Good: []string{"Maria", "Juan", "Marcela", "Luz", "Carmen", "Jose", "Octavio", "Soto", "Del Rosario", "Del Carmen", "Ñuñez", "Perez", "Cadiz", "Caro", "Dr. Maria Jose Diaz Cadiz", "son 4 (4 bidones)", "pc dental (1)", "equipo (4)"}, Bad: append([]string{"peréz del rozal", " estos son \\n los podria"}, badTestData...)},
	}
	new.Set(params)
	return new
}

func TextNumberCode(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "tel", customName: "textNumberCode"},
		permitted:  permitted{Letters: true, Numbers: true, Characters: []rune{'_', '-'}, Minimum: 2, Maximum: 15, StartWith: &permitted{Letters: true, Numbers: true}},
		testData:   testData{Good: []string{"et1", "12f", "GH03", "JJ10", "Wednesday", "los567", "677GH", "son_24_botellas"}, Bad: append([]string{"los cuatro", "son 2 cuadros"}, badTestData...)},
	}
	new.Set(params)
	return new
}

func TextOnly(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "text", customName: "textOnly"},
		permitted:  permitted{Letters: true, Minimum: 3, Maximum: 50, Characters: []rune{' '}},
		testData:   testData{Good: []string{"Ñuñez perez", "juli", "luz", "hola que tal", "Wednesday", "lost"}, Bad: append([]string{"Dr. xxx 788", "peréz. del jkjk", "los,true, vengadores"}, badTestData...)},
	}
	new.Set(params)
	return new
}

func TextNumber(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "text", customName: "textNumber"},
		permitted:  permitted{Letters: true, Numbers: true, Characters: []rune{'_'}, Minimum: 5, Maximum: 20},
		testData:   testData{Good: []string{"pc_caja", "pc_20", "info_40", "pc_50", "son_24_botellas", "los_cuatro", "son_2_cuadros"}, Bad: append([]string{"los cuatro", "tres", "et1_"}, badTestData...)},
	}
	new.Set(params)
	return new
}

func TextSearch(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "search", customName: "textSearch"},
		permitted:  permitted{Letters: true, Tilde: false, Numbers: true, Characters: []rune{'-', ' '}, Minimum: 2, Maximum: 20},
		testData:   testData{Good: []string{"Ñuñez perez", "Maria Jose Diaz", "12038-0", "1990-07-21", "190-07-21", "lost"}, Bad: []string{"Dr. xxx 788", "peréz del jkjk", "los,true, vengadores", "0", " ", "", "#", "& ", "% &", "SELECT * FROM", "="}},
	}
	new.Set(params)
	return new
}
