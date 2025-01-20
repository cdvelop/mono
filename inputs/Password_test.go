package inputs

import (
	"testing"
)

func Test_InputPassword(t *testing.T) {
	var (
		modelPassword = Password()

		dataPassword = map[string]struct {
			inputData string

			expected string
		}{
			"validates numbers letters and character": {"c0ntra3", ""},
			"validates many characters":               {"M1 contraseÑ4", ""},
			"validates 8 characters":                  {"contrase", ""},
			"validates 5 characters":                  {"contñ", ""},
			"validates only numbers":                  {"12345", ""},
			"does not validate less than 2":           {"1", Lang.T(D.MinSize, 5, D.Chars)},
			"no data":                                 {"", Lang.T(D.Field, D.Empty, D.NotAllowed)},
		}
	)
	for prueba, data := range dataPassword {
		t.Run((prueba + ": " + data.inputData), func(t *testing.T) {
			err := modelPassword.Validate(data.inputData)

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

func Test_TagPassword(t *testing.T) {
	tag := Password().Render(1)
	if tag == "" {
		t.Fatal("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputPassword(t *testing.T) {
	for _, data := range Password().GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := Password().Validate(data); ok != nil {
				t.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputPassword(t *testing.T) {
	for _, data := range Password().WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := Password().Validate(data); ok == nil {
				t.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

var modelPasswordMinimal = Password(`min="10"`, `max="30"`)

func Test_GoodInputPasswordMinimal(t *testing.T) {
	for _, data := range modelPasswordMinimal.GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelPasswordMinimal.Validate(data); ok != nil {
				t.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputPasswordMinimal(t *testing.T) {
	for _, data := range modelPasswordMinimal.WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := modelPasswordMinimal.Validate(data); ok == nil {
				t.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
