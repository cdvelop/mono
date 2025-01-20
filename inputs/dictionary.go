package inputs

type dictionary struct {
	Address           string `es:"dirección"`
	Allowed           string `es:"permitido"`
	BirthDate         string `es:"fecha de nacimiento"`
	Char              string `es:"carácter"`
	Chars             string `es:"caracteres"`
	City              string `es:"ciudad"`
	ConfirmPassword   string `es:"confirmar contraseña"`
	Country           string `es:"país"`
	Date              string `es:"fecha"`
	Day               string `es:"día"`
	Days              string `es:"días"`
	DoesNotExist      string `es:"no existe"`
	DoesNotHave       string `es:"no tiene"`
	DoNotStartWith    string `es:"no debe comenzar con"`
	Email             string `es:"correo electrónico"`
	Empty             string `es:"vacío"`
	Example           string `es:"ejemplo"`
	Field             string `es:"campo"`
	Format            string `es:"Formato"`
	Gender            string `es:"género"`
	Hour              string `es:"hora"`
	In                string `es:"en"`
	Is                string `es:"es"`
	LastName          string `es:"apellido"`
	Letters           string `es:"letras"`
	Max               string `es:"máx."`
	MaxSize           string `es:"tamaño máximo"`
	Min               string `es:"mín."`
	MinSize           string `es:"tamaño mínimo"`
	Name              string `es:"nombre"`
	Newline           string `es:"salto de linea"`
	NotAllowed        string `es:"no permitido"`
	NotLetter         string `es:"no es una letra"`
	NotFound          string `es:"no encontrado"`
	NotNumber         string `es:"no es un numero"`
	NotValid          string `es:"no es valido"`
	Numbers           string `es:"números"`
	Password          string `es:"contraseña"`
	Phone             string `es:"teléfono"`
	RequiredSelection string `es:"selección requerida"`
	Space             string `es:"espacio"`
	TabText           string `es:"tabulation de texto"`
	Text              string `es:"texto"`
	Terms             string `es:"términos y condiciones"`
	TildeNotAllowed   string `es:"tilde no permitida"`
	Value             string `es:"valor"`
	WhiteSpace        string `es:"espacio en blanco"`
	ZipCode           string `es:"código postal"`
	YearOutOfRange    string `es:"año fuera de rango"`
	MonthOutOfRange   string `es:"mes fuera de rango"`
	DayCannotBeZero   string `es:"día no puede ser cero"`
	InvalidDateFormat string `es:"formato de fecha ingresado incorrecto"`
	Year              string `es:"año"`
	Month             string `es:"mes"`
	January           string `es:"Enero"`
	February          string `es:"Febrero"`
	March             string `es:"Marzo"`
	April             string `es:"Abril"`
	May               string `es:"Mayo"`
	June              string `es:"Junio"`
	July              string `es:"Julio"`
	August            string `es:"Agosto"`
	September         string `es:"Septiembre"`
	October           string `es:"Octubre"`
	November          string `es:"Noviembre"`
	December          string `es:"Diciembre"`
}
