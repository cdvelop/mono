package css

import (
	"reflect"
	"strings"
	"unicode"
)

type cssVars struct {
	/* Font Sizes */
	FontSizeNormal string
	FontSizeSmall  string
	/* Colors */
	ColorPrimary    string
	ColorSecondary  string
	ColorTertiary   string
	ColorQuaternary string
	ColorGray       string
	ColorSelection  string
	ColorHover      string
	ColorSuccess    string
	ColorError      string
	/* Layout Sizes */
	MenuSize      string
	ContentHeight string
	ContentWidth  string
	/* Timing */
	TransitionWait string

	externalVars map[string]string
}

var Var = cssVars{
	FontSizeNormal:  "1.1rem",
	FontSizeSmall:   ".6rem",
	ColorPrimary:    "#ffffff",
	ColorSecondary:  "#3f88bf",
	ColorTertiary:   "#c2c1c1",
	ColorQuaternary: "#000000",
	ColorGray:       "#e9e9e9",
	ColorSelection:  "#ff9300",
	ColorHover:      "#ff95008e",
	ColorSuccess:    "#aadaff7c",
	ColorError:      "#f20707",
	MenuSize:        "6vh",
	ContentHeight:   "94vh",
	ContentWidth:    "100vw",
	TransitionWait:  "0s",

	externalVars: map[string]string{},
}

// agregar una variable CSS externa
func (c *cssVars) AddVariable(name, value string) {
	c.externalVars[name] = value
}

// obtener el valor de una variable (en formato CSS: var(--miVar))
func GetVariable(name string) string {
	return `var(` + name + `)`
}

// GenerateRoot genera la clase ":root" con todas las variables CSS
func GenerateRoot() string {
	var sb strings.Builder
	sb.WriteString(":root {\n")

	// Usa reflect para iterar sobre los campos de cssVars
	v := reflect.ValueOf(Var)
	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Name
		fieldValue := v.Field(i).String()

		// skip if it starts with lowercase
		if !unicode.IsUpper([]rune(fieldName)[0]) {
			continue
		}

		sb.WriteString("    --" + fieldName + ": " + fieldValue + ";\n")
	}

	sb.WriteString("}\n")
	return sb.String()
}
