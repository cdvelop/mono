package inputs

import (
	"log"
	"testing"
)

var (
	modelCheck = CheckBox("name=credentials", "options=1:Admin,2:Editor,3:Visitante")

	datacheck = map[string]struct {
		inputData string

		expected string
	}{
		// "una credencial ok?":         {modelCheck.GoodTestData()[0], ""},
		"editor y admin ok?":         {"1,2", ""},
		"todas las credenciales ok?": {`1,3`, ""},
		"0 existe?":                  {"0", "valor 0 no permitido en checkbox credentials"},
		"-1 valido?":                 {"-1", "valor -1 no permitido en checkbox credentials"},
		"todas existentes?":          {"1,5", "valor 5 no permitido en checkbox credentials"},
		"con data?":                  {"", "selección requerida campo credentials"},
		"sin espacios?":              {"luis ,true, 3", "valor luis  no permitido en checkbox credentials"},
	}
)

func Test_check(t *testing.T) {
	for prueba, data := range datacheck {
		t.Run((prueba + " " + data.inputData), func(t *testing.T) {
			err := modelCheck.ValidateInput(data.inputData)

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

func Test_TagCheck(t *testing.T) {
	tag := modelCheck.BuildHtmlInput("1")
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}
