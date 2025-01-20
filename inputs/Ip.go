package inputs

import (
	"strings"
)

// valid ip address with dot-separated fields
func Ip(params ...any) *ip {
	new := &ip{
		input: input{
			attributes: attributes{
				htmlName:   "text",
				customName: "ip",
				Title:      `title="` + Lang.T(D.Example, ':') + ` 192.168.0.8"`,
				// Pattern: `^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`,
			},
			permitted: permitted{
				Letters:         true,
				Numbers:         true,
				Characters:      []rune{'.', ':'},
				Minimum:         7,  //IPv4 - IPv6 es 39
				Maximum:         39, // IPv6 - IPv4 es 15
				ExtraValidation: &ip{},
			},
		},
	}
	new.Set(params)

	return new
}

type ip struct {
	input
}

func (i ip) ExtraValidation(value string) error {

	var ipV string

	if value == "0.0.0.0" {
		return Lang.Err(D.Example, "IP", D.NotAllowed, ':', "0.0.0.0")
	}

	if strings.Contains(value, ":") { //IPv6
		ipV = ":"
	} else if strings.Contains(value, ".") { //IPv4
		ipV = "."
	}

	part := strings.Split(value, ipV)

	if ipV == "." && len(part) != 4 {
		return Lang.Err(D.Format, "IPv4", D.NotValid)
	}

	if ipV == ":" && len(part) != 8 {
		return Lang.Err(D.Format, "IPv6", D.NotValid)
	}

	return nil

}

func (i ip) GoodTestData() (out []string) {

	out = []string{
		"120.1.3.206",
		"195.145.149.184",
		"179.183.230.16",
		"253.70.9.26",
		"215.35.117.51",
		"212.149.243.253",
		"126.158.214.250",
		"49.122.253.195",
		"53.218.195.25",
		"190.116.115.117",
		"115.186.149.240",
		"163.95.226.221",
	}

	return
}

func (i ip) WrongTestData() (out []string) {
	out = []string{
		"0.0.0.0",
		"192.168.1.1.8",
	}
	out = append(out, wrong_data...)
	return
}
