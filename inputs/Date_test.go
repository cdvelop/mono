package inputs

import (
	"log"
	"testing"
)

var (
	modelDate = Date()

	dataDate = map[string]struct {
		inputData string
		expected  string
	}{
		"correcto ":                      {"2002-12-03", ""},
		"dia 29 febrero año bisiesto":    {"2020-02-29", ""},
		"dia 29 febrero año no bisiesto": {"2023-02-29", "Febrero no contiene 29 días. año 2023 no es bisiesto."},
		"junio no tiene 31 días":         {"2023-06-31", "Junio no contiene 31 días."},
		"carácter de mas incorrecto ":    {"2002-12-03-", "CheckDateExists formato de fecha ingresado incorrecto ej: 2006-01-02"},
		"formato incorrecto ":            {"21/12/1998", "CheckDateExists formato de fecha ingresado incorrecto ej: 2006-01-02"},
		"fecha incorrecta ":              {"2020-31-01", "CheckDateExists error 31 es un formato de mes incorrecto ej: 01 a 12"},
		"fecha recortada sin año ok?":    {"31-01", "CheckDateExists formato de fecha ingresado incorrecto ej: 2006-01-02"},
		"data incorrecta ":               {"0000-00-00", "CheckDateExists error 00 es un formato de mes incorrecto ej: 01 a 12"},
		"toda la data correcta?":         {"", "CheckDateExists formato de fecha ingresado incorrecto ej: 2006-01-02"},
	}
)

func Test_InputDate(t *testing.T) {
	for prueba, data := range dataDate {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelDate.Validate(data.inputData)

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

func Test_TagDate(t *testing.T) {
	tag := modelDate.buildHtml("1")
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputDate(t *testing.T) {
	for _, data := range modelDate.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelDate.Validate(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputDate(t *testing.T) {
	for _, data := range modelDate.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelDate.Validate(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
