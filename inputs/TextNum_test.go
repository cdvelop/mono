package inputs

import (
	"log"
	"testing"
)

func Test_InputTextNum(t *testing.T) {
	var (
		modelTextNum = TextNumber()

		dataTextNum = map[string]struct {
			inputData string

			expected string
		}{
			"guion bajo ":           {"son_24_botellas", ""},
			"frase con guion bajo ": {"los_cuatro", ""},
			"frase sin guion bajo ": {"los cuatro", "espacio en blanco no permitido"},
			"palabras guion bajo ":  {"son_2_cuadros", ""},
			"palabras separadas ":   {"son 2 cuadros", "espacio en blanco no permitido"},
			"palabras guion medio ": {"son-2-cuadros", "carácter - no permitido"},
			"menos de 5 palabras ":  {"tres", "tamaño mínimo 5 caracteres"},
			"2 letras un numero ":   {"et1_", "tamaño mínimo 5 caracteres"},
		}
	)
	for prueba, data := range dataTextNum {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelTextNum.Validate(data.inputData)

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

func Test_TagTextNum(t *testing.T) {
	tag := TextNumber().Render(1)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputTextNum(t *testing.T) {
	for _, data := range TextNumber().GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := TextNumber().Validate(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputTextNum(t *testing.T) {
	for _, data := range TextNumber().WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := TextNumber().Validate(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
