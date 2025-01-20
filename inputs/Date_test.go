package inputs

import (
	"log"
	"testing"
)

func Test_InputDate(t *testing.T) {
	var (
		modelDate = Date()

		dataDate = map[string]struct {
			inputData string
			expected  string
		}{
			"correct ":                        {"2002-12-03", ""},
			"February 29th leap year":         {"2020-02-29", ""},
			"February 29th non leap year":     {"2023-02-29", Lang.T(D.February, D.DoesNotHave, "29", D.Days, D.Year, "2023")},
			"June does not have 31 days":      {"2023-06-31", Lang.T(D.June, D.DoesNotHave, "31", D.Days)},
			"incorrect extra character ":      {"2002-12-03-", Lang.T(D.InvalidDateFormat, "2006-01-02")},
			"incorrect format ":               {"21/12/1998", Lang.T(D.InvalidDateFormat, "2006-01-02")},
			"incorrect month 31":              {"2020-31-01", Lang.T(D.Month, 31, D.NotValid)},
			"shortened date without year ok?": {"31-01", Lang.T(D.InvalidDateFormat, "2006-01-02")},
			"incorrect data ":                 {"0000-00-00", Lang.T(D.InvalidDateFormat)},
			"all data correct?":               {"", Lang.T(D.InvalidDateFormat, "2006-01-02")}}
	)

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
	tag := Date().Render(1)
	if tag == "" {
		t.Fatal("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputDate(t *testing.T) {
	for _, data := range Date().GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := Date().Validate(data); ok != nil {
				t.Fatalf("result [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputDate(t *testing.T) {
	for _, data := range Date().WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := Date().Validate(data); ok == nil {
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
