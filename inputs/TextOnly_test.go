package inputs

import (
	"log"
	"testing"
)

var (
	modelTextOnly = TextOnly()

	dataTextOnly = map[string]struct {
		inputData string

		expected string
	}{
		"nombre correcto con punto?":       {"Dr. Maria Jose Diaz Cadiz", "carácter . no permitido"},
		"palabras con tilde?":              {"María Jose Diáz Cadíz", "í con tilde no permitida"},
		"caracteres 47 ok?":                {"juan marcos antonio del rosario de las carmenes", ""},
		"tilde ok ? ":                      {"peréz del rozal", "é con tilde no permitida"},
		"texto con ñ?":                     {"Ñuñez perez", ""},
		"texto correcto + 3 caracteres ":   {"juli", ""},
		"texto correcto 3 caracteres ":     {"luz", ""},
		"oración ok ":                      {"hola que tal", ""},
		"Dato numérico 100 no permitido? ": {"100", "carácter 1 no permitido"},
		"con caracteres y coma ?":          {"los,true, vengadores", "carácter , no permitido"},
		"sin data ok":                      {"", "tamaño mínimo 3 caracteres"},
		"un carácter numérico ?":           {"8", "tamaño mínimo 3 caracteres"},
		"palabra mas numero permitido ?":   {"son 4 bidones", "carácter 4 no permitido"},
		"con paréntesis y numero ?":        {"son {4 bidones}", "carácter { no permitido"},
		"con solo paréntesis ?":            {"son (bidones)", "carácter ( no permitido"},
		"palabras y numero ?":              {"apellido Actualizado 1", "carácter 1 no permitido"},
		"un carácter ok?":                  {"!", "tamaño mínimo 3 caracteres"},
	}
)

func Test_TagTextOnly(t *testing.T) {
	tag := modelTextOnly.buildHtml("1")
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_InputTextOnly(t *testing.T) {
	for prueba, data := range dataTextOnly {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelTextOnly.Validate(data.inputData)

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

func Test_GoodInputTextOnly(t *testing.T) {
	for _, data := range modelTextOnly.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextOnly.Validate(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputTextOnly(t *testing.T) {
	for _, data := range modelTextOnly.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextOnly.Validate(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
