package inputs

import (
	"log"
	"testing"
)

var (
	modelMail = Mail()

	dataMail = map[string]struct {
		inputData string

		expected string
	}{
		"correo normal ":   {"mi.correo@mail.com", ""},
		"correo un campo ": {"correo@mail.com", ""},
	}
)

func Test_TagMail(t *testing.T) {
	tag := modelMail.buildHtml("1")
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_InputMail(t *testing.T) {
	for prueba, data := range dataMail {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelMail.ValidateInput(data.inputData)

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

func Test_GoodInputMail(t *testing.T) {
	for _, data := range modelMail.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelMail.ValidateInput(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputMail(t *testing.T) {
	for _, data := range modelMail.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelMail.ValidateInput(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
