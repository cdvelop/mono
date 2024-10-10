package inputs

import (
	"log"
	"strconv"
	"testing"
)

var (
	modelTextArea = TextArea()

	dataTextArea = map[string]struct {
		inputData string
		expected  string
	}{
		"todo los caracteres permitidos?":   {"hola: esto, es. la - prueba 10", ""},
		"salto de linea permitido? y guion": {"hola:\n esto, es. la - \nprueba 10", ""},
		"letra ñ permitida? paréntesis y $": {"soy ñato o Ñato (aqui) costo $10000.", ""},
		"solo texto y espacio?":             {"hola esto es una prueba", ""},
		"texto y puntuación?":               {"hola: esto es una prueba", ""},
		"texto y puntuación y coma?":        {"hola: esto,true, es una prueba", ""},
		"5 caracteres?":                     {", .s5", ""},
		"sin data permitido?":               {"", "tamaño mínimo " + strconv.Itoa(modelTextArea.Minimum) + " caracteres"},
		"# permitido?":                      {"# son", ""},
		"¿ ? permitido?":                    {" ¿ si ?", "carácter ¿ no permitido"},
		"tildes si?":                        {" mí tílde", ""},
		"1 carácter?":                       {"1", "tamaño mínimo " + strconv.Itoa(modelTextArea.Minimum) + " caracteres"},
		"nombre correcto?":                  {"Dr. Pato Gomez", ""},
		"solo espacio en blanco?":           {" ", "tamaño mínimo " + strconv.Itoa(modelTextArea.Minimum) + " caracteres"},
		"texto largo correcto?":             {`IRRITACION EN PIEL DE ROSTRO. ALERGIAS NO. CIRUGIAS NO. ACTUAL TTO CON ISOTRETINOINA 10MG - ENERO 2022. EN TTO ACTUAL CON VIT D. EXAMEN DE LAB 20-12-2022. SIN OTROS ANTECEDENTES`, ""},
		"texto con salto de lineas ok": {`HOY......Referido por        : dr. ........
		Motivo                    : ........
		Premedicacion  : ........`, ""},
	}
)

func Test_InputTextArea(t *testing.T) {
	for prueba, data := range dataTextArea {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelTextArea.Validate(data.inputData)

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

func Test_TagTextArea(t *testing.T) {
	tag := modelTextArea.buildHtml("1")
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}
func Test_GoodInputTextArea(t *testing.T) {
	for _, data := range modelTextArea.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextArea.Validate(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputTextArea(t *testing.T) {
	for _, data := range modelTextArea.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelTextArea.Validate(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
