package inputs

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

type testStruct struct {
	name        string
	description string
	inputData   string
	expected    string
	model       interface{ Validate(string) error }
}

func Test_AllInputs(t *testing.T) {
	tests := []testStruct{
		// DataList tests
		{"DataList - Credencial válida", "una credencial ok?", "1", "", DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"DataList - Otro número válido", "otro numero ok?", "3", "", DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"DataList - Valor 0", "0 existe?", "0", "valor 0 no permitido en datalist credentials", DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"DataList - Valor negativo", "-1 valido?", "-1", "valor -1 no permitido en datalist credentials", DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"DataList - Carácter especial", "carácter permitido?", "%", "valor % no permitido en datalist credentials", DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"DataList - Sin datos", "con data?", "", "selección requerida campo credentials", DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"DataList - Espacios", "sin espacios?", "luis ", "valor luis  no permitido en datalist credentials", DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},

		// Date tests
		{"Date - Formato correcto", "correct", "2002-12-03", "", Date()},
		{"Date - Año bisiesto", "February 29th leap year", "2020-02-29", "", Date()},
		{"Date - Año no bisiesto", "February 29th non leap year", "2023-02-29", Lang.T(D.February, D.DoesNotHave, "29", D.Days, D.Year, "2023"), Date()},
		{"Date - Junio 31 días", "June does not have 31 days", "2023-06-31", Lang.T(D.June, D.DoesNotHave, "31", D.Days), Date()},
		{"Date - Carácter extra", "incorrect extra character", "2002-12-03-", Lang.T(D.InvalidDateFormat, "2006-01-02"), Date()},
		{"Date - Formato incorrecto", "incorrect format", "21/12/1998", Lang.T(D.InvalidDateFormat, "2006-01-02"), Date()},
		{"Date - Mes 31", "incorrect month 31", "2020-31-01", Lang.T(D.Month, 31, D.NotValid), Date()},
		{"Date - Sin año", "shortened date without year ok?", "31-01", Lang.T(D.InvalidDateFormat, "2006-01-02"), Date()},
		{"Date - Datos incorrectos", "incorrect data", "0000-00-00", Lang.T(D.InvalidDateFormat), Date()},
		{"Date - Sin datos", "all data correct?", "", Lang.T(D.InvalidDateFormat, "2006-01-02"), Date()},

		// FilePath tests
		{"FilePath - Ruta correcta", "correct path", ".\\files\\1234\\", "", FilePath()},
		{"FilePath - Ruta sin punto inicial", "correct path without starting point?", "\\files\\1234\\", Lang.Err(D.DoNotStartWith, '\\').Error(), FilePath()},
		{"FilePath - Ruta absoluta Linux", "absolute path in Linux", "/home/user/files/", "", FilePath()},
		{"FilePath - Ruta absoluta sin punto inicial", "absolute path without starting point", "/files/1234/", "", FilePath()},
		{"FilePath - Ruta relativa sin directorios", "relative path without directories ok?", ".\\", "", FilePath()},
		{"FilePath - Ruta relativa con slash final", "relative path with final slash", "./files/1234/", "", FilePath()},
		{"FilePath - Ruta con nombre de archivo", "path with filename", ".\\files\\1234\\archivo.txt", "", FilePath()},
		{"FilePath - Ruta con guiones bajos", "path with directory names using underscores", ".\\mi_directorio\\sub_directorio\\", "", FilePath()},
		{"FilePath - Número como ruta", "is a number a valid path?", "5", "", FilePath()},
		{"FilePath - Palabra como ruta", "is a single word a valid path?", "path", "", FilePath()},
		{"FilePath - Espacios en blanco", "white spaces allowed?", ".\\path with white space\\", Lang.T(D.WhiteSpace, D.NotAllowed), FilePath()},

		// IP tests
		{"IP - IPv4 válida", "IPv4 ok", "192.168.1.1", "", Ip()},
		{"IP - IPv6 válida", "IPv6 ok", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", "", Ip()},
		{"IP - Formato incorrecto", "ip incorrecta", "192.168.1.1.8", Lang.T(D.Format, "IPv4", D.NotValid), Ip()},
		{"IP - Dirección 0.0.0.0", "correcto?", "0.0.0.0", Lang.T(D.Example, "IP", D.NotAllowed, ':', "0.0.0.0"), Ip()},
		{"IP - Sin datos", "sin data", "", Lang.T(D.Field, D.Empty, D.NotAllowed), Ip()},

		// Mail tests
		{"Mail - Correo normal", "correo normal", "mi.correo@mail.com", "", Mail()},
		{"Mail - Correo simple", "correo un campo", "correo@mail.com", "", Mail()},

		// Number tests
		{"Number - Número correcto", "correct number 100", "100", "", Number()},
		{"Number - Dígito 0", "single digit 0", "0", "", Number()},
		{"Number - Dígito 1", "single digit 1", "1", "", Number()},
		{"Number - uint64 20 dígitos", "uint64 20 chars", "18446744073709551615", "", Number()},
		{"Number - uint64 21 dígitos", "uint64 21 chars", "184467440737095516150", Lang.T(D.MaxSize, 20, D.Chars), Number()},
		{"Number - int64 19 dígitos", "int64 19 chars", "9223372036854775807", "", Number()},
		{"Number - int32 10 dígitos", "int32 10 chars", "2147483647", "", Number()},
		{"Number - 18 dígitos", "18 digits", "100002323262637278", "", Number()},
		{"Number - Letra en número", "large number with letter", "10000232E26263727", Lang.T('E', D.NotNumber), Number()},
		{"Number - Número negativo", "negative number -100", "-100", Lang.T('-', D.NotNumber), Number()},
		{"Number - Texto en lugar de número", "text instead of number", "h", Lang.T('h', D.NotNumber), Number()},
		{"Number - Texto y número", "text and number", "h1", Lang.T('h', D.NotNumber), Number()},

		// Password tests
		{"Password - Números, letras y caracteres", "validates numbers letters and character", "c0ntra3", "", Password()},
		{"Password - Muchos caracteres", "validates many characters", "M1 contraseÑ4", "", Password()},
		{"Password - 8 caracteres", "validates 8 characters", "contrase", "", Password()},
		{"Password - 5 caracteres", "validates 5 characters", "contñ", "", Password()},
		{"Password - Solo números", "validates only numbers", "12345", "", Password()},
		{"Password - Menos de 2 caracteres", "does not validate less than 2", "1", Lang.T(D.MinSize, 5, D.Chars), Password()},
		{"Password - Sin datos", "no data", "", Lang.T(D.Field, D.Empty, D.NotAllowed), Password()},

		// Radio tests
		{"Radio - Dato correcto D", "D Dato correcto", "D", "", Radio("name=genre", "options=D:Dama,V:Varón")},
		{"Radio - Dato correcto V", "V Dato correcto", "V", "", Radio("name=genre", "options=D:Dama,V:Varón")},
		{"Radio - Dato en minúscula incorrecto d", "d Dato en minúscula incorrecto", "d", "valor d no permitido en radio genre", Radio("name=genre", "options=D:Dama,V:Varón")},
		{"Radio - Dato en minúscula incorrecto v", "v Dato en minúscula incorrecto", "v", "valor v no permitido en radio genre", Radio("name=genre", "options=D:Dama,V:Varón")},
		{"Radio - Dato 0 no permitido", "data ok?", "0", "valor 0 no permitido en radio genre", Radio("name=genre", "options=D:Dama,V:Varón")},

		// Rut tests
		{"Rut - Sin guión", "without hyphen 15890022k", "15890022k", Lang.T(D.HyphenMissing), Rut()},
		{"Rut - Sin guión", "no hyphen 177344788", "177344788", Lang.T(D.HyphenMissing), Rut()},
		{"Rut - Correcto", "ok 7863697-1", "7863697-1", "", Rut()},
		{"Rut - K mayúscula", "ok uppercase K 20373221-K", "20373221-k", "", Rut()},
		{"Rut - k minúscula", "ok lowercase k 20373221-k", "20373221-k", "", Rut()},
		{"Rut - Dígito W inválido", "valid run? allowed?", "7863697-W", Lang.T(D.Digit, D.Verifier, "w", D.NotValid), Rut()},
		{"Rut - Cambio dígito a k", "change digit to k 7863697-k", "7863697-k", Lang.T(D.Digit, D.Verifier, "k", D.NotValid), Rut()},
		{"Rut - Cambio dígito a 0", "change digit to 0 7863697-0", "7863697-0", Lang.T(D.Digit, D.Verifier, 0, D.NotValid), Rut()},
		{"Rut - Correcto", "ok 14080717-6", "14080717-6", "", Rut()},
		{"Rut - Dígito 0 inválido", "incorrect 14080717-0", "14080717-0", Lang.T(D.Digit, D.Verifier, 0, D.NotValid), Rut()},
		{"Rut - Cero al inicio", "correct zero at start?", "07863697-1", Lang.T(D.DoNotStartWith, D.Digit, 0), Rut()},
		{"Rut - Solo espacio", "correct data only space?", " ", Lang.T(D.DoNotStartWith, D.WhiteSpace), Rut()},

		// Select tests
		{"Select - Credencial válida", "una credencial ok?", "1", "", Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"Select - Otro número válido", "otro numero ok?", "3", "", Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"Select - 0 no permitido", "0 existe?", "0", "valor 0 no permitido en select credentials", Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"Select - -1 no permitido", "-1 valido?", "-1", "valor -1 no permitido en select credentials", Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"Select - Carácter % no permitido", "carácter permitido?", "%", "valor % no permitido en select credentials", Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"Select - Sin datos", "con data?", "", "selección requerida campo credentials", Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"Select - Espacios no permitidos", "sin espacios?", "luis ", "valor luis  no permitido en select credentials", Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},

		// Text tests
		{"Text - Nombre con punto válido", "nombre correcto con punto?", "Dr. Maria Jose Diaz Cadiz", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Tilde no permitida", "no tilde", "peréz del rozal", "é tilde no permitida", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Texto con ñ permitido", "texto con ñ", "Ñuñez perez", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Texto de 4 caracteres", "texto correcto + 3 caracteres", "hola", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Texto de 3 caracteres", "texto correcto 3 caracteres", "los", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Oración válida", "oración ok", "hola que tal", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Solo números permitidos", "solo Dato numérico permitido?", "100", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Caracteres y comas permitidos", "con caracteres y coma", "los,true, vengadores", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Campo vacío permitido", "sin data ok se permite campo vació", "", "", Text(`!required`)},
		{"Text - Campo vacío no permitido", "no se permite campo vació", "8", "tamaño mínimo 2 caracteres", Text(`required`)},
		{"Text - Palabra y número permitidos", "palabra mas numero permitido", "son 4 bidones", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Paréntesis y número permitidos", "con paréntesis y numero", "son 4 (4 bidones)", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Solo paréntesis permitidos", "con solo paréntesis", "son (bidones)", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Palabras y número permitidos", "palabras y numero", "apellido Actualizado 1", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Slash no permitido", "palabra con slash?", " estos son \\n los podria", "carácter \\ no permitido", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"Text - Nombres de archivos separados por coma", "nombre de archivos separados por ,", "dino.png, gatito.jpeg", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},

		// TextArea tests
		{"TextArea - Todos caracteres permitidos", "todo los caracteres permitidos?", "hola: esto, es. la - prueba 10", "", TextArea()},
		{"TextArea - Salto de línea y guion permitidos", "salto de linea permitido? y guion", "hola:\n esto, es. la - \nprueba 10", "", TextArea()},
		{"TextArea - Letra ñ, paréntesis y $ permitidos", "letra ñ permitida? paréntesis y $", "soy ñato o Ñato (aqui) costo $10000.", "", TextArea()},
		{"TextArea - Solo texto y espacio", "solo texto y espacio?", "hola esto es una prueba", "", TextArea()},
		{"TextArea - Texto y puntuación", "texto y puntuación?", "hola: esto es una prueba", "", TextArea()},
		{"TextArea - Texto, puntuación y coma", "texto y puntuación y coma?", "hola: esto,true, es una prueba", "", TextArea()},
		{"TextArea - 5 caracteres", "5 caracteres?", ", .s5", "", TextArea()},
		{"TextArea - Sin datos no permitido", "sin data permitido?", "", Lang.T(D.Field, D.Empty, D.NotAllowed), TextArea()},
		{"TextArea - Carácter # permitido", "# permitido?", "# son", "", TextArea()},
		{"TextArea - Carácter ¿ no permitido", "¿ ? permitido?", " ¿ si ?", "carácter ¿ no permitido", TextArea()},
		{"TextArea - Tildes permitidas", "tildes si?", " mí tílde", "", TextArea()},
		{"TextArea - 1 carácter no permitido", "1 carácter?", "1", "tamaño mínimo " + strconv.Itoa(TextArea().Minimum) + " caracteres", TextArea()},
		{"TextArea - Nombre correcto", "nombre correcto?", "Dr. Pato Gomez", "", TextArea()},
		{"TextArea - Solo espacio no permitido", "solo espacio en blanco?", " ", "tamaño mínimo " + strconv.Itoa(TextArea().Minimum) + " caracteres", TextArea()},
		{"TextArea - Texto largo válido", "texto largo correcto?", `IRRITACION EN PIEL DE ROSTRO. ALERGIAS NO. CIRUGIAS NO. ACTUAL TTO CON ISOTRETINOINA 10MG - ENERO 2022. EN TTO ACTUAL CON VIT D. EXAMEN DE LAB 20-12-2022. SIN OTROS ANTECEDENTES`, "", TextArea()},
		{"TextArea - Texto con saltos de línea", "texto con salto de lineas ok", `HOY......Referido por        : dr. ........
			Motivo                    : ........
			Premedicacion  : ........`, "", TextArea()},

		// TextSearch tests
		{"TextSearch - Palabra de 15 caracteres", "palabra solo texto 15 caracteres?", "Maria Jose Diaz", "", TextSearch()},
		{"TextSearch - Texto con ñ permitido", "texto con ñ ok?", "Ñuñez perez", "", TextSearch()},
		{"TextSearch - Tilde no permitida", "tilde permitido?", "peréz del rozal", "é tilde no permitida", TextSearch()},
		{"TextSearch - Más de 20 caracteres no permitido", "mas de 20 caracteres permitidos?", "hola son mas de 21 ca", "tamaño máximo 20 caracteres", TextSearch()},
		{"TextSearch - Guion permitido", "guion permitido", "12038-0", "", TextSearch()},
		{"TextSearch - Fecha correcta", "fecha correcta?", "1990-07-21", "", TextSearch()},
		{"TextSearch - Fecha incorrecta permitida", "fecha incorrecta permitida?", "190-07-21", "", TextSearch()},

		// Validation tests
		{"Validation - Números sin espacio", "números sin espacio ok", "5648", "", input{permitted: permitted{Numbers: true}}},
		{"Validation - Números con espacio", "números con espacio ok", "5648 78212", "", input{permitted: permitted{Numbers: true, Characters: []rune{' '}}}},
		{"Validation - Error: espacio no permitido", "error no permitido números con espacio", "5648 78212", "espacio en blanco no permitido", input{permitted: permitted{Numbers: true}}},
		{"Validation - Solo texto sin espacio", "solo texto sin espacio ok", "Maria", "", input{permitted: permitted{Letters: true}}},
		{"Validation - Texto con espacios", "texto con espacios ok", "Maria De Lourdes", "", input{permitted: permitted{Letters: true, Characters: []rune{' '}}}},
		{"Validation - Texto con tildes y espacios", "texto con tildes y espacios ok", "María Dé Lourdes", "", input{permitted: permitted{Tilde: true, Letters: true, Characters: []rune{' '}}}},
		{"Validation - Texto con número sin espacios", "texto con numero sin espacios ok", "equipo01", "", input{permitted: permitted{Letters: true, Numbers: true}}},
		{"Validation - Número al inicio y texto sin espacios", "numero al inicio y texto sin espacios ok", "9equipo01", "", input{permitted: permitted{Letters: true, Numbers: true}}},
		{"Validation - Número al inicio y texto con espacios", "numero al inicio y texto con espacios ok", "9equipo01 2equipo2", "", input{permitted: permitted{Letters: true, Numbers: true, Characters: []rune{' '}}}},
		{"Validation - Error: solo números permitidos", "error solo números no letras si espacios", "9equipo01 2equipo2", "carácter e no permitido", input{permitted: permitted{Numbers: true, Characters: []rune{' '}}}},
		{"Validation - Correo con punto y @", "correo con punto y @ ok", "mi.correo1@mail.com", "", input{permitted: permitted{Characters: []rune{'@', '.'}, Numbers: true, Letters: true}}},
		{"Validation - Error: tilde no permitida en correo", "error correo con tilde no permitido", "mí.correo@mail.com", "í tilde no permitida", input{permitted: permitted{Characters: []rune{'@', '.'}, Numbers: true, Letters: true}}},
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
				log.Fatalf("result: [%v] expected: [%v]\n%v", err, test.expected, test.inputData)
			}
		})
	}
}
