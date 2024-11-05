package inputs

import "github.com/cdvelop/godi/css"

func (h *input) setCssClasses() string {

	h.cssClasses = []*css.Class{}

	var class = `normal border`

	switch h.attributes.htmlName {

	case "checkbox", "textarea":

		h.cssClasses = append(h.cssClasses, &css.Class{
			Name: "all-width",
			Properties: []string{
				"width: 100%",
				"min-width: 100%",
			},
		})

	default:
		h.cssClasses = append(h.cssClasses, &css.Class{
			Name: "width-auto",
			Properties: []string{
				"max-height: 5em",
			},
		})

		class += ` width-auto`
	}

	return ` class="` + class + `"`
}
