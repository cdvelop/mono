package inputs

import (
	"strconv"
	"strings"
)

type rut struct {
	hideTyping bool
	input
}

// typing="hide" hides information while typing
func Rut(params ...any) *rut {
	new := &rut{
		input: input{
			attributes: attributes{
				htmlName:     "text",
				customName:   "rut",
				Autocomplete: `autocomplete="off"`,
				Title:        `title="rut sin puntos y con guion ejem.: 11222333-4"`,
				Class:        []className{"rut"},
			},
			permitted: permitted{
				Letters:    true,
				Numbers:    true,
				Minimum:    9,
				Maximum:    11,
				Characters: []rune{'-'},
				StartWith: &permitted{
					Numbers: true,
				},
				ExtraValidation: rut{},
			},
		},
	}

	new.Set(params)

	if !new.hideTyping {
		new.PlaceHolder = `placeholder="ej: 11222333-4"`
	}

	return new
}

// RUT validate formato "7863697-1"
func (r rut) ExtraValidation(value string) error {

	// Validar RUT chileno
	if !strings.Contains(value, "-") {
		return Lang.Err(D.HyphenMissing)
	}

	data, onlyRun, err := RunData(value)
	if err != nil {
		return err
	}

	if data[0][0:1] == "0" {
		return Lang.Err(D.DoNotStartWith, D.Digit, 0)
	}

	dv := DvRut(onlyRun)

	expectedDV := strings.ToLower(data[1])
	if dv != expectedDV {
		return Lang.Err(D.Digit, D.Verifier, expectedDV, D.NotValid)
	}

	return nil
}

// DvRut retorna dígito verificador de un run
func DvRut(rut int) string {
	var sum = 0
	var factor = 2
	for ; rut != 0; rut /= 10 {
		sum += rut % 10 * factor
		if factor == 7 {
			factor = 2
		} else {
			factor++
		}
	}

	if val := 11 - (sum % 11); val == 11 {
		return "0"
	} else if val == 10 {
		return "k"
	} else {
		return strconv.Itoa(val)
	}
}

func RunData(runIn string) (data []string, onlyRun int, err error) {

	if len(runIn) < 3 {
		return nil, 0, Lang.Err(D.Value, D.Empty)
	}

	// Separar número y dígito verificador
	data = strings.Split(runIn, "-")
	if len(data) != 2 {
		return nil, 0, Lang.Err(D.Format, D.NotValid)
	}

	// Validar caracteres del número
	if !isDigits(data[0]) {
		return nil, 0, Lang.Err(D.Chars, D.NotAllowed, D.In, D.Numbers)
	}

	// Validar dígito verificador
	dv := strings.ToLower(data[1])
	if len(dv) != 1 || (dv != "k" && !isDigits(dv)) {
		return nil, 0, Lang.Err(D.Digit, D.Verifier, dv, D.NotValid)
	}

	// Convertir número a entero
	onlyRun, err = strconv.Atoi(data[0])
	if err != nil {
		return nil, 0, Lang.Err(D.Numbers, D.NotValid)
	}

	return data, onlyRun, nil
}

func isDigits(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func (r rut) GoodTestData() (out []string) {
	ok_rut := []string{"22537160-1", "5008452-3", "10493788-8", "21821424-k", "15890022-k", "7467499-2", "21129619-4", "24287548-6", "15093641-1", "10245390-5"}
	return ok_rut
}

func (r rut) WrongTestData() (out []string) {

	out = []string{"7863697-k", "7863697-0", "14080717-0", "07863697-1", " - 100 ", "-100"}
	out = append(out, wrong_data...)

	return
}
