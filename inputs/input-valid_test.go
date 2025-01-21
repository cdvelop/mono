package inputs

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

// DataList tests
func Test_DataList(t *testing.T) {
	var (
		modelDataList = DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")

		dataList = map[string]struct {
			inputData string
			expected  string
		}{
			"una credencial ok?":  {"1", ""},
			"otro numero ok?":     {"3", ""},
			"0 existe?":           {"0", "valor 0 no permitido en datalist credentials"},
			"-1 valido?":          {"-1", "valor -1 no permitido en datalist credentials"},
			"carácter permitido?": {"%", "valor % no permitido en datalist credentials"},
			"con data?":           {"", "selección requerida campo credentials"},
			"sin espacios?":       {"luis ", "valor luis  no permitido en datalist credentials"},
		}
	)
	for prueba, data := range dataList {
		t.Run((prueba + " " + data.inputData), func(t *testing.T) {
			err := modelDataList.Validate(data.inputData)

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

// Date tests
func Test_InputDate(t *testing.T) {
	var (
		modelDate = Date()

		dataDate = map[string]struct {
			inputData string
			expected  string
		}{
			"correct ":                        {"2002-12-03", ""},
			"February 29th leap year":         {"2020-02-29", ""},
			"February 29th non leap year":     {"2023-02-29", Lang.T(D.February, D.DoesNotHave, "29", D.Days, D.Year, "2023")},
			"June does not have 31 days":      {"2023-06-31", Lang.T(D.June, D.DoesNotHave, "31", D.Days)},
			"incorrect extra character ":      {"2002-12-03-", Lang.T(D.InvalidDateFormat, "2006-01-02")},
			"incorrect format ":               {"21/12/1998", Lang.T(D.InvalidDateFormat, "2006-01-02")},
			"incorrect month 31":              {"2020-31-01", Lang.T(D.Month, 31, D.NotValid)},
			"shortened date without year ok?": {"31-01", Lang.T(D.InvalidDateFormat, "2006-01-02")},
			"incorrect data ":                 {"0000-00-00", Lang.T(D.InvalidDateFormat)},
			"all data correct?":               {"", Lang.T(D.InvalidDateFormat, "2006-01-02")},
		}
	)

	for prueba, data := range dataDate {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelDate.Validate(data.inputData)

			var err_str string
			if err != nil {
				err_str = err.Error()
			}

			if err_str != data.expected {
				log.Println(prueba)
				t.Fatalf("result: [%v] expected: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}

// FilePath tests
func Test_FilePath(t *testing.T) {
	var (
		filePathTestData = map[string]struct {
			inputData string
			expected  string
		}{
			"correct path":                                {".\\files\\1234\\", ""},
			"correct path without starting point?":        {"\\files\\1234\\", Lang.Err(D.DoNotStartWith, '\\').Error()},
			"absolute path in Linux":                      {"/home/user/files/", ""},
			"absolute path without starting point":        {"/files/1234/", ""},
			"relative path without directories ok?":       {".\\", ""},
			"relative path with final slash":              {"./files/1234/", ""},
			"path with filename":                          {".\\files\\1234\\archivo.txt", ""},
			"path with directory names using underscores": {".\\mi_directorio\\sub_directorio\\", ""},
			"is a number a valid path?":                   {"5", ""},
			"is a single word a valid path?":              {"path", ""},
			"white spaces allowed?":                       {".\\path with white space\\", Lang.T(D.WhiteSpace, D.NotAllowed)},
		}
	)

	for prueba, data := range filePathTestData {
		t.Run((prueba + " " + data.inputData), func(t *testing.T) {
			err := FilePath().Validate(data.inputData)

			var err_str string
			if err != nil {
				err_str = err.Error()
			}

			if err_str != data.expected {
				log.Println(prueba)
				log.Fatalf("got: [%v] expected: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}

// IP tests
func Test_InputIp(t *testing.T) {
	var (
		dataIp = map[string]struct {
			inputData string
			expected  string
		}{
			"IPv4 ok":        {"192.168.1.1", ""},
			"IPv6 ok":        {"2001:0db8:85a3:0000:0000:8a2e:0370:7334", ""},
			"ip incorrecta ": {"192.168.1.1.8", Lang.T(D.Format, "IPv4", D.NotValid)},
			"correcto?":      {"0.0.0.0", Lang.T(D.Example, "IP", D.NotAllowed, ':', "0.0.0.0")},
			"sin data ":      {"", Lang.T(D.Field, D.Empty, D.NotAllowed)},
		}
	)
	for prueba, data := range dataIp {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := Ip().Validate(data.inputData)

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

// Mail tests
func Test_InputMail(t *testing.T) {
	var (
		modelMail = Mail()

		dataMail = map[string]struct {
			inputData string
			expected  string
		}{
			"correo normal ":   {"mi.correo@mail.com", ""},
			"correo un campo ": {"correo@mail.com", ""},
		}
	)
	for prueba, data := range dataMail {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelMail.Validate(data.inputData)

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

// Number tests
func Test_InputNumber(t *testing.T) {
	var (
		modelNumber = Number()

		dataNumber = map[string]struct {
			inputData string
			expected  string
		}{
			"correct number 100":       {"100", ""},
			"single digit 0":           {"0", ""},
			"single digit 1":           {"1", ""},
			"uint64 20 chars":          {"18446744073709551615", ""},
			"uint64 21 chars":          {"184467440737095516150", Lang.T(D.MaxSize, 20, D.Chars)},
			"int64 19 chars":           {"9223372036854775807", ""},
			"int32 10 chars":           {"2147483647", ""},
			"18 digits":                {"100002323262637278", ""},
			"large number with letter": {"10000232E26263727", Lang.T('E', D.NotNumber)},
			"negative number -100":     {"-100", Lang.T('-', D.NotNumber)},
			"text instead of number":   {"h", Lang.T('h', D.NotNumber)},
			"text and number":          {"h1", Lang.T('h', D.NotNumber)},
		}
	)
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

// Password tests
func Test_InputPassword(t *testing.T) {
	var (
		modelPassword = Password()

		dataPassword = map[string]struct {
			inputData string
			expected  string
		}{
			"validates numbers letters and character": {"c0ntra3", ""},
			"validates many characters":               {"M1 contraseÑ4", ""},
			"validates 8 characters":                  {"contrase", ""},
			"validates 5 characters":                  {"contñ", ""},
			"validates only numbers":                  {"12345", ""},
			"does not validate less than 2":           {"1", Lang.T(D.MinSize, 5, D.Chars)},
			"no data":                                 {"", Lang.T(D.Field, D.Empty, D.NotAllowed)},
		}
	)
	for prueba, data := range dataPassword {
		t.Run((prueba + ": " + data.inputData), func(t *testing.T) {
			err := modelPassword.Validate(data.inputData)

			var err_str string
			if err != nil {
				err_str = err.Error()
			}

			if err_str != data.expected {
				t.Fatalf("result: [%v] expected: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}

// Radio tests
func Test_RadioButton(t *testing.T) {
	var (
		modelRadio = Radio("name=genre", "options=D:Dama,V:Varón")

		TestData = map[string]struct {
			inputData string
			expected  string
		}{
			"D Dato correcto":                {"D", ""},
			"V Dato correcto":                {"V", ""},
			"d Dato en minúscula incorrecto": {"d", "valor d no permitido en radio genre"},
			"v Dato en minúscula incorrecto": {"v", "valor v no permitido en radio genre"},
			"data ok?":                       {"0", "valor 0 no permitido en radio genre"},
		}
	)
	for prueba, data := range TestData {
		t.Run((prueba), func(t *testing.T) {
			err := modelRadio.Validate(data.inputData)

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

// Rut tests
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

// Select tests
func Test_Select(t *testing.T) {
	var (
		modelSelect = Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")

		dataSelect = map[string]struct {
			inputData string
			expected  string
		}{
			"una credencial ok?":  {"1", ""},
			"otro numero ok?":     {"3", ""},
			"0 existe?":           {"0", "valor 0 no permitido en select credentials"},
			"-1 valido?":          {"-1", "valor -1 no permitido en select credentials"},
			"carácter permitido?": {"%", "valor % no permitido en select credentials"},
			"con data?":           {"", "selección requerida campo credentials"},
			"sin espacios?":       {"luis ", "valor luis  no permitido en select credentials"},
		}
	)
	for prueba, data := range dataSelect {
		t.Run((prueba + " " + data.inputData), func(t *testing.T) {
			err := modelSelect.Validate(data.inputData)

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

// Text tests
func Test_InputText(t *testing.T) {
	var (
		modelText = Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)

		dataText = map[string]struct {
			inputData string
			expected  string
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

// TextArea tests
func Test_InputTextArea(t *testing.T) {
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
			"sin data permitido?":               {"", Lang.T(D.Field, D.Empty, D.NotAllowed)},
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

// TextSearch tests
func Test_InputTextSearch(t *testing.T) {
	var (
		modelTextSearch = TextSearch()

		dataTextSearch = map[string]struct {
			inputData string
			expected  string
		}{
			"palabra solo texto 15 caracteres?": {"Maria Jose Diaz", ""},
			"texto con ñ ok?":                   {"Ñuñez perez", ""},
			"tilde permitido?":                  {"peréz del rozal", "é tilde no permitida"},
			"mas de 20 caracteres permitidos?":  {"hola son mas de 21 ca", "tamaño máximo 20 caracteres"},
			"guion permitido":                   {"12038-0", ""},
			"fecha correcta?":                   {"1990-07-21", ""},
			"fecha incorrecta permitida?":       {"190-07-21", ""},
		}
	)
	for prueba, data := range dataTextSearch {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelTextSearch.Validate(data.inputData)

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

// Validation tests
func Test_Valid(t *testing.T) {
	var (
		validTestData = map[string]struct {
			text     string
			expected string
			permitted
		}{
			"números sin espacio ok":                    {"5648", "", permitted{Numbers: true}},
			"números con espacio ok":                    {"5648 78212", "", permitted{Numbers: true, Characters: []rune{' '}}},
			"error no permitido números con espacio":    {"5648 78212", "espacios en blanco no permitidos", permitted{Numbers: true}},
			"solo texto sin espacio ok":                 {"Maria", "", permitted{Letters: true}},
			"texto con espacios ok":                     {"Maria De Lourdes", "", permitted{Letters: true, Characters: []rune{' '}}},
			"texto con tildes y espacios ok":            {"María Dé Lourdes", "", permitted{Tilde: true, Letters: true, Characters: []rune{' '}}},
			"texto con numero sin espacios ok":          {"equipo01", "", permitted{Letters: true, Numbers: true}},
			"numero al inicio y texto sin espacios ok":  {"9equipo01", "", permitted{Letters: true, Numbers: true}},
			"numero al inicio y texto con espacios ok":  {"9equipo01 2equipo2", "", permitted{Letters: true, Numbers: true, Characters: []rune{' '}}},
			"error solo números no letras si espacios ": {"9equipo01 2equipo2", "carácter e no permitido", permitted{Numbers: true, Characters: []rune{' '}}},
			"correo con punto y @ ok":                   {"mi.correo1@mail.com", "", permitted{Characters: []rune{'@', '.'}, Numbers: true, Letters: true}},
			"error correo con tilde no permitido":       {"mí.correo@mail.com", "carácter í con tilde no permitida", permitted{Characters: []rune{'@', '.'}, Numbers: true, Letters: true}},
		}
	)

	for prueba, data := range validTestData {
		t.Run((prueba + " " + data.text), func(t *testing.T) {
			tempInput := input{permitted: data.permitted}
			err := tempInput.Validate(data.text)

			var err_str string
			if err != nil {
				err_str = err.Error()
			}

			if err_str != data.expected {
				t.Fatalf("\nexpected:\n[%v]\n\nresult:\n[%v]\n%v\n", data.expected, err, data.text)
			}
		})
	}
}
