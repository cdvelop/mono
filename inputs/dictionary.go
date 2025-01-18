package inputs

type dictionary struct {
	Allowed               string `es:"permitido:"`
	NotAllowed            string `es:" no permitido"`
	Character             string `es:"carácter "`
	Characters            string `es:"caracteres:"`
	Chars                 string `es:" caracteres"`
	Letters               string `es:"letras"`
	Max                   string `es:"máx."`
	MaxSize               string `es:"tamaño máximo "`
	Min                   string `es:"mín."`
	MinSize               string `es:"tamaño mínimo "`
	NewlineNotAllowed     string `es:"salto de linea no permitido"`
	NotLetter             string `es:" no es una letra"`
	NotNumber             string `es:" no es un numero"`
	Numbers               string `es:"números"`
	TabNotAllowed         string `es:"tabulation de texto no permitida"`
	TildeNotAllowed       string `es:" con tilde no permitida"`
	UnsupportedTilde      string `es:"tilde "`
	WhiteSpacesNotAllowed string `es:"espacios en blanco no permitidos"`
	Name                  string `es:"nombre"`
	LastName              string `es:"apellido"`
	Address               string `es:"dirección"`
	Email                 string `es:"correo electrónico"`
	Phone                 string `es:"teléfono"`
	Password              string `es:"contraseña"`
	ConfirmPassword       string `es:"confirmar contraseña"`
	City                  string `es:"ciudad"`
	Country               string `es:"país"`
	ZipCode               string `es:"código postal"`
	DateOfBirth           string `es:"fecha de nacimiento"`
	Gender                string `es:"género"`
	Terms                 string `es:"términos y condiciones"`
	BirthDate             string `es:"fecha de nacimiento"`
	Id                    string `es:"id"`
	WhiteSpaces           string `es:"espacios en blanco"`
	DoNotStartWith        string `es:"no debe comenzar con "`
	Space                 string `es:"espacio"`
}
