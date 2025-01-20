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

	for _, doc := range r.options {
		if doc == "ex" {
			err := r.dni.Validate(value)
			if err != nil && r.hideTyping {
				return errors.New(hidden_err)
			}
			return err
		} else {
			err := r.runValidate(value)
			if err != nil && r.hideTyping {
				return errors.New(hidden_err)
			}
			return err
		}
	}

	if r.dni_mode {
		if !strings.Contains(value, `-`) {
			err := r.Validate(value)
			if err != nil && r.hideTyping {
				return errors.New(hidden_err)
			}
			return err
		}
	}

	err := r.runValidate(value)
	if err != nil && r.hideTyping {
		return errors.New(hidden_err)
	}

	return err
}

const errCeroRut = "primer dígito no puede ser 0"

// RUT validate formato "7863697-1"
func (r rut) runValidate(rin string) error {
	data, onlyRun, err := RunData(rin)
	if err != "" {
		return errors.New(err)
	}

	if data[0][0:1] == "0" {
		return errors.New(errCeroRut)
	}

	dv := DvRut(onlyRun)

	if dv != strings.ToLower(data[1]) {
		return errors.New("dígito verificador " + data[1] + " inválido")
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

func RunData(runIn string) (data []string, onlyRun int, err string) {

	if len(runIn) < 3 {
		return nil, 0, errRut01
	}

	if !strings.Contains(runIn, "-") {
		return nil, 0, errGuionRut
	}

	data = strings.Split(string(runIn), "-")
	// fmt.Println("TAMAÑO", len(data), "RUT DATA -:", data)
	var e error
	onlyRun, e = strconv.Atoi(data[0])
	if e != nil {
		err = "caracteres no permitidos"
	}

	return
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
