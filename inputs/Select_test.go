package inputs

import (
	"log"
	"testing"
)

var (
	modelSelect = Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")

	dataSelect = map[string]struct {
		inputData string

		expected string
	}{
		"una credencial ok?":  {"1", ""},
		"otro numero ok?":     {"3", ""},
		"0 existe?":           {"0", "valor 0 no permitido en select credentials"},
		"-1 valido?":          {"-1", "valor -1 no permitido en select credentials"},
		"carácter permitido?": {"%", "valor % no permitido en select credentials"},
		"con data?":           {"", "selección requerida campo credentials"},
		"sin espacios?":       {"luis ", "valor luis  no permitido en select credentials"},
	}
)

func Test_TagSelect(t *testing.T) {
	tag := modelSelect.BuildHtmlInput("1")
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_Select(t *testing.T) {
	for prueba, data := range dataSelect {
		t.Run((prueba + " " + data.inputData), func(t *testing.T) {
			err := modelSelect.ValidateInput(data.inputData)

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
func Test_GoodInputSelect(t *testing.T) {
	for _, data := range modelSelect.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelSelect.ValidateInput(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputSelect(t *testing.T) {
	for _, data := range modelSelect.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelSelect.ValidateInput(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
