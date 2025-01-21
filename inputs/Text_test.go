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
		"no tilde ":                          {"peréz del rozal", "é tilde no permitida"},
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

func Test_TagText(t *testing.T) {
	result := modelText.Render(1)
	if result == "" {
		t.Fatal("ERROR NO TAG RENDERING ")
	}
	id := modelText.input.Id

	expected := `<input type="hidden" name="full_name" id="` + id + `" data-price="100">`

	if result != expected {
		t.Fatalf("error:\n-result: \n%v\n\n-expected: \n%v\n", result, expected)
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

func Test_InputTextNumCode(t *testing.T) {
	var (
		modelTextNumCode = TextNumberCode()

		dataTextNumCode = map[string]struct {
			inputData string

			expected string
		}{
			"2 letras un numero ":            {"et1", ""},
			"código venta ":                  {"V22400", ""},
			"código numero y letras":         {"12f", ""},
			"guion bajo permitido?":          {"son_24_botellas", ""},
			"espacio permitido?":             {"1os cuatro", "espacio en blanco no permitido"},
			"palabras guion_bajo si? ":       {"son_2_cuadros", ""},
			"palabras separadas si?":         {"son 2 cuadros", "espacio en blanco no permitido"},
			"palabras guion medio si?":       {"son-2-cuadros", ""},
			"solo texto ok":                  {"tres", ""},
			"friday ok":                      {"friday", ""},
			"saturday ok":                    {"saturday", ""},
			"wednesday ok":                   {"Wednesday", ""},
			"month 10 ok":                    {"10", ""},
			"month 03 ok":                    {"03", ""},
			"solo un carácter":               {"3", "tamaño mínimo 2 caracteres"},
			"guion al inicio ? 2 caracteres": {"-1", "no debe comenzar con -"},
			"/ al inicio 2 caracteres":       {"/1", "no debe comenzar con /"},
		}
	)
	for prueba, data := range dataTextNumCode {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelTextNumCode.Validate(data.inputData)
			var err_str string
			if err != nil {
				err_str = err.Error()
			}

			if err_str != data.expected {
				t.Fatalf("resultado: [%v] expectativa: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}
func Test_TagTextNumCode(t *testing.T) {
	tag := TextNumberCode().Render(1)
	if tag == "" {
		t.Fatal("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputTextNumCode(t *testing.T) {
	for _, data := range TextNumberCode().GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := TextNumberCode().Validate(data); ok != nil {
				t.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputTextNumCode(t *testing.T) {
	for _, data := range TextNumberCode().WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := TextNumberCode().Validate(data); ok == nil {
				t.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_InputTextOnly(t *testing.T) {
	var (
		modelTextOnly = TextOnly()

		dataTextOnly = map[string]struct {
			inputData string

			expected string
		}{
			"nombre correcto con punto?":       {"Dr. Maria Jose Diaz Cadiz", "carácter . no permitido"},
			"palabras con tilde?":              {"María Jose Diáz Cadíz", "carácter í con tilde no permitida"},
			"caracteres 47 ok?":                {"juan marcos antonio del rosario de las carmenes", ""},
			"tilde ok ? ":                      {"peréz del rozal", "é tilde no permitida"},
			"texto con ñ?":                     {"Ñuñez perez", ""},
			"texto correcto + 3 caracteres ":   {"juli", ""},
			"texto correcto 3 caracteres ":     {"luz", ""},
			"oración ok ":                      {"hola que tal", ""},
			"Dato numérico 100 no permitido? ": {"100", "carácter 1 no permitido"},
			"con caracteres y coma ?":          {"los,true, vengadores", "carácter , no permitido"},
			"sin data ok":                      {"", "tamaño mínimo 3 caracteres"},
			"un carácter numérico ?":           {"8", "tamaño mínimo 3 caracteres"},
			"palabra mas numero permitido ?":   {"son 4 bidones", "carácter 4 no permitido"},
			"con paréntesis y numero ?":        {"son {4 bidones}", "carácter { no permitido"},
			"con solo paréntesis ?":            {"son (bidones)", "carácter ( no permitido"},
			"palabras y numero ?":              {"apellido Actualizado 1", "carácter 1 no permitido"},
			"un carácter ok?":                  {"!", "tamaño mínimo 3 caracteres"},
		}
	)
	for prueba, data := range dataTextOnly {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelTextOnly.Validate(data.inputData)

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

func Test_TagTextOnly(t *testing.T) {
	tag := TextOnly().Render(1)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputTextOnly(t *testing.T) {
	for _, data := range TextOnly().GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := TextOnly().Validate(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputTextOnly(t *testing.T) {
	for _, data := range TextOnly().WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := TextOnly().Validate(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
