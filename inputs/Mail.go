package inputs

import (
	"errors"
	"strings"
)

func Mail(params ...any) mail {
	new := mail{
		input: input{
			attributes: attributes{
				htmlName:    "mail",
				PlaceHolder: `placeHolder="ej: mi.correo@mail.com"`,
				// Pattern:     `[a-zA-Z0-9!#$%&'*_+-]([\.]?[a-zA-Z0-9!#$%&'*_+-])+@[a-zA-Z0-9]([^@&%$\/()=?¿!.,:;]|\d)+[a-zA-Z0-9][\.][a-zA-Z]{2,4}([\.][a-zA-Z]{2})?`,
			},
			permitted: permitted{
				Letters:    true,
				Numbers:    true,
				Characters: []rune{'@', '.', '_'},
				Minimum:    0,
				Maximum:    0,
			},
		},
	}
	new.Set(params)

	return new
}

type mail struct {
	input
}

// validación con datos de entrada
func (m mail) Validate(value string) error {

	if strings.Contains(value, "example") {
		return errors.New(value + " es un correo de ejemplo")
	}

	parts := strings.Split(value, "@")
	if len(parts) != 2 {
		return errors.New("error en @ del correo " + value)
	}

	return m.permitted.Validate(value)

}

func (mail) GoodTestData() (out []string) {

	out = []string{
		"mi.correo@mail.com",
		"alguien@algunlugar.es",
		"ramon.bonachea@email.com",
		"r.bonachea@email.com",
		"laura@hellos.email.tk",
	}

	return
}

func (mail) WrongTestData() (out []string) {

	out = []string{
		"email@example.com",
		"correomao@n//.oo",
		"son_24_bcoreos",
		"email@example",
	}
	out = append(out, wrong_data...)

	return
}
