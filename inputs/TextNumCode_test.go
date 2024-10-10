package inputs

import (
	"log"
	"testing"
)

var (
	modelTextNumCode = TextNumberCode()

	dataTextNumCode = map[string]struct {
		inputData string

		expected string
	}{
		"2 letras un numero ":            {"et1", ""},
		"código venta ":                  {"V22400", ""},
		"código numero y letras":         {"12f", ""},
		"guion bajo permitido?":          {"son_24_botellas", ""},
		"espacio permitido?":             {"1os cuatro", "espacios en blanco no permitidos"},
		"palabras guion_bajo si? ":       {"son_2_cuadros", ""},
		"palabras separadas si?":         {"son 2 cuadros", "espacios en blanco no permitidos"},
		"palabras guion medio si?":       {"son-2-cuadros", ""},
		"solo texto ok":                  {"tres", ""},
		"friday ok":                      {"friday", ""},
		"saturday ok":                    {"saturday", ""},
		"wednesday ok":                   {"Wednesday", ""},
		"month 10 ok":                    {"10", ""},
		"month 03 ok":                    {"03", ""},
		"solo un carácter":               {"3", "tamaño mínimo 2 caracteres"},
		"guion al inicio ? 2 caracteres": {"-1", "no se puede comenzar con -"},
		"/ al inicio 2 caracteres":       {"/1", "no se puede comenzar con /"},
	}
)

func Test_InputTextNumCode(t *testing.T) {
	for prueba, data := range dataTextNumCode {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelTextNumCode.Validate(data.inputData)

			var err_str string
			if err != nil {
				err_str = err.Error()
			}

			if err_str != data.expected {
				log.Println(prueba)
				log.Fatalf("resultado: [%v] expectativa: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}
func Test_TagTextNumCode(t *testing.T) {
	tag := modelTextNumCode.buildHtml("1")
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputTextNumCode(t *testing.T) {
	for _, data := range modelTextNumCode.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextNumCode.Validate(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputTextNumCode(t *testing.T) {
	for _, data := range modelTextNumCode.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextNumCode.Validate(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
