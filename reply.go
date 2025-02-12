package mono

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// var R reply
var R reply

type reply struct {
	current      string // "es" or "en"
	translations map[string]map[string]string
	out          strings.Builder
	err          errMessage
}

type errMessage struct {
	message string
}

// var d dictionary
var D dictionary

func init() {

	langSupported := []string{"es"}

	R = reply{
		current: "es",
		translations: map[string]map[string]string{
			"en": {}, // do not complete manually!
		},
		err: errMessage{
			message: "",
		},
	}

	// initialize translations map
	for _, lang := range langSupported {
		R.translations[lang] = map[string]string{}
	}

	v := reflect.ValueOf(&D).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		dbFieldType := t.Field(i)

		if field.CanSet() {
			// Convert field name to: snake case
			snakeCaseName := G.String.SnakeCase(dbFieldType.Name)
			// Assign field name to dictionary structure
			field.SetString(snakeCaseName)
			// Separate words
			separateName := G.String.SnakeCase(dbFieldType.Name, " ")
			// Update translations map
			R.translations["en"][snakeCaseName] = separateName
			for _, reply := range langSupported {
				// Get tags for other languages "es","pt"
				esTag := dbFieldType.Tag.Get(reply)
				if esTag != "" {
					R.translations[reply][snakeCaseName] = esTag
				}
			}
		}
	}
	// fmt.Println("dictionary initialized", R)
}

// Set set the language eg: "es", "en", "pt", "fr"
func (l *reply) Set(reply string) {
	l.current = reply
}

// T returns the translation of the given arguments.
// eg: R.T("hello", "world") returns "hello world"
func (l *reply) T(args ...interface{}) string {
	l.out.Reset()
	var space string
	for _, arg := range args {
		switch v := arg.(type) {
		case string:
			if v == "" {
				continue
			}

			if trans, ok := l.translations[l.current][v]; ok {
				l.out.WriteString(space + trans)
			} else {
				l.out.WriteString(space + v)
			}
		case []string:
			for _, s := range v {
				if s == "" {
					continue
				}
				if trans, ok := l.translations[l.current][s]; ok {
					l.out.WriteString(space + trans)
				} else {
					l.out.WriteString(space + s)
				}
				space = " "
			}
		case rune:

			if v == ':' {
				l.out.WriteString(":")
				continue
			}

			l.out.WriteString(space + string(v))
		case int:
			l.out.WriteString(space + strconv.Itoa(v))
		case float64:
			l.out.WriteString(space + strconv.FormatFloat(v, 'f', -1, 64))
		case bool:
			l.out.WriteString(space + strconv.FormatBool(v))
		case error:
			l.out.WriteString(space + v.Error())
		default:
			l.out.WriteString(space + fmt.Sprint(v))
		}
		space = " "
	}
	return l.out.String()
}

func (l reply) Err(args ...any) error {
	l.T(args...)
	l.err.message = l.out.String()
	return l.err
}

func (e errMessage) Error() string {
	return e.message
}
