package inputs

import (
	"log"
	"testing"
)

var (
	modelMonthDay = MonthDay()

	dataMonthDay = map[string]struct {
		inputData string

		expected string
	}{
		"dia ok?":                     {"01", ""},
		"caracteres?":                 {"0l", "l no es un numero"},
		"mes ok?":                     {"31", ""},
		"fecha recortada sin año ok?": {"31-01", "tamaño máximo 2 caracteres"},
		"correcto ?":                  {"1-1", "tamaño máximo 2 caracteres"},
		"incorrecto ":                 {"2002-12-03", "tamaño máximo 2 caracteres"},
		"formato incorrecto ":         {"21/12", "tamaño máximo 2 caracteres"},
		"data incorrecta ":            {"0000-00-00", "tamaño máximo 2 caracteres"},
		"toda la data correcta?":      {"", "tamaño mínimo 2 caracteres"},
	}
)

func Test_InputMonthDay(t *testing.T) {
	for prueba, data := range dataMonthDay {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelMonthDay.Validate(data.inputData)

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

func Test_TagMonthDay(t *testing.T) {
	tag := modelMonthDay.Render("1")
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputMonthDay(t *testing.T) {
	for _, data := range modelMonthDay.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelMonthDay.Validate(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputMonthDay(t *testing.T) {
	for _, data := range modelMonthDay.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelMonthDay.Validate(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
