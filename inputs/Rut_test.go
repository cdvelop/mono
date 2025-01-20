package inputs

import (
	"fmt"
	"log"
	"testing"
)

func Test_InputRut(t *testing.T) {
	var (
		modelRut = Rut()

		dataRut = map[string]struct {
			inputData string
			expected  string
		}{
			"without hyphen 15890022k":    {"15890022k", Lang.T(D.HyphenMissing)},
			"no hyphen 177344788":         {"177344788", Lang.T(D.HyphenMissing)},
			"ok 7863697-1":                {"7863697-1", ""},
			"ok 20373221-K":               {"20373221-k", ""},
			"valid run? allowed?":         {"7863697-W", Lang.T(D.Digit, D.Verifier, "W", D.NotValid)},
			"change digit to k 7863697-k": {"7863697-k", Lang.T(D.Digit, D.Verifier, "k", D.NotValid)},
			"change digit to 0 7863697-0": {"7863697-0", Lang.T(D.Digit, D.Verifier, 0, D.NotValid)},
			"ok 14080717-6":               {"14080717-6", ""},
			"incorrect 14080717-0":        {"14080717-0", Lang.T(D.Digit, D.Verifier, 0, D.NotValid)},
			"correct zero at start?":      {"07863697-1", Lang.T(D.DoNotStartWith, D.Digit, 0, D.NotValid)},
			"correct data only space?":    {" ", Lang.T(D.MinSize, 9, D.Chars)},
			"ok 17734478-8":               {"17734478-8", ""},
			"allowed characters?":         {`%$"1 `, Lang.T(D.Chars, D.NotAllowed)},
			"pasaporte ax001223b ok?":     {"ax001223b", ""},
			"caída con dato":              {"123", Lang.T(D.MinSize, 9, D.Chars)},
		}
	)
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
	tag := Rut().Render(1)
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
	for _, data := range Rut().GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := Rut().Validate(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputRut(t *testing.T) {
	for _, data := range Rut().WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := Rut().Validate(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
