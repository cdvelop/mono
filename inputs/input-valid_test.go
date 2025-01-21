package inputs

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

type testStruct struct {
	description string
	inputData   string
	expected    string
	model       interface{ Validate(string) error }
}

func Test_AllInputs(t *testing.T) {
	tests := []testStruct{
		// DataList tests
		{"una credencial ok?", "1", "", DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"otro numero ok?", "3", "", DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"0 existe?", "0", "valor 0 no permitido en datalist credentials", DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"-1 valido?", "-1", "valor -1 no permitido en datalist credentials", DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"carácter permitido?", "%", "valor % no permitido en datalist credentials", DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"con data?", "", "selección requerida campo credentials", DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"sin espacios?", "luis ", "valor luis  no permitido en datalist credentials", DataList("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},

		// Date tests
		{"correct", "2002-12-03", "", Date()},
		{"February 29th leap year", "2020-02-29", "", Date()},
		{"February 29th non leap year", "2023-02-29", Lang.T(D.February, D.DoesNotHave, "29", D.Days, D.Year, "2023"), Date()},
		{"June does not have 31 days", "2023-06-31", Lang.T(D.June, D.DoesNotHave, "31", D.Days), Date()},
		{"incorrect extra character", "2002-12-03-", Lang.T(D.InvalidDateFormat, "2006-01-02"), Date()},
		{"incorrect format", "21/12/1998", Lang.T(D.InvalidDateFormat, "2006-01-02"), Date()},
		{"incorrect month 31", "2020-31-01", Lang.T(D.Month, 31, D.NotValid), Date()},
		{"shortened date without year ok?", "31-01", Lang.T(D.InvalidDateFormat, "2006-01-02"), Date()},
		{"incorrect data", "0000-00-00", Lang.T(D.InvalidDateFormat), Date()},
		{"all data correct?", "", Lang.T(D.InvalidDateFormat, "2006-01-02"), Date()},

		// FilePath tests
		{"correct path", ".\\files\\1234\\", "", FilePath()},
		{"correct path without starting point?", "\\files\\1234\\", Lang.Err(D.DoNotStartWith, '\\').Error(), FilePath()},
		{"absolute path in Linux", "/home/user/files/", "", FilePath()},
		{"absolute path without starting point", "/files/1234/", "", FilePath()},
		{"relative path without directories ok?", ".\\", "", FilePath()},
		{"relative path with final slash", "./files/1234/", "", FilePath()},
		{"path with filename", ".\\files\\1234\\archivo.txt", "", FilePath()},
		{"path with directory names using underscores", ".\\mi_directorio\\sub_directorio\\", "", FilePath()},
		{"is a number a valid path?", "5", "", FilePath()},
		{"is a single word a valid path?", "path", "", FilePath()},
		{"white spaces allowed?", ".\\path with white space\\", Lang.T(D.WhiteSpace, D.NotAllowed), FilePath()},

		// IP tests
		{"IPv4 ok", "192.168.1.1", "", Ip()},
		{"IPv6 ok", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", "", Ip()},
		{"ip incorrecta", "192.168.1.1.8", Lang.T(D.Format, "IPv4", D.NotValid), Ip()},
		{"correcto?", "0.0.0.0", Lang.T(D.Example, "IP", D.NotAllowed, ':', "0.0.0.0"), Ip()},
		{"sin data", "", Lang.T(D.Field, D.Empty, D.NotAllowed), Ip()},

		// Mail tests
		{"correo normal", "mi.correo@mail.com", "", Mail()},
		{"correo un campo", "correo@mail.com", "", Mail()},

		// Number tests
		{"correct number 100", "100", "", Number()},
		{"single digit 0", "0", "", Number()},
		{"single digit 1", "1", "", Number()},
		{"uint64 20 chars", "18446744073709551615", "", Number()},
		{"uint64 21 chars", "184467440737095516150", Lang.T(D.MaxSize, 20, D.Chars), Number()},
		{"int64 19 chars", "9223372036854775807", "", Number()},
		{"int32 10 chars", "2147483647", "", Number()},
		{"18 digits", "100002323262637278", "", Number()},
		{"large number with letter", "10000232E26263727", Lang.T('E', D.NotNumber), Number()},
		{"negative number -100", "-100", Lang.T('-', D.NotNumber), Number()},
		{"text instead of number", "h", Lang.T('h', D.NotNumber), Number()},
		{"text and number", "h1", Lang.T('h', D.NotNumber), Number()},

		// Password tests
		{"validates numbers letters and character", "c0ntra3", "", Password()},
		{"validates many characters", "M1 contraseÑ4", "", Password()},
		{"validates 8 characters", "contrase", "", Password()},
		{"validates 5 characters", "contñ", "", Password()},
		{"validates only numbers", "12345", "", Password()},
		{"does not validate less than 2", "1", Lang.T(D.MinSize, 5, D.Chars), Password()},
		{"no data", "", Lang.T(D.Field, D.Empty, D.NotAllowed), Password()},

		// Radio tests
		{"D Dato correcto", "D", "", Radio("name=genre", "options=D:Dama,V:Varón")},
		{"V Dato correcto", "V", "", Radio("name=genre", "options=D:Dama,V:Varón")},
		{"d Dato en minúscula incorrecto", "d", "valor d no permitido en radio genre", Radio("name=genre", "options=D:Dama,V:Varón")},
		{"v Dato en minúscula incorrecto", "v", "valor v no permitido en radio genre", Radio("name=genre", "options=D:Dama,V:Varón")},
		{"data ok?", "0", "valor 0 no permitido en radio genre", Radio("name=genre", "options=D:Dama,V:Varón")},

		// Rut tests
		{"without hyphen 15890022k", "15890022k", Lang.T(D.HyphenMissing), Rut()},
		{"no hyphen 177344788", "177344788", Lang.T(D.HyphenMissing), Rut()},
		{"ok 7863697-1", "7863697-1", "", Rut()},
		{"ok uppercase K 20373221-K", "20373221-k", "", Rut()},
		{"ok lowercase k 20373221-k", "20373221-k", "", Rut()},
		{"valid run? allowed?", "7863697-W", Lang.T(D.Digit, D.Verifier, "w", D.NotValid), Rut()},
		{"change digit to k 7863697-k", "7863697-k", Lang.T(D.Digit, D.Verifier, "k", D.NotValid), Rut()},
		{"change digit to 0 7863697-0", "7863697-0", Lang.T(D.Digit, D.Verifier, 0, D.NotValid), Rut()},
		{"ok 14080717-6", "14080717-6", "", Rut()},
		{"incorrect 14080717-0", "14080717-0", Lang.T(D.Digit, D.Verifier, 0, D.NotValid), Rut()},
		{"correct zero at start?", "07863697-1", Lang.T(D.DoNotStartWith, D.Digit, 0), Rut()},
		{"correct data only space?", " ", Lang.T(D.DoNotStartWith, D.WhiteSpace), Rut()},

		// Select tests
		{"una credencial ok?", "1", "", Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"otro numero ok?", "3", "", Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"0 existe?", "0", "valor 0 no permitido en select credentials", Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"-1 valido?", "-1", "valor -1 no permitido en select credentials", Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"carácter permitido?", "%", "valor % no permitido en select credentials", Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"con data?", "", "selección requerida campo credentials", Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},
		{"sin espacios?", "luis ", "valor luis  no permitido en select credentials", Select("name=credentials", "options=1:Admin,2:Editor,3:Visitante")},

		// Text tests
		{"nombre correcto con punto?", "Dr. Maria Jose Diaz Cadiz", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"no tilde", "peréz del rozal", "é tilde no permitida", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"texto con ñ", "Ñuñez perez", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"texto correcto + 3 caracteres", "hola", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"texto correcto 3 caracteres", "los", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"oración ok", "hola que tal", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"solo Dato numérico permitido?", "100", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"con caracteres y coma", "los,true, vengadores", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"sin data ok se permite campo vació", "", "", Text(`!required`)},
		{"no se permite campo vació", "8", "tamaño mínimo 2 caracteres", Text(`required`)},
		{"palabra mas numero permitido", "son 4 bidones", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"con paréntesis y numero", "son 4 (4 bidones)", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"con solo paréntesis", "son (bidones)", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"palabras y numero", "apellido Actualizado 1", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"palabra con slash?", " estos son \\n los podria", "carácter \\ no permitido", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},
		{"nombre de archivos separados por ,", "dino.png, gatito.jpeg", "", Text(`name=full_name;hidden;!required; data=price:100 ;placeholder=tu nombre`)},

		// TextArea tests
		{"todo los caracteres permitidos?", "hola: esto, es. la - prueba 10", "", TextArea()},
		{"salto de linea permitido? y guion", "hola:\n esto, es. la - \nprueba 10", "", TextArea()},
		{"letra ñ permitida? paréntesis y $", "soy ñato o Ñato (aqui) costo $10000.", "", TextArea()},
		{"solo texto y espacio?", "hola esto es una prueba", "", TextArea()},
		{"texto y puntuación?", "hola: esto es una prueba", "", TextArea()},
		{"texto y puntuación y coma?", "hola: esto,true, es una prueba", "", TextArea()},
		{"5 caracteres?", ", .s5", "", TextArea()},
		{"sin data permitido?", "", Lang.T(D.Field, D.Empty, D.NotAllowed), TextArea()},
		{"# permitido?", "# son", "", TextArea()},
		{"¿ ? permitido?", " ¿ si ?", "carácter ¿ no permitido", TextArea()},
		{"tildes si?", " mí tílde", "", TextArea()},
		{"1 carácter?", "1", "tamaño mínimo " + strconv.Itoa(TextArea().Minimum) + " caracteres", TextArea()},
		{"nombre correcto?", "Dr. Pato Gomez", "", TextArea()},
		{"solo espacio en blanco?", " ", "tamaño mínimo " + strconv.Itoa(TextArea().Minimum) + " caracteres", TextArea()},
		{"texto largo correcto?", `IRRITACION EN PIEL DE ROSTRO. ALERGIAS NO. CIRUGIAS NO. ACTUAL TTO CON ISOTRETINOINA 10MG - ENERO 2022. EN TTO ACTUAL CON VIT D. EXAMEN DE LAB 20-12-2022. SIN OTROS ANTECEDENTES`, "", TextArea()},
		{"texto con salto de lineas ok", `HOY......Referido por        : dr. ........
			Motivo                    : ........
			Premedicacion  : ........`, "", TextArea()},

		// TextSearch tests
		{"palabra solo texto 15 caracteres?", "Maria Jose Diaz", "", TextSearch()},
		{"texto con ñ ok?", "Ñuñez perez", "", TextSearch()},
		{"tilde permitido?", "peréz del rozal", "é tilde no permitida", TextSearch()},
		{"mas de 20 caracteres permitidos?", "hola son mas de 21 ca", "tamaño máximo 20 caracteres", TextSearch()},
		{"guion permitido", "12038-0", "", TextSearch()},
		{"fecha correcta?", "1990-07-21", "", TextSearch()},
		{"fecha incorrecta permitida?", "190-07-21", "", TextSearch()},

		// Validation tests
		{"números sin espacio ok", "5648", "", input{permitted: permitted{Numbers: true}}},
		{"números con espacio ok", "5648 78212", "", input{permitted: permitted{Numbers: true, Characters: []rune{' '}}}},
		{"error no permitido números con espacio", "5648 78212", "espacio en blanco no permitido", input{permitted: permitted{Numbers: true}}},
		{"solo texto sin espacio ok", "Maria", "", input{permitted: permitted{Letters: true}}},
		{"texto con espacios ok", "Maria De Lourdes", "", input{permitted: permitted{Letters: true, Characters: []rune{' '}}}},
		{"texto con tildes y espacios ok", "María Dé Lourdes", "", input{permitted: permitted{Tilde: true, Letters: true, Characters: []rune{' '}}}},
		{"texto con numero sin espacios ok", "equipo01", "", input{permitted: permitted{Letters: true, Numbers: true}}},
		{"numero al inicio y texto sin espacios ok", "9equipo01", "", input{permitted: permitted{Letters: true, Numbers: true}}},
		{"numero al inicio y texto con espacios ok", "9equipo01 2equipo2", "", input{permitted: permitted{Letters: true, Numbers: true, Characters: []rune{' '}}}},
		{"error solo números no letras si espacios", "9equipo01 2equipo2", "carácter e no permitido", input{permitted: permitted{Numbers: true, Characters: []rune{' '}}}},
		{"correo con punto y @ ok", "mi.correo1@mail.com", "", input{permitted: permitted{Characters: []rune{'@', '.'}, Numbers: true, Letters: true}}},
		{"error correo con tilde no permitido", "mí.correo@mail.com", "í tilde no permitida", input{permitted: permitted{Characters: []rune{'@', '.'}, Numbers: true, Letters: true}}},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			err := test.model.Validate(test.inputData)

			var err_str string
			if err != nil {
				err_str = err.Error()
			}

			if err_str != test.expected {
				fmt.Println(test.description)
				log.Fatalf("resultado: [%v] expectativa: [%v]\n%v", err, test.expected, test.inputData)
			}
		})
	}
}
