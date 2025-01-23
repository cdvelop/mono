package mono

import (
	"strconv"
	"strings"
)

type inputs struct{}

// html an validation inputs
var IN = inputs{}

func (i inputs) Date(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "date", Title: `title="` + R.T(D.Format, D.Date, ':') + ` DD-MM-YYYY"`},
		permitted: permitted{
			ExtraValidation: G.Date.DateExists,
		},
	}
	new.Set(params)
	return new
}

// typing="hide" hides information while typing
func (i inputs) Rut(params ...any) *input {
	new := &input{
		attributes: attributes{
			htmlName: "text", customName: "rut", Autocomplete: `autocomplete="off"`,
			Title: `title="rut sin puntos y con guion ejem.: 11222333-4"`,
			Class: []className{"rut"},
		},
		permitted: permitted{Numbers: true, Letters: true, Minimum: 9, Maximum: 11, Characters: []rune{'-'},
			StartWith: &permitted{Numbers: true},
			ExtraValidation: func(value string) error {

				// Validar RUT chileno
				if !strings.Contains(value, "-") {
					return R.Err(D.HyphenMissing)
				}

				data, onlyRun, err := runData(value)
				if err != nil {
					return err
				}

				if data[0][0:1] == "0" {
					return R.Err(D.DoNotStartWith, D.Digit, 0)
				}

				dv := G.Rut.DvRut(onlyRun)

				expectedDV := strings.ToLower(data[1])
				if dv != expectedDV {
					return R.Err(D.Digit, D.Verifier, expectedDV, D.NotValid)
				}

				return nil

			},
		},
	}

	new.Set(params)

	if !new.hideTyping {
		new.PlaceHolder = `placeholder="ej: 11222333-4"`
	}

	return new
}

// eg: options=1:Admin,2:Editor,3:Visitante
func (i inputs) DataList(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "datalist"},
	}
	new.Set(params)

	return new
}

// value="valor a mostrar"
func (i inputs) Info(params ...any) *input {
	new := &input{
		attributes: attributes{
			htmlName:   "text",
			customName: "info",
		},
	}

	return new
}

// eg: options=1:Admin,2:Editor,3:Visitante
func (i inputs) CheckBox(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "checkbox"},
	}
	new.Set(params)
	return new
}

// formato fecha: DD-MM-YYYY
func (i inputs) DateAge(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "date", customName: "DateAge", Title: `title="formato fecha: DD-MM-YYYY"`, Onchange: `onchange="DateAgeChange(this)"`},
	}
	new.Set(params)
	return new
}

// formato dia DD como palabra ej. Lunes 24 Diciembre
func (i inputs) DayWord(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "text", customName: "DayWord", DataSet: []map[string]string{{"spanish": ""}}},
	}
	new.Set(params)
	return new
}

// ej: options=m:male,f:female
func (i inputs) Radio(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "radio", Onchange: `onchange="RadioChange(this);"`},
	}
	new.Set(params)
	return new
}

// ej: {"f": "Femenino", "m": "Masculino"}.
func (i inputs) RadioGender(params ...any) *input {
	options := append(params, "name=genre", `options=f:`+R.T(D.Female)+`,m:`+R.T(D.Male)+``)
	return i.Radio(options...)
}

func (i inputs) Text(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "text"},
		permitted:  permitted{Letters: true, Tilde: false, Numbers: true, Characters: []rune{' ', '.', ',', '(', ')'}, Minimum: 2, Maximum: 100},
	}
	new.Set(params)
	return new
}

func (i inputs) TextNumberCode(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "tel", customName: "textNumberCode"},
		permitted:  permitted{Letters: true, Numbers: true, Characters: []rune{'_', '-'}, Minimum: 2, Maximum: 15, StartWith: &permitted{Letters: true, Numbers: true}},
	}
	new.Set(params)
	return new
}

func (i inputs) TextOnly(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "text", customName: "textOnly"},
		permitted:  permitted{Letters: true, Minimum: 3, Maximum: 50, Characters: []rune{' '}},
	}
	new.Set(params)
	return new
}

func (i inputs) TextNumber(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "text", customName: "textNumber"},
		permitted:  permitted{Letters: true, Numbers: true, Characters: []rune{'_'}, Minimum: 5, Maximum: 20},
	}
	new.Set(params)
	return new
}

func (i inputs) TextSearch(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "search", customName: "textSearch"},
		permitted:  permitted{Letters: true, Tilde: false, Numbers: true, Characters: []rune{'-', ' '}, Minimum: 2, Maximum: 20},
	}
	new.Set(params)
	return new
}

