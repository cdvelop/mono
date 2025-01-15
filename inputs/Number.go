package inputs

import (
	"strconv"
)

// options ej: data-type, data-after=" Años"
// hidden, el campo se mantendrá oculto
// title="xxx"
// for phone ej: `min="7"`, `max="11"`
func Number(params ...any) *number {

	new := &number{
		input: input{
			attributes: attributes{
				htmlName: "number",
				// Title:    `title="solo valores numéricos positivos >= 0 máximo 20 char 18446744073709551615"`,
				// Pattern: `^[0-9]{1,20}`,
			},
			permitted: permitted{
				Numbers: true,
				Minimum: 1,
				Maximum: 20,
			},
		},
	}
	new.Set(params)

	if new.Min != "" {
		new.Minimum, _ = strconv.Atoi(new.Min)
	}

	if new.Max != "" {
		new.Maximum, _ = strconv.Atoi(new.Max)
	}

	return new
}

type number struct {
	input
}

// func (n number) FieldAddEventListener(fieldName string) string {
// 	return fmt.Sprintf(`input_%v.addEventListener("input", InputValidationWithPattern);`, fieldName)
// }

// func (n number) FieldRemoveEventListener(fieldName string) string {
// 	return fmt.Sprintf(`input_%v.removeEventListener("input", InputValidationWithPattern);`, fieldName)
// }

func (n number) GoodTestData() (out []string) {

	temp := []string{
		"56988765432",
		"1234567",
		"0",
		"123456789",
		"100",
		"5000",
		"423456789",
		"31",
		"523756789",
		"10000232326263727",
		"29",
		"923726789",
		"3234567",
		"823456789",
		"29",
	}

	for _, v := range temp {
		if len(v) >= n.Minimum && len(v) <= n.Maximum {
			out = append(out, v)
		}
	}

	return
}

func (n number) WrongTestData() (out []string) {
	out = []string{"1-1", "-100", "h", "h1", "-1", " ", "", "#", "& ", "% &"}

	return
}
