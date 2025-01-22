package inputs

import (
	"strconv"
	"strings"
)

var badTestData = []string{"", " ", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "ñ", "Ñ", "á", "é", "í", "ó", "ú", "Á", "É", "Í", "Ó", "Ú", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "+", "=", "{", "}", "[", "]", "|", "\\", ":", ";", "\"", "'", "<", ">", ",", ".", "?", "/", "`", "~"}

func Date(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "date", Title: `title="` + Lang.T(D.Format, D.Date, ':') + `: DD-MM-YYYY"`},
		permitted: permitted{
			ExtraValidation: CheckDateExists,
		},
		testData: testData{
			Good: []string{"2002-01-03", "1998-02-01", "1999-03-08", "2022-04-21", "1999-05-30", "2020-09-29", "1991-10-02", "2000-11-12", "1993-12-15"},
			Bad:  append([]string{"21/12/1998", "0000-00-00", "31-01"}, wrong_data...),
		},
	}
	new.Set(params)
	return new
}

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

func TextArea(params ...any) *input {
	new := &input{
		attributes: attributes{Rows: `rows="3"`, Cols: `cols="1"`, Oninput: `oninput="TexAreaOninput(this)"`},
		permitted:  permitted{Letters: true, Tilde: true, Numbers: true, BreakLine: true, WhiteSpaces: true, Tabulation: true, Characters: []rune{'$', '%', '+', '#', '-', '.', ',', ':', '(', ')'}, Minimum: 2, Maximum: 1000},
		testData: testData{
			Good: []string{"hola: esto, es. la - prueba 10", "soy ñato o Ñato (aqui)", "son dos examenes", "costo total es de $23.230. pesos"},
			Bad:  []string{"] son", " ", "& ", "SELECT * FROM", "="}},
	}
	new.Set(params)
	return new
}

func Password(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "password"},
		permitted:  permitted{Letters: true, Tilde: true, Numbers: true, Characters: []rune{' ', '#', '%', '?', '.', ',', '-', '_'}, Minimum: 5, Maximum: 50},
		testData: testData{
			Good: []string{"c0ntra3", "M1 contraseÑ4", "contrase", "cont", "12345", "UNA Frase tambien Cuenta", "DOS Frases tambien CuentaN", "CUATRO FraseS tambien CuentaN"},
			Bad:  []string{"", "Ñ", "c", " ", "2", "%", "sdlksññs092830928309280%%%%%9382¿323294720&&/0kdlskdlskdskdñskdlskdsññdkslkdñskdslkdsñ", "sdlksññs0928309283092809382%%¿323294720&&/0kdlskdlskdskdñskdlskdsññdkslkdñskdslkdsñ", "sdlksññs0928309283092809382¿78%%323294720&&/0kdlskdlskdskdñskdlskdsññdkslkdñskdslkdsñ"},
		},
	}
	new.Set(params)

	if new.Min != "" {
		new.Minimum, _ = strconv.Atoi(new.Min)
	}

	if new.Max != "" {
		new.Maximum, _ = strconv.Atoi(new.Max)
	}

	return new
}

func Number(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "number"},
		permitted:  permitted{Numbers: true, Minimum: 1, Maximum: 20},
		testData: testData{
			Good: []string{
				"56988765432", "1234567", "0", "123456789", "100", "5000",
				"423456789", "31", "523756789", "10000232326263727", "29",
				"923726789", "3234567", "823456789", "29",
			},
			Bad: []string{"1-1", "-100", "h", "h1", "-1", " ", "", "#", "& ", "% &"},
		},
	}
	new.Set(params)

	if new.Min != "" {
		new.Minimum, _ = strconv.Atoi(new.Min)
	}

	if new.Max != "" {
		new.Maximum, _ = strconv.Atoi(new.Max)
	}

	return new
}

func Phone(params ...any) *input {
	return Number(`min="7"`, `max="11"`)
}

