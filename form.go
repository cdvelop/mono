package mono

import (
	"strings"
)

// attributes eg: class=myForm1, data-field_id=user.id_user.0
// default: autocomplete="off" spellcheck="false", class=form-distributed-fields
func (e *entity) FormRender(attributes ...string) string {
	if e.HtmlForm != "" {
		return e.HtmlForm
	}

	class := ` class="form-distributed-fields"`
	autocomplete := ` autocomplete="off"`
	spellcheck := ` spellcheck="false"`
	for _, a := range attributes {
		if strings.Contains(a, "class=") {
			class = ` ` + a
		}
		if strings.Contains(a, "autocomplete=") {
			autocomplete = ` ` + a
		}
		if strings.Contains(a, "spellcheck=") {
			spellcheck = ` ` + a
		}
	}

	e.HtmlForm = `<form name="` + e.Name + `"` + class + autocomplete + spellcheck + `>
	
	`
	var tabIndex int
	for _, f := range e.Fields {

		if f.Input == nil {
			continue
		}

		e.HtmlForm += f.Input.Render(&tabIndex)
		e.HtmlForm += "\n\n"
		tabIndex++
	}

	e.HtmlForm += `
	</form>`

	return e.HtmlForm
}
