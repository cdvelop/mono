package inputs

import (
	"log"
	"testing"
)

var (
	modelNumber = Number()

	dataNumber = map[string]struct {
		inputData string

		expected string
	}{
		"numero correcto 100": {"100", ""},
		"un carácter 0":       {"0", ""},
		"un carácter 1":       {"1", ""},
		"uint64 +20 char":     {"18446744073709551615", ""},
		"uint64 +21 char?":    {"184467440737095516150", "tamaño máximo 20 caracteres"},
		"int 64 -o+ 19 char":  {"9223372036854775807", ""},
		"int 32 -o+ 10  char": {"2147483647", ""},
		"18 cifras":           {"100002323262637278", ""},

		"grande y letra":          {"10000232E26263727", "carácter E no es un numero"},
		"numero negativo -100 ":   {"-100", "carácter - no es un numero"},
		"texto en vez de numero ": {"h", "carácter h no es un numero"},
		"texto y numero":          {"h1", "carácter h no es un numero"},
	}
)

func Test_InputNumber(t *testing.T) {
	for prueba, data := range dataNumber {
		t.Run((prueba), func(t *testing.T) {
			err := modelNumber.Validate(data.inputData)

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

func Test_TagNumber(t *testing.T) {
	tag := modelNumber.Render(1)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

var (
	// 1 408 XXX XXXX
	// 5 699 524 9966

	dataPhoneNumber = map[string]struct {
		inputData string

		expected string
	}{
		"numero correcto 7 dígitos":      {"1234567", ""},
		"numero correcto 9 dígitos":      {"123456789", ""},
		"numero correcto 11 dígitos ok?": {"12345678911", ""},
		"con código país":                {"56988765432", ""},
		"signo mas + ok?":                {"+56988765432", "tamaño máximo 11 caracteres"},
		"numero correcto 6 dígitos ok?":  {"123456", "tamaño mínimo 7 caracteres"},
		"numero correcto 1 dígitos ok?":  {"0", "tamaño mínimo 7 caracteres"},
	}
)

func Test_InputPhoneNumber(t *testing.T) {
	for prueba, data := range dataPhoneNumber {
		t.Run((prueba), func(t *testing.T) {
			err := Phone().Validate(data.inputData)

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

func Test_GoodInputPhoneNumber(t *testing.T) {
	for _, data := range Phone().GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := Phone().Validate(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputNumber(t *testing.T) {
	for _, data := range modelNumber.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelNumber.Validate(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputNumber(t *testing.T) {
	for _, data := range modelNumber.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelNumber.Validate(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