func Mail(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "mail", PlaceHolder: `placeHolder="ej: mi.correo@mail.com"`},
		permitted: permitted{Letters: true, Numbers: true, Characters: []rune{'@', '.', '_'}, Minimum: 0, Maximum: 0, ExtraValidation: func(s string) error {
			if strings.Contains(s, "example") {
				return Lang.Err(D.Example, D.Email, D.NotAllowed)
			}

			parts := strings.Split(s, "@")
			if len(parts) != 2 {
				return Lang.Err(D.Format, D.Email, D.NotValid)
			}

			return nil
		}},
		testData: testData{
			Good: []string{"mi.correo@mail.com", "alguien@algunlugar.es", "ramon.bonachea@email.com", "r.bonachea@email.com", "laura@hellos.email.tk"},
			Bad:  []string{"email@example.com", "correomao@n//.oo", "son_24_bcoreos", "email@example"},
		},
	}
	new.Set(params)
	return new
}

func Ip(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "text", customName: "ip", Title: `title="` + Lang.T(D.Example, ':') + ` 192.168.0.8"`},
		permitted: permitted{Letters: true, Numbers: true, Characters: []rune{'.', ':'}, Minimum: 7, Maximum: 39, ExtraValidation: func(value string) error {
			var ipV string

			if value == "0.0.0.0" {
				return Lang.Err(D.Example, "IP", D.NotAllowed, ':', "0.0.0.0")
			}

			if strings.Contains(value, ":") { //IPv6
				ipV = ":"
			} else if strings.Contains(value, ".") { //IPv4
				ipV = "."
			}

			part := strings.Split(value, ipV)

			if ipV == "." && len(part) != 4 {
				return Lang.Err(D.Format, "IPv4", D.NotValid)
			}

			if ipV == ":" && len(part) != 8 {
				return Lang.Err(D.Format, "IPv6", D.NotValid)
			}

			return nil
		}},
		testData: testData{
			Good: []string{"120.1.3.206", "195.145.149.184", "179.183.230.16", "253.70.9.26", "215.35.117.51", "212.149.243.253", "126.158.214.250", "49.122.253.195", "53.218.195.25", "190.116.115.117", "115.186.149.240", "163.95.226.221"},
			Bad:  []string{"0.0.0.0", "192.168.1.1.8"},
		},
	}
	new.Set(params)
	return new
}

func ID(params ...any) *input {
	new := &input{
		attributes: attributes{allowSkipCompleted: true, htmlName: "hidden", customName: "id"},
		permitted:  permitted{Letters: false, Numbers: true, Characters: []rune{'.'}, Minimum: 1, Maximum: 39},
		testData: testData{
			Good: []string{
				"56988765432", "1234567", "0", "123456789", "100", "5000",
				"423456789", "31", "523756789", "10000232326263727", "29",
				"923726789", "3234567", "823456789", "29",
			},
			Bad: []string{"1-1", "-100", "h", "h1", "-1", " ", "", "#", "& ", "% &"},
		},
	}
	new.Set(params)
	return new
}

func Hour(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "time", customName: "Hour", Title: `title="` + Lang.T(D.Format, D.Hour, ':') + ` HH:MM"`},
		permitted:  permitted{Numbers: true, Characters: []rune{':'}, Minimum: 5, Maximum: 5, TextNotAllowed: []string{"24:"}},
		testData: testData{
			Good: []string{"23:59", "00:00", "12:00", "13:17", "21:53", "00:40", "08:30", "12:00", "15:01"},
			Bad:  append([]string{"24:00", "13-34"}, wrong_data...),
		},
	}
	new.Set(params)
	return new
}

func FilePath(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "file", customName: "FilePath"},
		permitted: permitted{
			Letters: true, Tilde: false, Numbers: true, Characters: []rune{'\\', '/', '.', '_'}, Minimum: 1, Maximum: 100, StartWith: &permitted{Letters: true, Numbers: true, Characters: []rune{'.', '_', '/'}},
		},
		testData: testData{Good: []string{".\\misArchivos", ".\\todos\\videos"}, Bad: []string{"\\-", "///."}}}
	new.Set(params)
	return new
}

func MonthDay(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "text", customName: "MonthDay"},
		permitted: permitted{Numbers: true, Minimum: 2, Maximum: 2, ExtraValidation: func(value string) error {
			_, err := validateDay(value)
			return err
		}},
		testData: testData{
			Good: []string{"01", "30", "03", "22", "31", "29", "10", "12", "05"},
			Bad:  append([]string{"1-1", "21/12"}, wrong_data...),
		},
	}
	new.Set(params)
	return new
}
