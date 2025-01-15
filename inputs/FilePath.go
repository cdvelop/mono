package inputs

import (
	"errors"
	"strings"
)

// options:
// "multiple"
// accept="image/*"
// title="Imágenes jpg"
func FilePath(params ...any) *filePath {
	new := &filePath{
		input: input{
			attributes: attributes{
				htmlName:   "file",
				customName: "FilePath",
			},
			permitted: permitted{
				Letters:    true,
				Tilde:      false,
				Numbers:    true,
				Characters: []rune{'\\', '/', '.', '_'},
				Minimum:    1,
				Maximum:    100,
			},
		},
	}
	new.Set(params)

	return new
}

type filePath struct {
	input
}

var errPath = errors.New("la ruta no puede comenzar con \\ o / ")

// validación con datos de entrada
func (f filePath) Validate(value string) error {
	if value == "" {
		return errors.New("la ruta no puede estar vacía")
	}

	if value[0] == '\\' {
		return errPath
	}

	// Reemplazar las barras diagonales hacia adelante con barras diagonales hacia atrás.
	value = strings.ReplaceAll(value, "/", "\\")

	// Eliminar barras diagonales dobles al principio y al final de la cadena.
	value = strings.ReplaceAll(value, "\\", "")

	// Dividir la cadena en partes utilizando las barras diagonales como delimitadores.
	parts := strings.Split(value, "\\")

	for _, part := range parts {
		err := f.permitted.Validate(part)
		if err != nil {
			return err
		}
	}

	// Verificar que la ruta sea válida para Linux y Windows
	return nil
}

func (f filePath) GoodTestData() (out []string) {

	return []string{
		".\\misArchivos",
		".\\todos\\videos",
	}
}

func (f filePath) WrongTestData() (out []string) {
	out = []string{
		"\\-",
		"///.",
	}
	return
}
