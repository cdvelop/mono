package inputs

import (
	"log"
	"testing"
)

var (
	modelText = Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)

	dataText = map[string]struct {
		inputData string

		expected string
	}{
		"nombre correcto con punto?":         {"Dr. Maria Jose Diaz Cadiz", ""},
		"no tilde ":                          {"peréz del rozal", "carácter é con tilde no permitida"},
		"texto con ñ ":                       {"Ñuñez perez", ""},
		"texto correcto + 3 caracteres ":     {"hola", ""},
		"texto correcto 3 caracteres ":       {"los", ""},
		"oración ok ":                        {"hola que tal", ""},
		"solo Dato numérico permitido?":      {"100", ""},
		"con caracteres y coma ":             {"los,true, vengadores", ""},
		"sin data ok":                        {"", "tamaño mínimo 2 caracteres"},
		"un carácter numérico ":              {"8", "tamaño mínimo 2 caracteres"},
		"palabra mas numero permitido ":      {"son 4 bidones", ""},
		"con paréntesis y numero ":           {"son 4 (4 bidones)", ""},
		"con solo paréntesis ":               {"son (bidones)", ""},
		"palabras y numero":                  {"apellido Actualizado 1", ""},
		"palabra con slash?":                 {" estos son \\n los podria", "carácter \\ no permitido"},
		"nombre de archivos separados por ,": {"dino.png, gatito.jpeg", ""},
	}
)

func Test_TagText(t *testing.T) {
	result := modelText.Render(1)
	if result == "" {
		t.Fatal("ERROR NO TAG RENDERING ")
	}

	expected := `<fieldset data-name="full_name" tabindex="1"">
	<legend class="basic-legend"><label for="17"></label></legend>
	<input type="hidden" name="full_name" id="17" placeholder="tu nombre" title="permitido: letras números caracteres: ␣ . , ( ) mín. 2 máx. 100" data-price="100" data-name="text"></fieldset>`

	if result != expected {
		t.Fatalf("error:\n-result: \n%v\n\n-expected: \n%v\n", result, expected)
	}

}

func Test_InputText(t *testing.T) {
	for prueba, data := range dataText {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelText.Validate(data.inputData)

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

func Test_GoodInputText(t *testing.T) {
	for _, data := range modelText.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelText.Validate(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputTextFirsNames(t *testing.T) {
	for _, data := range modelText.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelText.Validate(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputText(t *testing.T) {
	for _, data := range modelText.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelText.Validate(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
