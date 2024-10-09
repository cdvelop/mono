package inputs

import (
	"log"
	"testing"
)

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
		"error correo con tilde no permitido":       {"mí.correo@mail.com", "í con tilde no permitida", permitted{Characters: []rune{'@', '.'}, Numbers: true, Letters: true}},
	}
)

func Test_Valid(t *testing.T) {

	for prueba, data := range validTestData {
		t.Run((prueba + " " + data.text), func(t *testing.T) {
			err := data.Validate(data.text)

			var err_str string
			if err != nil {
				err_str = err.Error()
			}

			if err_str != data.expected {
				log.Println(prueba)
				log.Fatalf("expectativa [%v] resultado [%v]\n%v", data.expected, err, data.text)
			}

		})
	}
}
