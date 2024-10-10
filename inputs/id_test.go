package inputs

import (
	"fmt"
	"log"
	"testing"
)

var (
	idTestData = map[string]struct {
		input     id
		inputData string
		expected  string
	}{
		"id 1 correcto?":                   {ID(), "1624397134562544800", ""},
		"id 2 ok?":                         {ID(), "1624397172303448900", ""},
		"id 3 ok?":                         {ID(), "1634394443466878800", ""},
		"id 4 con punto usuario numero ok": {ID(), "1624397134562544800.30", ""},
		"id 5 con - usuario numero error":  {ID(), "1624397134562544800-30", "carácter - no permitido"},
		"numero 5 correcto?":               {ID(), "5", ""},
		"numero 45 correcto?":              {ID(), "45", ""},
		"id con letra no valido":           {ID(), "E624397172303448900", "carácter E no permitido"},
		"id con letra valido":              {ID("letters"), "E624397172303448900", ""},
		"id con carácter valido":           {ID("chars='$','-'"), "$624397172303448900-2", ""},
		"id con carácter no validos":       {ID(), "$624397172303448900-2", "carácter $ no permitido"},
		"primary key se permite vació ?":   {ID(), "", "tamaño mínimo 1 caracteres"},
		"id cero?":                         {ID(), "0", ""},
	}
)

// 9223372036854775807

func Test_InputPrimaryKey(t *testing.T) {

	for prueba, data := range idTestData {
		t.Run((prueba + ": " + data.inputData), func(t *testing.T) {
			err := data.input.Validate(data.inputData)
			var resp string
			if err != nil {
				resp = err.Error()
			}

			if resp != data.expected {
				log.Println(prueba)
				log.Fatalf("resultado: [%v] expectativa: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}

func Test_TagPrimaryKey(t *testing.T) {
	tag := ID().BuildHtmlInput("1")
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}

	fmt.Println(tag)
}

func Test_GoodInputPrimaryKey(t *testing.T) {
	for _, data := range ID().GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := ID().Validate(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputPrimaryKey(t *testing.T) {
	for _, data := range ID().WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := ID().Validate(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
