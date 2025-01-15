package inputs

type rut struct {
	hideTyping bool
	dni_mode   bool
	input
}

// typing="hide" ocultar información  al escribir
// dni-mode: acepta otro documentos
func Rut(params ...any) *rut {
	new := &rut{
		input: input{
			attributes: attributes{
				htmlName:     "text",
				customName:   "rut",
				Autocomplete: `autocomplete="off"`,
				Class:        []string{"rut"},
				DataSet:      []map[string]string{{"option": "ch"}},
			},
			permitted: permitted{
				Letters: true,
				Numbers: true,
				Minimum: 9,
				Maximum: 15,
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
		new.Title = `title="Documento Chileno (ch) o Extranjero (ex)"`
		if !new.hideTyping {
			new.PlaceHolder = `placeholder="ej: (ch) 11222333-k  /  (ex) 1b2334"`
		}

		new.Maxlength = `maxlength="15"`
	} else {
		new.Title = `title="rut sin puntos y con guion ejem.: 11222333-4"`
		if !new.hideTyping {
			new.PlaceHolder = `placeholder="ej: 11222333-4"`
		}
		new.Maxlength = `maxlength="10"`
	}

	return new
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
