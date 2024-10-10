package inputs

import (
	"log"
	"testing"
)

var (
	modelTextSearch = TextSearch()

	dataTextSearch = map[string]struct {
		inputData string

		expected string
	}{
		"palabra solo texto 15 caracteres?": {"Maria Jose Diaz", ""},
		"texto con ñ ok?":                   {"Ñuñez perez", ""},
		"tilde permitido?":                  {"peréz del rozal", "é con tilde no permitida"},
		"mas de 20 caracteres permitidos?":  {"hola son mas de 21 ca", "tamaño máximo 20 caracteres"},
		"guion permitido":                   {"12038-0", ""},
		"fecha correcta?":                   {"1990-07-21", ""},
		"fecha incorrecta permitida?":       {"190-07-21", ""},
	}
)

func Test_TagTextSearch(t *testing.T) {
	tag := modelTextSearch.Render("1")
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_InputTextSearch(t *testing.T) {
	for prueba, data := range dataTextSearch {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelTextSearch.Validate(data.inputData)

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

func Test_GoodInputTextSearch(t *testing.T) {
	for _, data := range modelTextSearch.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextSearch.Validate(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputTextSearch(t *testing.T) {
	for _, data := range modelTextSearch.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextSearch.Validate(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
