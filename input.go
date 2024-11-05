package godi

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/cdvelop/godi/inputs"
)

type input interface {
	Render(tabIndex int) string
	Validate(value string) error
}

func (f *field) setInput(structureFrom reflect.Type, rf *reflect.StructField) {

	var params []any

	if structureFrom != nil {
		params = append(params, structureFrom)
	}

	inputData := rf.Tag.Get("Input")
	if inputData == "" {
		return
	}
	params = append(params, getParams(inputData))

	inputType := strings.Split(inputData, "(")[0]
	if inputType == "" {
		return
	}
	// fmt.Println("field name:", f.Name)
	params = append(params, `name=`+f.Name)
	params = append(params, `entity=`+f.Parent.Name)
	params = append(params, `legend=`+f.Legend)

	fmt.Println("setInput params:", params)
	// fmt.Println("inputType", inputType)

	switch inputType {
	case "Checkbox":
		f.Input = inputs.CheckBox(params...)
	case "DataList":
		f.Input = inputs.DataList(params...)
	case "Date":
		f.Input = inputs.Date(params...)
	case "DateAge":
		f.Input = inputs.DateAge(params...)
	case "DayWord":
		f.Input = inputs.DayWord(params...)
	case "FilePath":
		f.Input = inputs.FilePath(params...)
	case "Hour":
		f.Input = inputs.Hour(params...)
	case "Id":
		f.Input = inputs.ID(params...)
	case "Info":
		f.Input = inputs.Info(params...)
	case "Ip":
		f.Input = inputs.Ip(params...)
	case "List":
		f.Input = inputs.List(params...)
	case "Mail":
		f.Input = inputs.Mail(params...)
	case "MonthDay":
		f.Input = inputs.MonthDay(params...)
	case "Number":
		f.Input = inputs.Number(params...)
	case "Password":
		f.Input = inputs.Password(params...)
	case "Phone":
		f.Input = inputs.Phone(params...)
	case "Radio":
		f.Input = inputs.Radio(params...)
	case "RadioGender":
		f.Input = inputs.RadioGender(params...)
	case "Rut":
		f.Input = inputs.Rut(params...)
	case "Select":
		f.Input = inputs.Select(params...)
	case "TextArea":
		f.Input = inputs.TextArea(params...)
	case "TextNumber":
		f.Input = inputs.TextNumber(params...)
	case "TextNumberCode":
		f.Input = inputs.TextNumberCode(params...)
	case "TextOnly":
		f.Input = inputs.TextOnly(params...)
	case "TextSearch":
		f.Input = inputs.TextSearch(params...)
	default: //"Text"
		f.Input = inputs.Text(params...)
	}

}

func getParams(inputTag string) []string {
	start := strings.Index(inputTag, "(")
	end := strings.Index(inputTag, ")")
	if start == -1 || end == -1 {
		return nil
	}
	paramsStr := inputTag[start+1 : end]
	return strings.Split(paramsStr, ",")
}
