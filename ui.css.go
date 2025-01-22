package mono

type className string

const (
	cssClassNormal    className = "normal"
	cssClassBorder    className = "border"
	cssClassAllWidth  className = "all-width"
	cssClassWidthAuto className = "width-auto"
)

// func (h *input) setCssClasses() string {

// 	h.cssClasses = []*css.class{}

// 	var class = `normal border`

// 	switch h.attributes.htmlName {

// 	case "checkbox", "textarea":

// 		h.cssClasses = append(h.cssClasses, &css.class{
// 			Name: "all-width",
// 			Properties: []string{
// 				"width: 100%",
// 				"min-width: 100%",
// 			},
// 		})

// 	default:
// 		h.cssClasses = append(h.cssClasses, &css.class{
// 			Name: "width-auto",
// 			Properties: []string{
// 				"max-height: 5em",
// 			},
// 		})

// 		class += ` width-auto`
// 	}

// 	return ` class="` + class + `"`
// }
