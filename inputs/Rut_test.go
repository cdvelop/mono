package inputs

import (
	"fmt"
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
			"ok uppercase K 20373221-K":   {"20373221-k", ""},
			"ok lowercase k 20373221-k":   {"20373221-k", ""},
			"valid run? allowed?":         {"7863697-W", Lang.T(D.Digit, D.Verifier, "w", D.NotValid)},
			"change digit to k 7863697-k": {"7863697-k", Lang.T(D.Digit, D.Verifier, "k", D.NotValid)},
			"change digit to 0 7863697-0": {"7863697-0", Lang.T(D.Digit, D.Verifier, 0, D.NotValid)},
			"ok 14080717-6":               {"14080717-6", ""},
			"incorrect 14080717-0":        {"14080717-0", Lang.T(D.Digit, D.Verifier, 0, D.NotValid)},
			"correct zero at start?":      {"07863697-1", Lang.T(D.DoNotStartWith, D.Digit, 0)},
			"correct data only space?":    {" ", Lang.T(D.DoNotStartWith, D.WhiteSpace)},
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
				fmt.Println(prueba)
				t.Fatalf("result: [%v] expected: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}

func Test_TagRut(t *testing.T) {
	tag := Rut().Render(1)
	if tag == "" {
		t.Fatal("ERROR NO TAG RENDERING ")
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
				t.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputRut(t *testing.T) {
	for _, data := range Rut().WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := Rut().Validate(data); ok == nil {
				t.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
