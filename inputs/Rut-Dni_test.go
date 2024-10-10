package inputs

import (
	"log"
	"testing"
)

var (
	modelDNI = Rut("dni-mode")

	dataDNI = map[string]struct {
		inputData string
		expected  string
	}{
		"ok 17734478-8":               {"17734478-8", ""},
		"ok 7863697-1":                {"7863697-1", ""},
		"ok 20373221-K":               {"20373221-k", ""},
		"run validado W? permitido?":  {"7863697-W", "dígito verificador W inválido"},
		"cambio dígito a k 7863697-k": {"7863697-k", "dígito verificador k inválido"},
		"cambio dígito a 0 7863697-0": {"7863697-0", "dígito verificador 0 inválido"},
		"ok 14080717-6":               {"14080717-6", ""},
		"incorrecto 14080717-0":       {"14080717-0", "dígito verificador 0 inválido"},
		"correcto cero al inicio? ":   {"07863697-1", errCeroRut},
		"data correcta solo espacio?": {" ", "tamaño mínimo 9 caracteres"},
		"caracteres permitidos?":      {`%$"1 uut4%%oo`, "% no es un numero"},
		"pasaporte ax001223b ok?":     {"ax001223b", ""},
		"caída con dato":              {"123", "tamaño mínimo 9 caracteres"},
	}
)

func Test_TagDNI(t *testing.T) {
	tag := modelDNI.Render("1")
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_InputDNI(t *testing.T) {
	for prueba, data := range dataDNI {
		t.Run((prueba), func(t *testing.T) {
			err := modelDNI.Validate(data.inputData)

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

func Test_GoodInputDNI(t *testing.T) {
	for _, data := range modelDNI.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelDNI.Validate(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputDNI(t *testing.T) {
	for _, data := range modelDNI.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelDNI.Validate(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
