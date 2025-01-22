package inputs

import (
	"fmt"
	"strconv"
	"testing"
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
		{"DataList - Credencial válida", "1", "", DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"DataList - Otro número válido", "3", "", DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"DataList - Valor 0", "0", "valor 0 no permitido en datalist credentials", DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"DataList - Valor negativo", "-1", "valor -1 no permitido en datalist credentials", DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"DataList - Carácter especial", "%", "valor % no permitido en datalist credentials", DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"DataList - Sin datos", "", "selección requerida campo credentials", DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"DataList - Espacios", "luis ", "valor luis  no permitido en datalist credentials", DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},

		// Date tests
		{"Date - Formato correcto", "2002-12-03", "", Date()},
		{"Date - Año bisiesto", "2020-02-29", "", Date()},
		{"Date - Año no bisiesto", "2023-02-29", Lang.T(D.February, D.DoesNotHave, "29", D.Days, D.Year, "2023"), Date()},
		{"Date - Junio 31 días", "2023-06-31", Lang.T(D.June, D.DoesNotHave, "31", D.Days), Date()},
		{"Date - Carácter extra", "2002-12-03-", Lang.T(D.InvalidDateFormat, "2006-01-02"), Date()},
		{"Date - Formato incorrecto", "21/12/1998", Lang.T(D.InvalidDateFormat, "2006-01-02"), Date()},
		{"Date - Mes 31", "2020-31-01", Lang.T(D.Month, 31, D.NotValid), Date()},
		{"Date - Sin año", "31-01", Lang.T(D.InvalidDateFormat, "2006-01-02"), Date()},
		{"Date - Datos incorrectos", "0000-00-00", Lang.T(D.InvalidDateFormat), Date()},
		{"Date - Sin datos", "", Lang.T(D.Field, D.Empty, D.NotAllowed), Date()},

		// FilePath tests
		{"FilePath - Ruta correcta", ".\\files\\1234\\", "", FilePath()},
		{"FilePath - Ruta sin punto inicial", "\\files\\1234\\", Lang.Err(D.DoNotStartWith, '\\').Error(), FilePath()},
		{"FilePath - Ruta absoluta Linux", "/home/user/files/", "", FilePath()},
		{"FilePath - Ruta absoluta sin punto inicial", "/files/1234/", "", FilePath()},
		{"FilePath - Ruta relativa sin directorios", ".\\", "", FilePath()},
		{"FilePath - Ruta relativa con slash final", "./files/1234/", "", FilePath()},
		{"FilePath - Ruta con nombre de archivo", ".\\files\\1234\\archivo.txt", "", FilePath()},
		{"FilePath - Ruta con guiones bajos", ".\\mi_directorio\\sub_directorio\\", "", FilePath()},
		{"FilePath - Número como ruta", "5", "", FilePath()},
		{"FilePath - Palabra como ruta", "path", "", FilePath()},
		{"FilePath - Espacios en blanco", ".\\path with white space\\", Lang.T(D.WhiteSpace, D.NotAllowed), FilePath()},

		// IP tests
		{"IP - IPv4 válida", "192.168.1.1", "", Ip()},
		{"IP - IPv6 válida", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", "", Ip()},
		{"IP - Formato incorrecto", "192.168.1.1.8", Lang.T(D.Format, "IPv4", D.NotValid), Ip()},
		{"IP - Dirección 0.0.0.0", "0.0.0.0", Lang.T(D.Example, "IP", D.NotAllowed, ':', "0.0.0.0"), Ip()},
		{"IP - Sin datos", "", Lang.T(D.Field, D.Empty, D.NotAllowed), Ip()},

		// Mail tests
		{"Mail - Correo normal", "mi.correo@mail.com", "", Mail()},
		{"Mail - Correo simple", "correo@mail.com", "", Mail()},

		// Number tests
		{"Number - Número correcto", "100", "", Number()},
		{"Number - Dígito 0", "0", "", Number()},
		{"Number - Dígito 1", "1", "", Number()},
		{"Number - uint64 20 dígitos", "18446744073709551615", "", Number()},
		{"Number - uint64 21 dígitos", "184467440737095516151", Lang.T(D.MaxSize, 20, D.Chars), Number()},
		{"Number - int64 19 dígitos", "9223372036854775807", "", Number()},
		{"Number - int32 10 dígitos", "2147483647", "", Number()},
		{"Number - 18 dígitos", "100002323262637278", "", Number()},
		{"Number - Letra en número", "100002323E262637278", Lang.T('E', D.NotNumber), Number()},
		{"Number - Número negativo", "-100", Lang.T('-', D.NotNumber), Number()},
		{"Number - Texto en lugar de número", "lOO", Lang.T('l', D.NotNumber), Number()},
		{"Number - Texto y número", "l500", Lang.T('l', D.NotNumber), Number()},

		// Password tests
		{"Password - Números, letras y caracteres", "c0ntra3", "", Password()},
		{"Password - Muchos caracteres", "M1 contraseÑ4", "", Password()},
		{"Password - 8 caracteres", "contrase", "", Password()},
		{"Password - 5 caracteres", "contñ", "", Password()},
		{"Password - Solo números", "12345", "", Password()},
		{"Password - Menos de 2 caracteres", "1", Lang.T(D.MinSize, 5, D.Chars), Password()},
		{"Password - Sin datos", "", Lang.T(D.Field, D.Empty, D.NotAllowed), Password()},

		// Radio tests
		{"Radio - Dato correcto D", "D", "", Radio("name=genre", "options=D:Dama,V:Varón")},
		{"Radio - Dato correcto V", "V", "", Radio("name=genre", "options=D:Dama,V:Varón")},
		{"Radio - Dato en minúscula incorrecto d", "d", "valor d no permitido en radio genre", Radio("name=genre", "options=D:Dama,V:Varón")},
		{"Radio - Dato en minúscula incorrecto v", "v", "valor v no permitido en radio genre", Radio("name=genre", "options=D:Dama,V:Varón")},
		{"Radio - Dato 0 no permitido", "0", "valor 0 no permitido en radio genre", Radio("name=genre", "options=D:Dama,V:Varón")},

		// Rut tests
		{"Rut - Sin guión", "15890022k", Lang.T(D.HyphenMissing), Rut()},
		{"Rut - Sin guión", "177344788", Lang.T(D.HyphenMissing), Rut()},
		{"Rut - Correcto", "7863697-1", "", Rut()},
		{"Rut - K mayúscula", "20373221-K", "", Rut()},
		{"Rut - k minúscula", "20373221-k", "", Rut()},
		{"Rut - Dígito W inválido", "7863697-W", Lang.T(D.Digit, D.Verifier, "w", D.NotValid), Rut()},
		{"Rut - Cambio dígito a k", "7863697-k", Lang.T(D.Digit, D.Verifier, "k", D.NotValid), Rut()},
		{"Rut - Cambio dígito a 0", "7863697-0", Lang.T(D.Digit, D.Verifier, 0, D.NotValid), Rut()},
		{"Rut - Correcto", "14080717-6", "", Rut()},
		{"Rut - Dígito 0 inválido", "14080717-0", Lang.T(D.Digit, D.Verifier, 0, D.NotValid), Rut()},
		{"Rut - Cero al inicio", "07863697-1", Lang.T(D.DoNotStartWith, D.Digit, 0), Rut()},
		{"Rut - Solo espacio", " ", Lang.T(D.DoNotStartWith, D.WhiteSpace), Rut()},

		// Select tests
		{"Select - Credencial válida", "1", "", Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"Select - Otro número válido", "3", "", Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"Select - 0 no permitido", "0", "valor 0 no permitido en select credentials", Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"Select - -1 no permitido", "-1", "valor -1 no permitido en select credentials", Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"Select - Carácter % no permitido", "%", "valor % no permitido en select credentials", Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"Select - Sin datos", "", "selección requerida campo credentials", Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"Select - Espacios no permitidos", "luis ", "valor luis  no permitido en select credentials", Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},

		// Text tests
		{"Text - Nombre con punto válido", "Dr. Maria Jose Diaz Cadiz", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Tilde no permitida", "peréz del rozal", "é tilde no permitida", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Texto con ñ permitido", "Ñuñez perez", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Texto de 4 caracteres", "hola", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Texto de 3 caracteres", "los", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Oración válida", "hola que tal", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Solo números permitidos", "100", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Caracteres y comas permitidos", "los,true, vengadores", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Campo vacío permitido", "", "", Text(`!required`)},
		{"Text - Campo vacío no permitido", "8", "tamaño mínimo 2 caracteres", Text(`required`)},
		{"Text - Palabra y número permitidos", "son 4 bidones", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Paréntesis y número permitidos", "son 4 (4 bidones)", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Solo paréntesis permitidos", "son (bidones)", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Palabras y número permitidos", "apellido Actualizado 1", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Slash no permitido", " estos son \\n los podria", "carácter \\ no permitido", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Nombres de archivos separados por coma", "dino.png, gatito.jpeg", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},

		// TextArea tests
		{"TextArea - Todos caracteres permitidos", "hola: esto, es. la - prueba 10", "", TextArea()},
		{"TextArea - Salto de línea y guion permitidos", "hola:\n esto, es. la - \nprueba 10", "", TextArea()},
		{"TextArea - Letra ñ, paréntesis y $ permitidos", "soy ñato o Ñato (aqui) costo $10000.", "", TextArea()},
		{"TextArea - Solo texto y espacio", "hola esto es una prueba", "", TextArea()},
		{"TextArea - Texto y puntuación", "hola: esto es una prueba", "", TextArea()},
		{"TextArea - Texto, puntuación y coma", "hola: esto,true, es una prueba", "", TextArea()},
		{"TextArea - 5 caracteres", ", .s5", "", TextArea()},
		{"TextArea - Sin datos no permitido", "", Lang.T(D.Field, D.Empty, D.NotAllowed), TextArea()},
		{"TextArea - Carácter # permitido", "# son", "", TextArea()},
		{"TextArea - Carácter ¿ no permitido", " ¿ si ?", "carácter ¿ no permitido", TextArea()},
		{"TextArea - Tildes permitidas", " mí tílde", "", TextArea()},
		{"TextArea - 1 carácter no permitido", "1", "tamaño mínimo " + strconv.Itoa(TextArea().Minimum) + " caracteres", TextArea()},
		{"TextArea - Nombre correcto", "Dr. Pato Gomez", "", TextArea()},
		{"TextArea - Solo espacio no permitido", " ", "tamaño mínimo " + strconv.Itoa(TextArea().Minimum) + " caracteres", TextArea()},
		{"TextArea - Texto largo válido", `IRRITACION EN PIEL DE ROSTRO. ALERGIAS NO. CIRUGIAS NO. ACTUAL TTO CON ISOTRETINOINA 10MG - ENERO 2022. EN TTO ACTUAL CON VIT D. EXAMEN DE LAB 20-12-2022. SIN OTROS ANTECEDENTES`, "", TextArea()},
		{"TextArea - Texto con saltos de línea", `HOY......Referido por        : dr. ........
			Motivo                    : ........
			Premedicacion  : ........`, "", TextArea()},

		// TextSearch tests
		{"TextSearch - Palabra de 15 caracteres", "Maria Jose Diaz", "", TextSearch()},
		{"TextSearch - Texto con ñ permitido", "Ñuñez perez", "", TextSearch()},
		{"TextSearch - Tilde no permitida", "peréz del rozal", "é tilde no permitida", TextSearch()},
		{"TextSearch - Más de 20 caracteres no permitido", "hola son mas de 21 ca", "tamaño máximo 20 caracteres", TextSearch()},
		{"TextSearch - Guion permitido", "12038-0", "", TextSearch()},
		{"TextSearch - Fecha correcta", "1990-07-21", "", TextSearch()},
		{"TextSearch - Fecha incorrecta permitida", "190-07-21", "", TextSearch()},

		// Validation tests
		{"Validation - Números sin espacio", "5648", "", input{permitted: permitted{Numbers: true}}},
		{"Validation - Números con espacio", "5648 78212", "", input{permitted: permitted{Numbers: true, Characters: []rune{' '}}}},
		{"Validation - Error: espacio no permitido", "5648 78212", "espacio en blanco no permitido", input{permitted: permitted{Numbers: true}}},
		{"Validation - Solo texto sin espacio", "Maria", "", input{permitted: permitted{Letters: true}}},
		{"Validation - Texto con espacios", "Maria De Lourdes", "", input{permitted: permitted{Letters: true, Characters: []rune{' '}}}},
		{"Validation - Texto con tildes y espacios", "María Dé Lourdes", "", input{permitted: permitted{Tilde: true, Letters: true, Characters: []rune{' '}}}},
		{"Validation - Texto con número sin espacios", "equipo01", "", input{permitted: permitted{Letters: true, Numbers: true}}},
		{"Validation - Número al inicio y texto sin espacios", "9equipo01", "", input{permitted: permitted{Letters: true, Numbers: true}}},
		{"Validation - Número al inicio y texto con espacios", "9equipo01 2equipo2", "", input{permitted: permitted{Letters: true, Numbers: true, Characters: []rune{' '}}}},
		{"Validation - Error: solo números permitidos", "9equipo01 2equipo2", "carácter e no permitido", input{permitted: permitted{Numbers: true, Characters: []rune{' '}}}},
		{"Validation - Correo con punto y @", "mi.correo1@mail.com", "", input{permitted: permitted{Characters: []rune{'@', '.'}, Numbers: true, Letters: true}}},
		{"Validation - Error: tilde no permitida en correo", "mí.correo@mail.com", "í tilde no permitida", input{permitted: permitted{Characters: []rune{'@', '.'}, Numbers: true, Letters: true}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.model.Validate(test.inputData)

			var err_str string
			if err != nil {
				err_str = err.Error()
			}

			if err_str != test.expected {
				fmt.Println(test.name)
				t.Fatalf("result: [%v] expected: [%v]\n%v", err, test.expected, test.inputData)
			}
		})
	}
}
