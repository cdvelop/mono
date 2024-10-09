package inputs

import (
	"log"
	"testing"
)

var (
	modelPassword = Password()

	dataPassword = map[string]struct {
		inputData string

		expected string
	}{
		"valida numero letras y carácter": {"c0ntra3", ""},
		"valida muchos caracteres":        {"M1 contraseÑ4", ""},
		"valida 8 caracteres":             {"contrase", ""},
		"valida 5 caracteres":             {"contñ", ""},
		"valida solo números":             {"12345", ""},
		"no valida menos de 2":            {"1", "tamaño mínimo 5 caracteres"},
		"sin data":                        {"", "tamaño mínimo 5 caracteres"},
	}
)

func Test_InputPassword(t *testing.T) {
	for prueba, data := range dataPassword {
		t.Run((prueba + ": " + data.inputData), func(t *testing.T) {
			err := modelPassword.ValidateInput(data.inputData)

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

func Test_TagPassword(t *testing.T) {
	tag := modelPassword.buildHtml("1")
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputPassword(t *testing.T) {
	for _, data := range modelPassword.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelPassword.ValidateInput(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputPassword(t *testing.T) {
	for _, data := range modelPassword.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelPassword.ValidateInput(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

var modelPasswordMinimal = Password(`min="10"`, `max="30"`)

func Test_GoodInputPasswordMinimal(t *testing.T) {
	for _, data := range modelPasswordMinimal.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelPasswordMinimal.ValidateInput(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputPasswordMinimal(t *testing.T) {
	for _, data := range modelPasswordMinimal.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelPasswordMinimal.ValidateInput(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
