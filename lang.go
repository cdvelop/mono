package monogo

import "strconv"

type lang struct {
	Current string // "es" or "en"
}

var Lang = lang{
	Current: "es",
}

var translations = map[string]map[string]string{
	"en": {
		"allowed":                  "allowed:",
		"not_allowed":              " not allowed",
		"character":                "character ",
		"characters":               "characters:",
		"chars":                    " characters",
		"letters":                  "letters",
		"max":                      "max.",
		"max_size":                 "maximum size ",
		"min":                      "min.",
		"min_size":                 "minimum size ",
		"newline_not_allowed":      "line break not allowed",
		"not_letter":               " is not a letter",
		"not_number":               " is not a number",
		"numbers":                  "numbers",
		"tab_not_allowed":          "text tabulation not allowed",
		"tilde_not_allowed":        " with tilde not allowed",
		"unsupported_tilde":        "unsupported tilde ",
		"white_spaces_not_allowed": "white spaces not allowed",
		"name":                     "name",
		"last_name":                "last name",
		"address":                  "address",
		"email":                    "email",
		"phone":                    "phone",
		"password":                 "password",
		"confirm_password":         "confirm password",
		"city":                     "city",
		"country":                  "country",
		"zip_code":                 "zip code",
		"date_of_birth":            "date of birth",
		"gender":                   "gender",
		"terms":                    "terms and conditions",
		"birth_date":               "birth date",
		"id":                       "id",
	},
	"es": {
		"allowed":                  "permitido:",
		"not_allowed":              " no permitido",
		"character":                "carácter ",
		"characters":               "caracteres:",
		"chars":                    " caracteres",
		"letters":                  "letras",
		"max":                      "máx.",
		"max_size":                 "tamaño máximo ",
		"min":                      "mín.",
		"min_size":                 "tamaño mínimo ",
		"newline_not_allowed":      "salto de linea no permitido",
		"not_letter":               " no es una letra",
		"not_number":               " no es un numero",
		"numbers":                  "números",
		"tab_not_allowed":          "tabulation de texto no permitida",
		"tilde_not_allowed":        " con tilde no permitida",
		"unsupported_tilde":        "tilde ",
		"white_spaces_not_allowed": "espacios en blanco no permitidos",
		"name":                     "nombre",
		"last_name":                "apellido",
		"address":                  "dirección",
		"email":                    "correo electrónico",
		"phone":                    "teléfono",
		"password":                 "contraseña",
		"confirm_password":         "confirmar contraseña",
		"city":                     "ciudad",
		"country":                  "país",
		"zip_code":                 "código postal",
		"date_of_birth":            "fecha de nacimiento",
		"gender":                   "género",
		"terms":                    "términos y condiciones",
		"birth_date":               "fecha de nacimiento",
		"id":                       "id",
	},
}

func (l lang) T(key string) string {
	if trans, ok := translations[l.Current][key]; ok {
		return trans
	}
	return key
}

func (l lang) TNum(key string, num int) string {
	return l.T(key) + strconv.Itoa(num) + l.T("chars")
}

func (l lang) TChar(key, char string) string {
	return l.T("character") + char + l.T(key)
}
