package mono_test

import (
	"strconv"
	"testing"

	. "github.com/cdvelop/mono"
)

type testStruct struct {
	name      string // Test case description including input details and expected behavior
	inputData string
	expected  string
	model     interface{ Validate(string) error }
}

func Test_AllInputs(t *testing.T) {

	tests := []testStruct{
		// DataList tests
		{"DataList - Credencial válida", "1", "", IN.DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"DataList - Otro número válido", "3", "", IN.DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"DataList - Valor 0", "0", "valor 0 no permitido en credentials", IN.DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"DataList - Valor negativo", "-1", "valor -1 no permitido en credentials", IN.DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"DataList - Carácter especial", "%", "valor % no permitido en credentials", IN.DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"DataList - Sin datos", "", "selección requerida campo credentials", IN.DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"DataList - Espacios", "luis ", "valor luis  no permitido en credentials", IN.DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},

		// Date tests
		{"Date - Formato correcto", "2002-12-03", "", IN.Date()},
		{"Date - Año bisiesto", "2020-02-29", "", IN.Date()},
		{"Date - Año no bisiesto", "2023-02-29", R.T(D.February, D.DoesNotHave, "29", D.Days, D.Year, "2023"), IN.Date()},
		{"Date - Junio 31 días", "2023-06-31", R.T(D.June, D.DoesNotHave, "31", D.Days), IN.Date()},
		{"Date - Carácter extra", "2002-12-03-", R.T(D.InvalidDateFormat, "2006-01-02"), IN.Date()},
		{"Date - Formato incorrecto", "21/12/1998", R.T(D.InvalidDateFormat, "2006-01-02"), IN.Date()},
		{"Date - Mes 31", "2020-31-01", R.T(D.Month, 31, D.NotValid), IN.Date()},
		{"Date - Sin año", "31-01", R.T(D.InvalidDateFormat, "2006-01-02"), IN.Date()},
		{"Date - Datos incorrectos", "0000-00-00", R.T(D.InvalidDateFormat), IN.Date()},
		{"Date - Sin datos", "", R.T(D.Field, D.Empty, D.NotAllowed), IN.Date()},

		// FilePath tests
		{"FilePath - Ruta correcta", ".\\files\\1234\\", "", IN.FilePath()},
		{"FilePath - Ruta sin punto inicial", "\\files\\1234\\", R.Err(D.DoNotStartWith, '\\').Error(), IN.FilePath()},
		{"FilePath - Ruta absoluta Linux", "/home/user/files/", "", IN.FilePath()},
		{"FilePath - Ruta absoluta sin punto inicial", "/files/1234/", "", IN.FilePath()},
		{"FilePath - Ruta relativa sin directorios", ".\\", "", IN.FilePath()},
		{"FilePath - Ruta relativa con slash final", "./files/1234/", "", IN.FilePath()},
		{"FilePath - Ruta con nombre de archivo", ".\\files\\1234\\archivo.txt", "", IN.FilePath()},
		{"FilePath - Ruta con guiones bajos", ".\\mi_directorio\\sub_directorio\\", "", IN.FilePath()},
		{"FilePath - Número como ruta", "5", "", IN.FilePath()},
		{"FilePath - Palabra como ruta", "path", "", IN.FilePath()},
		{"FilePath - Espacios en blanco", ".\\path with white space\\", R.T(D.WhiteSpace, D.NotAllowed), IN.FilePath()},

		// IP tests
		{"IP - IPv4 válida", "192.168.1.1", "", IN.Ip()},
		{"IP - IPv6 válida", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", "", IN.Ip()},
		{"IP - Formato incorrecto", "192.168.1.1.8", R.T(D.Format, "IPv4", D.NotValid), IN.Ip()},
		{"IP - Dirección 0.0.0.0", "0.0.0.0", R.T(D.Example, "IP", D.NotAllowed, ':', "0.0.0.0"), IN.Ip()},
		{"IP - Sin datos", "", R.T(D.Field, D.Empty, D.NotAllowed), IN.Ip()},

		// Mail tests
		{"Mail - Correo normal", "mi.correo@mail.com", "", IN.Mail()},
		{"Mail - Correo simple", "correo@mail.com", "", IN.Mail()},

		// Number tests
		{"Number - Número correcto", "100", "", IN.Number()},
		{"Number - Dígito 0", "0", "", IN.Number()},
		{"Number - Dígito 1", "1", "", IN.Number()},
		{"Number - uint64 20 dígitos", "18446744073709551615", "", IN.Number()},
		{"Number - uint64 21 dígitos", "184467440737095516151", R.T(D.MaxSize, 20, D.Chars), IN.Number()},
		{"Number - int64 19 dígitos", "9223372036854775807", "", IN.Number()},
		{"Number - int32 10 dígitos", "2147483647", "", IN.Number()},
		{"Number - 18 dígitos", "100002323262637278", "", IN.Number()},
		{"Number - Letra en número", "100002323E262637278", R.T('E', D.NotNumber), IN.Number()},
		{"Number - Número negativo", "-100", R.T('-', D.NotNumber), IN.Number()},
		{"Number - Texto en lugar de número", "lOO", R.T('l', D.NotNumber), IN.Number()},
		{"Number - Texto y número", "l500", R.T('l', D.NotNumber), IN.Number()},

		// Password tests
		{"Password - Números, letras y caracteres", "c0ntra3", "", IN.Password()},
		{"Password - Muchos caracteres", "M1 contraseÑ4", "", IN.Password()},
		{"Password - 8 caracteres", "contrase", "", IN.Password()},
		{"Password - 5 caracteres", "contñ", "", IN.Password()},
		{"Password - Solo números", "12345", "", IN.Password()},
		{"Password - Menos de 2 caracteres", "1", R.T(D.MinSize, 5, D.Chars), IN.Password()},
		{"Password - Sin datos", "", R.T(D.Field, D.Empty, D.NotAllowed), IN.Password()},

		// Radio tests
		{"Radio - Dato correcto D", "D", "", IN.Radio("name=genre", "options=D:Dama,V:Varón")},
		{"Radio - Dato correcto V", "V", "", IN.Radio("name=genre", "options=D:Dama,V:Varón")},
		{"Radio - Dato en minúscula incorrecto d", "d", "valor d no permitido en genre", IN.Radio("name=genre", "options=D:Dama,V:Varón")},
		{"Radio - Dato en minúscula incorrecto v", "v", "valor v no permitido en genre", IN.Radio("name=genre", "options=D:Dama,V:Varón")},
		{"Radio - Dato 0 no permitido", "0", "valor 0 no permitido en genre", IN.Radio("name=genre", "options=D:Dama,V:Varón")},

		// Rut tests
		{"Rut - Sin guión", "15890022k", R.T(D.HyphenMissing), IN.Rut()},
		{"Rut - Sin guión", "177344788", R.T(D.HyphenMissing), IN.Rut()},
		{"Rut - Correcto", "7863697-1", "", IN.Rut()},
		{"Rut - K mayúscula", "20373221-K", "", IN.Rut()},
		{"Rut - k minúscula", "20373221-k", "", IN.Rut()},
		{"Rut - Dígito W inválido", "7863697-W", R.T(D.Digit, D.Verifier, "w", D.NotValid), IN.Rut()},
		{"Rut - Cambio dígito a k", "7863697-k", R.T(D.Digit, D.Verifier, "k", D.NotValid), IN.Rut()},
		{"Rut - Cambio dígito a 0", "7863697-0", R.T(D.Digit, D.Verifier, 0, D.NotValid), IN.Rut()},
		{"Rut - Correcto", "14080717-6", "", IN.Rut()},
		{"Rut - Dígito 0 inválido", "14080717-0", R.T(D.Digit, D.Verifier, 0, D.NotValid), IN.Rut()},
		{"Rut - Cero al inicio", "07863697-1", R.T(D.DoNotStartWith, D.Digit, 0), IN.Rut()},
		{"Rut - Solo espacio", " ", R.T(D.DoNotStartWith, D.WhiteSpace), IN.Rut()},

		// Select tests
		{"Select - Credencial válida", "1", "", IN.Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"Select - Otro número válido", "3", "", IN.Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"Select - 0 no permitido", "0", "valor 0 no permitido en credentials", IN.Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"Select - -1 no permitido", "-1", "valor -1 no permitido en credentials", IN.Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"Select - Carácter % no permitido", "%", "valor % no permitido en credentials", IN.Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"Select - Sin datos", "", "selección requerida campo credentials", IN.Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"Select - Espacios no permitidos", "luis ", "valor luis  no permitido en credentials", IN.Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},

		// Text tests
		{"Text - Nombre con punto válido", "Dr. Maria Jose Diaz Cadiz", "", IN.Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Tilde no permitida", "peréz del rozal", "é tilde no permitida", IN.Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Texto con ñ permitido", "Ñuñez perez", "", IN.Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Texto de 4 caracteres", "hola", "", IN.Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Texto de 3 caracteres", "los", "", IN.Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Oración válida", "hola que tal", "", IN.Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Solo números permitidos", "100", "", IN.Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Caracteres y comas permitidos", "los,true, vengadores", "", IN.Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Campo vacío permitido", "", "", IN.Text(`!required`)},
		{"Text - Campo vacío no permitido", "8", "tamaño mínimo 2 caracteres", IN.Text(`required`)},
		{"Text - Palabra y número permitidos", "son 4 bidones", "", IN.Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Paréntesis y número permitidos", "son 4 (4 bidones)", "", IN.Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Solo paréntesis permitidos", "son (bidones)", "", IN.Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Palabras y número permitidos", "apellido Actualizado 1", "", IN.Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Slash no permitido", " estos son \\n los podria", "carácter \\ no permitido", IN.Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Nombres de archivos separados por coma", "dino.png, gatito.jpeg", "", IN.Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},

		// TextArea tests
		{"TextArea - Todos caracteres permitidos", "hola: esto, es. la - prueba 10", "", IN.TextArea()},
		{"TextArea - Salto de línea y guion permitidos", "hola:\n esto, es. la - \nprueba 10", "", IN.TextArea()},
		{"TextArea - Letra ñ, paréntesis y $ permitidos", "soy ñato o Ñato (aqui) costo $10000.", "", IN.TextArea()},
		{"TextArea - Solo texto y espacio", "hola esto es una prueba", "", IN.TextArea()},
		{"TextArea - Texto y puntuación", "hola: esto es una prueba", "", IN.TextArea()},
		{"TextArea - Texto, puntuación y coma", "hola: esto,true, es una prueba", "", IN.TextArea()},
		{"TextArea - 5 caracteres", ", .s5", "", IN.TextArea()},
		{"TextArea - Sin datos no permitido", "", R.T(D.Field, D.Empty, D.NotAllowed), IN.TextArea()},
		{"TextArea - Carácter # permitido", "# son", "", IN.TextArea()},
		{"TextArea - Carácter ¿ no permitido", " ¿ si ?", "carácter ¿ no permitido", IN.TextArea()},
		{"TextArea - Tildes permitidas", " mí tílde", "", IN.TextArea()},
		{"TextArea - 1 carácter no permitido", "1", "tamaño mínimo " + strconv.Itoa(IN.TextArea().Minimum) + " caracteres", IN.TextArea()},
		{"TextArea - Nombre correcto", "Dr. Pato Gomez", "", IN.TextArea()},
		{"TextArea - Solo espacio no permitido", " ", "tamaño mínimo " + strconv.Itoa(IN.TextArea().Minimum) + " caracteres", IN.TextArea()},
		{"TextArea - Texto largo válido", `IRRITACION EN PIEL DE ROSTRO. ALERGIAS NO. CIRUGIAS NO. ACTUAL TTO CON ISOTRETINOINA 10MG - ENERO 2022. EN TTO ACTUAL CON VIT D. EXAMEN DE LAB 20-12-2022. SIN OTROS ANTECEDENTES`, "", IN.TextArea()},
		{"TextArea - Texto con saltos de línea", `HOY......Referido por        : dr. ........
			Motivo                    : ........
			Premedicacion  : ........`, "", IN.TextArea()},

		// TextSearch tests
		{"TextSearch - Palabra de 15 caracteres", "Maria Jose Diaz", "", IN.TextSearch()},
		{"TextSearch - Texto con ñ permitido", "Ñuñez perez", "", IN.TextSearch()},
		{"TextSearch - Tilde no permitida", "peréz del rozal", "é tilde no permitida", IN.TextSearch()},
		{"TextSearch - Más de 20 caracteres no permitido", "hola son mas de 21 ca", "tamaño máximo 20 caracteres", IN.TextSearch()},
		{"TextSearch - Guion permitido", "12038-0", "", IN.TextSearch()},
		{"TextSearch - Fecha correcta", "1990-07-21", "", IN.TextSearch()},
		{"TextSearch - Fecha incorrecta permitida", "190-07-21", "", IN.TextSearch()},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.model.Validate(test.inputData)

			var err_str string
			if err != nil {
				err_str = err.Error()
			}

			if err_str != test.expected {

				t.Fatalf("\n❌ %v:\n- result:\n[%v]\n- expected:\n[%v]\n- input:\n[%v]\n", test.name, err, test.expected, test.inputData)
			}
		})
	}
}
