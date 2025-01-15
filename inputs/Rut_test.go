package inputs

import (
	"fmt"
	"log"
	"testing"
)

var (
	modelRut = Rut()

	dataRut = map[string]struct {
		inputData string

		expected string
	}{
		"sin guion 15890022k":         {"15890022k", errGuionRut},
		"no tiene guion 177344788":    {"177344788", errGuionRut},
		"ok 7863697-1":                {"7863697-1", ""},
		"ok 20373221-K":               {"20373221-k", ""},
		"run validado? permitido?":    {"7863697-W", "dígito verificador W inválido"},
		"cambio dígito a k 7863697-k": {"7863697-k", "dígito verificador k inválido"},
		"cambio dígito a 0 7863697-0": {"7863697-0", "dígito verificador 0 inválido"},
		"ok 14080717-6":               {"14080717-6", ""},
		"incorrecto 14080717-0":       {"14080717-0", "dígito verificador 0 inválido"},
		"correcto cero al inicio? ":   {"07863697-1", errCeroRut},
		"data correcta solo espacio?": {" ", errRut01},
		"ok 17734478-8":               {"17734478-8", ""},
		"caracteres permitidos?":      {`%$"1 `, errGuionRut},
		"no tiene guion 20373221K":    {"20373221k", errGuionRut},
	}
)

func Test_InputRut(t *testing.T) {
	for prueba, data := range dataRut {
		t.Run((prueba), func(t *testing.T) {
			err := modelRut.Validate(data.inputData)

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

func Test_TagRut(t *testing.T) {
	tag := modelRut.Render(1)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_RutDigito(t *testing.T) {
	run := 17734478
	dv := DvRut(run)
	fmt.Printf("RUN: %v DIGITO: %v", run, dv)
}

func Test_GoodInputRut(t *testing.T) {
	for _, data := range modelRut.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelRut.Validate(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputRut(t *testing.T) {
	for _, data := range modelRut.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelRut.Validate(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