func (i inputs) TextArea(params ...any) *input {
	new := &input{
		attributes: attributes{Rows: `rows="3"`, Cols: `cols="1"`, Oninput: `oninput="TexAreaOninput(this)"`},
		permitted:  permitted{Letters: true, Tilde: true, Numbers: true, BreakLine: true, WhiteSpaces: true, Tabulation: true, Characters: []rune{'$', '%', '+', '#', '-', '.', ',', ':', '(', ')'}, Minimum: 2, Maximum: 1000},
	}
	new.Set(params)
	return new
}

func (i inputs) Password(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "password"},
		permitted:  permitted{Letters: true, Tilde: true, Numbers: true, Characters: []rune{' ', '#', '%', '?', '.', ',', '-', '_'}, Minimum: 5, Maximum: 50},
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

func (i inputs) Number(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "number"},
		permitted:  permitted{Numbers: true, Minimum: 1, Maximum: 20},
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

func (i inputs) Phone(params ...any) *input {
	return i.Number(`min="7"`, `max="11"`)
}

func (i inputs) Mail(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "mail", PlaceHolder: `placeHolder="ej: mi.correo@mail.com"`},
		permitted: permitted{Letters: true, Numbers: true, Characters: []rune{'@', '.', '_'}, Minimum: 0, Maximum: 0, ExtraValidation: func(s string) error {
			if strings.Contains(s, "example") {
				return R.Err(D.Example, D.Email, D.NotAllowed)
			}

			parts := strings.Split(s, "@")
			if len(parts) != 2 {
				return R.Err(D.Format, D.Email, D.NotValid)
			}

			return nil
		}},
	}
	new.Set(params)
	return new
}

func (i inputs) Ip(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "text", customName: "ip", Title: `title="` + R.T(D.Example, ':') + ` 192.168.0.8"`},
		permitted: permitted{Letters: true, Numbers: true, Characters: []rune{'.', ':'}, Minimum: 7, Maximum: 39, ExtraValidation: func(value string) error {
			var ipV string

			if value == "0.0.0.0" {
				return R.Err(D.Example, "IP", D.NotAllowed, ':', "0.0.0.0")
			}

			if strings.Contains(value, ":") { //IPv6
				ipV = ":"
			} else if strings.Contains(value, ".") { //IPv4
				ipV = "."
			}

			part := strings.Split(value, ipV)

			if ipV == "." && len(part) != 4 {
				return R.Err(D.Format, "IPv4", D.NotValid)
			}

			if ipV == ":" && len(part) != 8 {
				return R.Err(D.Format, "IPv6", D.NotValid)
			}

			return nil
		}},
	}
	new.Set(params)
	return new
}

func (i inputs) ID(params ...any) *input {
	new := &input{
		attributes: attributes{allowSkipCompleted: true, htmlName: "hidden", customName: "id"},
		permitted:  permitted{Letters: false, Numbers: true, Characters: []rune{'.'}, Minimum: 1, Maximum: 39},
	}
	new.Set(params)
	return new
}

func (i inputs) Hour(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "time", customName: "Hour", Title: `title="` + R.T(D.Format, D.Hour, ':') + ` HH:MM"`},
		permitted:  permitted{Numbers: true, Characters: []rune{':'}, Minimum: 5, Maximum: 5, TextNotAllowed: []string{"24:"}},
	}
	new.Set(params)
	return new
}

func (i inputs) FilePath(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "file", customName: "FilePath"},
		permitted: permitted{
			Letters: true, Tilde: false, Numbers: true, Characters: []rune{'\\', '/', '.', '_'}, Minimum: 1, Maximum: 100, StartWith: &permitted{Letters: true, Numbers: true, Characters: []rune{'.', '_', '/'}},
		},
	}
	new.Set(params)
	return new
}

// eg: options=1:Admin,2:Editor,3:Visitante
func (i inputs) List(params ...any) *input {
	new := &input{
		attributes: attributes{
			htmlName: "ol",
		},
	}
	new.Set(params)

	return new
}

func (i inputs) MonthDay(params ...any) *input {
	new := &input{
		attributes: attributes{htmlName: "text", customName: "MonthDay"},
		permitted: permitted{Numbers: true, Minimum: 2, Maximum: 2, ExtraValidation: func(value string) error {
			_, err := G.Date.ValidateDay(value)
			return err
		}},
	}
	new.Set(params)
	return new
}

// eg: options=1:Admin,2:Editor,3:Visitante
func (i inputs) Select(params ...any) *input {
	new := input{
		attributes: attributes{
			htmlName: "select",
		},
	}
	new.Set(params)

	return &new
}
