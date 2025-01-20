package inputs

import (
	"log"
	"testing"
)

var (
	modelDate = Date()

	dataDate = map[string]struct {
		inputData string
		expected  string
	}{
		"correcto ":                      {"2002-12-03", ""},
		"dia 29 febrero año bisiesto":    {"2020-02-29", ""},
		"dia 29 febrero año no bisiesto": {"2023-02-29", Lang.T(D.February, D.DoesNotExist, "29", D.Days) + ". " + Lang.T(D.Year, "2023", D.Is, D.NotValid) + "."},
		"junio no tiene 31 días":         {"2023-06-31", Lang.T(D.June, D.DoesNotExist, "31", D.Days) + "."},
		"carácter de mas incorrecto ":    {"2002-12-03-", Lang.T(D.InvalidDateFormat, "2006-01-02")},
		"formato incorrecto ":            {"21/12/1998", Lang.T(D.InvalidDateFormat, "2006-01-02")},
		"fecha incorrecta ":              {"2020-31-01", " error 31 " + Lang.T(D.Is, D.InvalidDateFormat, D.Month)},
		"fecha recortada sin año ok?":    {"31-01", Lang.T(D.InvalidDateFormat, "2006-01-02")},
		"data incorrecta ":               {"0000-00-00", " error 00 " + Lang.T(D.Is, D.InvalidDateFormat, D.Month)},
		"toda la data correcta?":         {"", Lang.T(D.InvalidDateFormat, "2006-01-02")},
	}
)

func Test_InputDate(t *testing.T) {
	for prueba, data := range dataDate {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelDate.Validate(data.inputData)

			var err_str string
			if err != nil {
				err_str = err.Error()
			}

			if err_str != data.expected {
				log.Println(prueba)
				t.Fatalf("result: [%v] expected: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}

func Test_TagDate(t *testing.T) {
	tag := modelDate.Render(1)
	if tag == "" {
		t.Fatal("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputDate(t *testing.T) {
	for _, data := range modelDate.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelDate.Validate(data); ok != nil {
				t.Fatalf("result [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputDate(t *testing.T) {
	for _, data := range modelDate.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelDate.Validate(data); ok == nil {
				t.Fatalf("result [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_InputMonthDay(t *testing.T) {
	var (
		modelMonthDay = MonthDay()

		dataMonthDay = map[string]struct {
			inputData string
			expected  string
		}{
			"day ok?":               {"01", ""},
			"characters?":           {"0l", Lang.T(D.NotNumber)},
			"month ok?":             {"31", ""},
			"date without year ok?": {"31-01", Lang.T(D.MaxSize, "2", D.Chars)},
			"incorrect":             {"2002-12-03", Lang.T(D.MaxSize, "2", D.Chars)},
			"incorrect format":      {"21/12", Lang.T(D.MaxSize, "2", D.Chars)},
			"incorrect data":        {"0000-00-00", Lang.T(D.MaxSize, "2", D.Chars)},
			"empty data":            {"", Lang.T(D.Field, D.Empty, D.NotAllowed)},
			"incorrect 32 day":      {"32", Lang.T(D.Field, D.Day, D.NotValid)},
		}
	)
	for prueba, data := range dataMonthDay {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := modelMonthDay.Validate(data.inputData)

			var err_str string
			if err != nil {
				err_str = err.Error()
			}

			if err_str != data.expected {
				t.Fatalf("result: [%v] expected: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}

func Test_TagMonthDay(t *testing.T) {
	tag := MonthDay().Render(1)
	if tag == "" {
		t.Fatal("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputMonthDay(t *testing.T) {
	for _, data := range MonthDay().GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := MonthDay().Validate(data); ok != nil {
				t.Fatalf("result [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputMonthDay(t *testing.T) {
	for _, data := range MonthDay().WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := MonthDay().Validate(data); ok == nil {
				t.Fatalf("result [%v] [%v]", ok, data)
			}
		})
	}
}
