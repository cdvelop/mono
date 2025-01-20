package inputs

import (
	"log"
	"testing"
)

func Test_InputIp(t *testing.T) {
	var (
		dataIp = map[string]struct {
			inputData string
			expected  string
		}{
			"IPv4 ok":        {"192.168.1.1", ""},
			"IPv6 ok":        {"2001:0db8:85a3:0000:0000:8a2e:0370:7334", ""},
			"ip incorrecta ": {"192.168.1.1.8", Lang.T(D.Format, "IPv4", D.NotValid)},
			"correcto?":      {"0.0.0.0", Lang.T(D.Example, "IP", D.NotAllowed, ':', "0.0.0.0")},
			"sin data ":      {"", Lang.T(D.Field, D.Empty, D.NotAllowed)},
		}
	)
	for prueba, data := range dataIp {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			err := Ip().Validate(data.inputData)

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

func Test_TagIp(t *testing.T) {
	tag := Ip().Render(1)
	if tag == "" {
		log.Fatalln("ERROR NO TAG RENDERING ")
	}
}

func Test_GoodInputIp(t *testing.T) {
	for _, data := range Ip().GoodTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := Ip().Validate(data); ok != nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}

func Test_WrongInputIp(t *testing.T) {
	for _, data := range Ip().WrongTestData() {
		t.Run((data), func(t *testing.T) {
			if ok := Ip().Validate(data); ok == nil {
				log.Fatalf("resultado [%v] [%v]", ok, data)
			}
		})
	}
}
