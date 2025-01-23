package mono

type input struct {
	attributes
	permitted
	dataSource
	fieldset
}

type attributes struct {
	Id   string
	Name string //eg address,phone

	Type     string // input type  eg : text, password, email
	htmlName string //eg input,select,textarea

	allowSkipCompleted bool //permite que el campo no sea completado
	hideTyping         bool //oculta el valor mientras se escribe

	entity string //eg: user, product, order
	legend string

	PlaceHolder string
	Title       string //info

	Min string //valor mínimo
	Max string //valor máximo

	Maxlength string //ej: maxlength="12"

	Autocomplete string

	Rows string //filas ej 4,5,6
	Cols string //columnas ej 50,80

	Step     string
	Oninput  string // ej: "miRealtimeFunction()" = oninput="miRealtimeFunction()"
	Onkeyup  string // ej: "miNormalFuncion()" = onkeyup="miNormalFuncion()"
	Onchange string // ej: "miNormalFuncion()" = onchange="myFunction()"

	// https://developer.mozilla.org/en-US/docs/Web/HTML/attributes/accept
	// https://developer.mozilla.org/es/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types
	// accept="image/*"
	Accept   string
	Multiple string // multiple

	Value string

	customName string //eg onlyText,onlyNumber...

	Class []className // clase css ej: class="age"

	DataSet []map[string]string // dataset ej: data-id="123" = map[string]string{"id": "123"}

	options []map[string]string // ej: [{"m": "male"}, { "f": "female"}]

}

type className string

type permitted struct {
	Letters         bool
	Tilde           bool
	Numbers         bool
	BreakLine       bool     // line breaks allowed
	WhiteSpaces     bool     // white spaces allowed
	Tabulation      bool     // tabulation allowed
	TextNotAllowed  []string // text not allowed eg: "hola" not allowed
	Characters      []rune   // other special characters eg: '\','/','@'
	Minimum         int      // min characters eg 2 "lo" ok default 0 no defined
	Maximum         int      // max characters eg 1 "l" ok default 0 no defined}
	ExtraValidation func(string) error
	StartWith       *permitted // characters allowed at the beginning
}

type fieldset struct {
	cssClasses []className
}

type container struct {
	cssClasses []className
}

type sourceData interface {
	DataSource() any
}

type dataSource struct {
	// 	DataSource() []map[string]string
	data sourceData
}
