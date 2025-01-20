package inputs

import (
	"log"
	"testing"
)

func Test_InputNumber(t *testing.T) {
	var (
		modelNumber = Number()

		dataNumber = map[string]struct {
			inputData string

			expected string
		}{
			"correct number 100": {"100", ""},
			"single digit 0":     {"0", ""},
			"single digit 1":     {"1", ""},
			"uint64 20 chars":    {"18446744073709551615", ""},
			"uint64 21 chars":    {"184467440737095516150", Lang.T(D.MaxSize, 20, D.Chars)},
			"int64 19 chars":     {"9223372036854775807", ""},
			"int32 10 chars":     {"2147483647", ""},
			"18 digits":          {"100002323262637278", ""},

			"large number with letter": {"10000232E26263727", Lang.T('E', D.NotNumber)},
			"negative number -100":     {"-100", Lang.T('-', D.NotNumber)},
			"text instead of number":   {"h", Lang.T('h', D.NotNumber)},
			"text and number":          {"h1", Lang.T('h', D.NotNumber)},
		}
	)
	for prueba, data := range dataNumber {
		t.Run((prueba), func(t *testing.T) {
			err := modelNumber.Validate(data.inputData)

			var err_str string
			if err != nil {
				err_str = err.Error()
			}

			if err_str != data.expected {
				log.Println(prueba)
				log.Fatalf("resultado: [%v] expectativa: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}

func Test_TagNumber(t *testing.T) {
	tag := Number().Render(1)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_InputPhoneNumber(t *testing.T) {
	var (
		// 1 408 XXX XXXX
		// 5 699 524 9966

		dataPhoneNumber = map[string]struct {
			inputData string

			expected string
		}{
			"correct 7 digit number":  {"1234567", ""},
			"correct 9 digit number":  {"123456789", ""},
			"correct 11 digit number": {"12345678911", ""},
			"with country code":       {"56988765432", ""},
			"plus sign +":             {"+56988765432", Lang.T(D.MaxSize, 11, D.Chars)},
			"6 digit number":          {"123456", Lang.T(D.MinSize, 7, D.Chars)},
			"1 digit number":          {"0", Lang.T(D.MinSize, 7, D.Chars)},
		}
	)
	for prueba, data := range dataPhoneNumber {
		t.Run((prueba), func(t *testing.T) {
			err := Phone().Validate(data.inputData)

			var err_str string
			if err != nil {
				err_str = err.Error()
			}

			if err_str != data.expected {
				log.Println(prueba)
				log.Fatalf("resultado: [%v] expectativa: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}

func Test_GoodInputPhoneNumber(t *testing.T) {
	for _, data := range Phone().GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := Phone().Validate(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_GoodInputNumber(t *testing.T) {
	for _, data := range Number().GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := Number().Validate(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputNumber(t *testing.T) {
	for _, data := range Number().WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := Number().Validate(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
