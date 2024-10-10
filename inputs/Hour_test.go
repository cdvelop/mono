package inputs

import (
	"log"
	"testing"
)

var (
	modelHour = Hour()

	dataHour = map[string]struct {
		inputData string
		expected  string
	}{
		"correcto":    {"23:59", ""},
		"correcto 00": {"00:00", ""},
		"correcto 12": {"12:00", ""},

		"incorrecto 24":       {"24:00", "la hora 24 no existe"},
		"incorrecto sin data": {"", "tamaño mínimo 5 caracteres"},
		"incorrecto carácter": {"13-34", "carácter - no permitido"},
	}
)

func Test_InputHour(t *testing.T) {
	for prueba, data := range dataHour {
		t.Run((prueba + " " + data.inputData), func(t *testing.T) {
			err := modelHour.Validate(data.inputData)

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

func Test_TagHour(t *testing.T) {
	tag := modelHour.buildHtml("1")
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputHour(t *testing.T) {
	for _, data := range modelHour.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelHour.Validate(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputHour(t *testing.T) {
	for _, data := range modelHour.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelHour.Validate(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
