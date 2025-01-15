package inputs

import (
	"log"
	"testing"
)

var (
	modelDataList = DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")

	dataList = map[string]struct {
		inputData string

		expected string
	}{
		"una credencial ok?":  {"1", ""},
		"otro numero ok?":     {"3", ""},
		"0 existe?":           {"0", "valor 0 no permitido en datalist credentials"},
		"-1 valido?":          {"-1", "valor -1 no permitido en datalist credentials"},
		"carácter permitido?": {"%", "valor % no permitido en datalist credentials"},
		"con data?":           {"", "selección requerida campo credentials"},
		"sin espacios?":       {"luis ", "valor luis  no permitido en datalist credentials"},
	}
)

func Test_DataList(t *testing.T) {
	for prueba, data := range dataList {
		t.Run((prueba + " " + data.inputData), func(t *testing.T) {
			err := modelDataList.Validate(data.inputData)

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

func Test_TagDataList(t *testing.T) {
	tag := modelDataList.Render(1)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}
