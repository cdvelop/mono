package mono

import (
	"reflect"
	"strconv"
	"strings"
	"unicode"
	"unsafe"
)

// G global functions variable
var G = global{
	String: stringUtils{},
	Number: numberUtils{},
	Date:   dateUtils{},
	Rut:    rutUtils{},
}

type global struct {
	String stringUtils
	Number numberUtils
	Date   dateUtils
	Rut    rutUtils
}

// NUMERIC UTILS
type numberUtils struct{}

// returns the number as a string and its size, which will never be more than 19 characters (int64)
func (n numberUtils) IsNumericValue(refValue *reflect.Value) (numStr string, size uint8, ok bool) {

	switch refValue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		numStr = strconv.FormatInt(refValue.Int(), 10)
		ok = true

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		numStr = strconv.FormatUint(refValue.Uint(), 10)
		ok = true

	case reflect.Float32:
		buf := make([]byte, 0, 10) // allocate a 32-byte buffer to avoid reallocation
		buf = strconv.AppendFloat(buf, refValue.Float(), 'f', -1, 32)
		numStr = unsafe.String(unsafe.SliceData(buf), len(buf))
		ok = true
	case reflect.Float64:
		buf := make([]byte, 0, 19) // allocate a 64-byte buffer to avoid reallocation
		buf = strconv.AppendFloat(buf, refValue.Float(), 'f', -1, 64)
		numStr = unsafe.String(unsafe.SliceData(buf), len(buf))
		ok = true

	}

	return numStr, uint8(len(numStr)), ok
}

// STRING UTILS
type stringUtils struct{}

// snakeCase converts a string to snake_case format with optional separator.
// If no separator is provided, underscore "_" is used as default.
// Example:
//
//	Input: "camelCase" -> Output: "camel_case"
//	Input: "PascalCase", "-" -> Output: "pascal-case"
//	Input: "APIResponse" -> Output: "api_response"
//	Input: "user123Name", "." -> Output: "user123.name"
func (s stringUtils) SnakeCase(str string, sep ...string) string {
	separator := "_"
	if len(sep) > 0 {
		separator = sep[0]
	}
	var out string
	for i, r := range str {
		if unicode.IsUpper(r) {
			// If it's uppercase and not the first character, add separator
			if i > 0 && (unicode.IsLower(rune(str[i-1])) || unicode.IsDigit(rune(str[i-1]))) {
				out += separator
			}
			// Convert uppercase to lowercase
			out += strings.ToLower(string(r))
		} else {
			// If it's not uppercase, simply add it
			out += string(r)
		}
	}
	return out
}

// DATE UTILS
type dateUtils struct{}

// verifica formato 2006-01-02 y si los rangos de el año, mes y dia son validos
// y si los Dias existen según año y mes bisiesto
func (d dateUtils) DateExists(date string) error {
	err := d.CheckFormatDate(date)
	if err != nil {
		return err
	}

	year, month, day, err := d.StringToDateNumberSeparate(date)
	if err != nil {
		return err
	}

	// Verificar los rangos para año, mes y día
	if year < 1000 || year > 9999 {
		return R.Err(D.YearOutOfRange)
	}

	if month < 1 || month > 12 {
		return R.Err(D.Month, strconv.Itoa(month), D.NotValid)
	}

	// Validación específica para febrero
	if month == 2 {
		febDays := 28
		if d.IsLeap(year) {
			febDays = 29
		}

		if day < 1 {
			return R.Err(D.DayCannotBeZero)
		}
		if day > febDays {
			yearMsg := ""
			if !d.IsLeap(year) {
				yearMsg = R.T(D.Year, strconv.Itoa(year))
			}
			return R.Err(d.MonthNames()[month], D.DoesNotHave, strconv.Itoa(day), D.Days, yearMsg)
		}
		return nil
	}

	// Validación para otros meses
	month_days := d.MonthDays(year)[month]
	if day < 1 {
		return R.Err(D.DayCannotBeZero)
	}
	if day > month_days {
		return R.Err(d.MonthNames()[month], D.DoesNotHave, strconv.Itoa(day), D.Days)
	}

	return nil
}

