package inputs

import (
	"errors"
	"strconv"
	"strings"
)

type rut struct {
	hideTyping bool
	dni_mode   bool
	input
}

// typing="hide" hides information while typing
// dni-mode: accepts other documents
func Rut(params ...any) *rut {
	new := &rut{
		input: input{
			attributes: attributes{
				htmlName:     "text",
				customName:   "rut",
				Autocomplete: `autocomplete="off"`,
				Class:        []className{"rut"},
				DataSet:      []map[string]string{{"option": "ch"}},
			},
			permitted: permitted{
				Letters:    true,
				Numbers:    true,
				Minimum:    9,
				Maximum:    15,
				Characters: []rune{'-'},
			},
		},
	}

	new.Set(params)

	for _, opt := range params {
		if opt.(string) == "dni-mode" {
			new.dni_mode = true
			break
		}
	}

	if new.dni_mode {
		new.attributes.customName = "rutDni"
		new.Title = Lang.T("title", "Chilean Document (ch) or Foreign (ex)")
		if !new.hideTyping {
			new.PlaceHolder = Lang.T("placeholder", "e.g.: (ch) 11222333-k  /  (ex) 1b2334")
		}

		new.Maxlength = `maxlength="15"`
	} else {
		new.Title = `title="rut sin puntos y con guion ejem.: 11222333-4"`
		if !new.hideTyping {
			new.PlaceHolder = `placeholder="ej: 11222333-4"`
		}
		new.Maxlength = `maxlength="10"`
		// new.permitted.Characters = []rune{'-'}
	}

	return new
} // validación con datos de entrada
func (r rut) Validate(value string) error {
	const hidden_err = "campo invalido"

	// Limpiar espacios y convertir a minúsculas
	value = strings.TrimSpace(value)
	value = strings.ToLower(value)

	// Validar tamaño mínimo primero
	if len(value) < 9 {
		if r.hideTyping {
			return errors.New(hidden_err)
		}
		return Lang.Err(D.MinSize, 9, D.Chars)
	}

	// Validar caracteres permitidos
	if len(value) < 9 {
		if r.hideTyping {
			return errors.New(hidden_err)
		}
		return Lang.Err(D.MinSize, 9, D.Chars)
	}

	// Validar modo DNI
	if r.dni_mode {
		// Si es documento extranjero
		if len(r.options) > 0 && r.options[0]["option"] == "ex" {
			// Validar formato básico
			if len(value) < 3 {
				if r.hideTyping {
					return errors.New(hidden_err)
				}
				return errors.New(errRut01)
			}
			// Permitir letras y números sin guion
			return nil
		}

		// Si es RUT chileno, validar formato
		if !strings.Contains(value, "-") {
			if r.hideTyping {
				return errors.New(hidden_err)
			}
			return Lang.Err(D.HyphenMissing)
		}
	}

	// Validar RUT chileno
	if !r.dni_mode {
		if !strings.Contains(value, "-") {
			if r.hideTyping {
				return errors.New(hidden_err)
			}
			return Lang.Err(D.HyphenMissing)
		}
	}

	// Ejecutar validación completa
	err := r.runValidate(value)
	if err != nil {
		if r.hideTyping {
			return errors.New(hidden_err)
		}
		return err
	}

	return nil
}

// RUT validate formato "7863697-1"
func (r rut) runValidate(rin string) error {
	data, onlyRun, err := RunData(rin)
	if err != nil {
		return err
	}

	if data[0][0:1] == "0" {
		return Lang.Err(D.DoNotStartWith, D.Digit, 0)
	}

	dv := DvRut(onlyRun)

	originalDv := data[1]
	expectedDV := strings.ToLower(data[1])
	if dv != expectedDV {
		return Lang.Err(D.Digit, D.Verifier, originalDv, D.NotValid)
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

const errRut01 = "datos ingresados insuficientes"
const errGuionRut = "guion (-) dígito verificador inexistente"

func RunData(runIn string) (data []string, onlyRun int, err error) {
	// Limpiar espacios y convertir a minúsculas
	runIn = strings.TrimSpace(strings.ToLower(runIn))

	if len(runIn) < 3 {
		return nil, 0, Lang.Err(D.Value, D.Empty)
	}

	// Verificar guion
	if !strings.Contains(runIn, "-") {
		return nil, 0, Lang.Err(D.HyphenMissing)
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
	var e error
	onlyRun, e = strconv.Atoi(data[0])
	if e != nil {
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

func Dni(params ...any) *rut {
	options := "dni-mode"
	return Rut(options)
}

func (r rut) Render(tabIndex int) string {

	if r.dni_mode {

		tag := `<div class="run-type">`

		tag += r.input.Render(tabIndex)

		tag += `<div class="rut-label-container"><label class="rut-radio-label">
			<input type="radio" name="type-dni" data-name="dni-ch" value="ch" checked="checked" onchange="changeDniType(this)">
			<span title="Documento Chileno">ch</span>
		</label>
	
		<label class="rut-radio-label">
			<input type="radio" name="type-dni" data-name="dni-ex" value="ex" onchange="changeDniType(this)">
			<span title="Documento Extranjero">ex</span>
		</label>
	  </div>
    </div>`

		return tag

	} else {
		return r.input.Render(tabIndex)
	}
}

func (r rut) GoodTestData() (out []string) {

	ok_rut := []string{"22537160-1", "5008452-3", "10493788-8", "21821424-k", "15890022-k", "7467499-2", "21129619-4", "24287548-6", "15093641-1", "10245390-5"}
	if r.dni_mode {

		out = []string{"15890022-k", "ax001223b", "A0C00A389", "B0004DF678"}
		out = append(out, ok_rut...)

		return

	} else {
		return ok_rut
	}

}

func (r rut) WrongTestData() (out []string) {

	out = []string{"7863697-k", "7863697-0", "14080717-0", "07863697-1", " - 100 ", "-100"}
	out = append(out, wrong_data...)

	return
}
