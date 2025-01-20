package inputs

import (
	"log"
	"testing"
)

func Test_InputHour(t *testing.T) {
	var (
		modelHour = Hour()

		dataHour = map[string]struct {
			inputData string
			expected  string
		}{
			"correct":    {"23:59", ""},
			"correct 00": {"00:00", ""},
			"correct 12": {"12:00", ""},

			"incorrect 24":        {"24:00", Lang.T(D.NotAllowed, ':', "24:")},
			"incorrect no data":   {"", Lang.T(D.Field, D.Empty, D.NotAllowed)},
			"incorrect character": {"13-34", Lang.T(D.Char, "-", D.NotAllowed)},
		}
	)
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
	modelHour := Hour()
	tag := modelHour.Render(1)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputHour(t *testing.T) {
	modelHour := Hour()
	for _, data := range modelHour.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelHour.Validate(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputHour(t *testing.T) {
	modelHour := Hour()
	for _, data := range modelHour.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelHour.Validate(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