// verifica formato y valores numéricos en sus posiciones ej: "2006-01-02"
func (d dateUtils) CheckFormatDate(date string) error {
	if len(date) != 10 {
		return R.Err(D.InvalidDateFormat, "2006-01-02")
	}

	numMap := map[byte]bool{
		'0': true, '1': true, '2': true, '3': true, '4': true,
		'5': true, '6': true, '7': true, '8': true, '9': true,
	}

	// Verificar que los guiones estén en las posiciones correctas y que los caracteres sean números
	for i, char := range date {
		if i == 4 || i == 7 {
			if char != '-' {
				return R.Err(D.InvalidDateFormat, "2006-01-02")
			}
		} else {
			if !numMap[byte(char)] {
				return R.Err(D.InvalidDateFormat, "2006-01-02")
			}
		}
	}

	return nil
}

func (d dateUtils) MonthNames() map[int]string {
	return map[int]string{
		1:  R.T(D.January),
		2:  R.T(D.February),
		3:  R.T(D.March),
		4:  R.T(D.April),
		5:  R.T(D.May),
		6:  R.T(D.June),
		7:  R.T(D.July),
		8:  R.T(D.August),
		9:  R.T(D.September),
		10: R.T(D.October),
		11: R.T(D.November),
		12: R.T(D.December),
	}
}

func (d dateUtils) MonthDays(year int) map[int]int {
	var feb_days = 28

	if d.IsLeap(year) {
		feb_days = 29
	}

	return map[int]int{
		1:  31,
		2:  feb_days,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
}

// es bisiesto este año?
func (d dateUtils) IsLeap(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}

// date format "2006-01-02" returns: 2006,1,2.
// NOTE: DOES NOT VERIFY THE INPUT FORMAT
func (d dateUtils) StringToDateNumberSeparate(date string) (year, month, day int, err error) {
	//YEAR
	year, err = strconv.Atoi(date[:4])
	if err != nil {
		return
	}

	// MONTH
	monthText := date[5:7]

	if monthText == "00" {
		err = R.Err(D.InvalidDateFormat)
		return
	}

	month, err = strconv.Atoi(monthText)
	if err != nil || month < 1 || month > 12 {
		err = R.Err(D.Month, monthText, D.NotValid)
		return
	}

	//DAY
	day, err = d.ValidateDay(date[8:10])
	if err != nil {
		return
	}

	return
}

// validateDay converts a day string to an integer
// Parameters:
//   - dayTxt: string representing a day in format "01" to "31"
//
// Returns:
//   - day: integer value of the day
//   - err: error if the day cannot be converted
func (d dateUtils) ValidateDay(dayTxt string) (day int, err error) {
	if len(dayTxt) > 2 {
		return 0, R.Err(D.MaxSize, "2", D.Chars)
	}

	// Verificar que todos los caracteres sean dígitos
	for _, c := range dayTxt {
		if c < '0' || c > '9' {
			return 0, R.Err(D.NotNumber)
		}
	}

	day, err = strconv.Atoi(dayTxt)
	if err != nil {
		return 0, R.Err(D.Field, D.Day, D.NotValid)
	}

	return day, nil
}

// RUT UTILS
type rutUtils struct{}

// DvRut retorna dígito verificador de un run
func (r rutUtils) DvRut(rut int) string {
	var sum = 0
	var factor = 2
	for ; rut != 0; rut /= 10 {
		sum += rut % 10 * factor
		if factor == 7 {
			factor = 2
		} else {
			factor++
		}
	}

	if val := 11 - (sum % 11); val == 11 {
		return "0"
	} else if val == 10 {
		return "k"
	} else {
		return strconv.Itoa(val)
	}
}

func runData(runIn string) (data []string, onlyRun int, err error) {

	if len(runIn) < 3 {
		return nil, 0, R.Err(D.Value, D.Empty)
	}

	// Separar número y dígito verificador
	data = strings.Split(runIn, "-")
	if len(data) != 2 {
		return nil, 0, R.Err(D.Format, D.NotValid)
	}

	// Validar caracteres del número
	if !isDigits(data[0]) {
		return nil, 0, R.Err(D.Chars, D.NotAllowed, D.In, D.Numbers)
	}

	// Validar dígito verificador
	dv := strings.ToLower(data[1])
	if len(dv) != 1 || (dv != "k" && !isDigits(dv)) {
		return nil, 0, R.Err(D.Digit, D.Verifier, dv, D.NotValid)
	}

	// Convertir número a entero
	onlyRun, err = strconv.Atoi(data[0])
	if err != nil {
		return nil, 0, R.Err(D.Numbers, D.NotValid)
	}

	return data, onlyRun, nil
}

func isDigits(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
