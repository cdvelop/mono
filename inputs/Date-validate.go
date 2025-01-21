package inputs

import (
	"strconv"
)

// verifica formato 2006-01-02 y si los rangos de el año, mes y dia son validos
// y si los Dias existen según año y mes bisiesto

func (d date) CheckDateExists(date string) error {
	err := d.CorrectFormatDate(date)
	if err != nil {
		return err
	}

	year, month, day, err := stringToDateNumberSeparate(date)
	if err != nil {
		return err
	}

	// Verificar los rangos para año, mes y día

	if year < 1000 || year > 9999 {
		return Lang.Err(D.YearOutOfRange)
	}

	if month < 1 || month > 12 {
		return Lang.Err(D.Month, strconv.Itoa(month), D.NotValid)
	}

	month_days := d.MonthDays(year)[month]
	if day < 1 {
		return Lang.Err(D.DayCannotBeZero)
	}

	if day > month_days {
		var yearMsg string
		// Solo agregar el año para febrero en años no bisiestos
		if month == 2 && !d.IsLeap(year) {
			yearMsg = Lang.T(D.Year, date[:4])
		}

		return Lang.Err(d.NameMonths()[month], D.DoesNotHave, strconv.Itoa(day), D.Days, yearMsg)
	}

	return nil
}

// verifica formato y valores numéricos en sus posiciones ej: "2006-01-02"
func (d date) CorrectFormatDate(date string) error {
	if len(date) != 10 {
		return Lang.Err(D.InvalidDateFormat, "2006-01-02")
	}

	numMap := map[byte]bool{
		'0': true, '1': true, '2': true, '3': true, '4': true,
		'5': true, '6': true, '7': true, '8': true, '9': true,
	}

	// Verificar que los guiones estén en las posiciones correctas y que los caracteres sean números
	for i, char := range date {
		if i == 4 || i == 7 {
			if char != '-' {
				return Lang.Err(D.InvalidDateFormat, "2006-01-02")
			}
		} else {
			if !numMap[byte(char)] {
				return Lang.Err(D.InvalidDateFormat, "2006-01-02")
			}
		}
	}

	return nil
}

func (d date) NameMonths() map[int]string {
	return map[int]string{
		1:  Lang.T(D.January),
		2:  Lang.T(D.February),
		3:  Lang.T(D.March),
		4:  Lang.T(D.April),
		5:  Lang.T(D.May),
		6:  Lang.T(D.June),
		7:  Lang.T(D.July),
		8:  Lang.T(D.August),
		9:  Lang.T(D.September),
		10: Lang.T(D.October),
		11: Lang.T(D.November),
		12: Lang.T(D.December),
	}
}

func (d date) MonthDays(year int) map[int]int {

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
func (d date) IsLeap(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}

// date format "2006-01-02" returns: 2006,1,2.
// NOTE: DOES NOT VERIFY THE INPUT FORMAT
func stringToDateNumberSeparate(date string) (year, month, day int, err error) {

	//YEAR
	year, err = strconv.Atoi(date[:4])
	if err != nil {
		return
	}

	// MONTH
	monthText := date[5:7]

	if monthText == "00" {
		err = Lang.Err(D.InvalidDateFormat)
		return
	}

	month, err = strconv.Atoi(monthText)
	if err != nil || month < 1 || month > 12 {
		err = Lang.Err(D.Month, monthText, D.NotValid)
		return
	}

	//DAY
	day, err = validateDay(date[8:10])
	if err != nil {
		return
	}

	return
}

// validateDay validates and converts a day string to an integer
// Parameters:
//   - dayTxt: string representing a day in format "01" to "31"
//
// Returns:
//   - day: integer value of the day
//   - err: error if the day is not valid or cannot be converted
func validateDay(dayTxt string) (day int, err error) {

	errOut := Lang.Err(D.Field, D.Day, D.NotValid)

	if len(dayTxt) > 2 {
		err = Lang.Err(D.MaxSize, "2", D.Chars)
		return
	}

	// Verificar que todos los caracteres sean dígitos
	for _, c := range dayTxt {
		if c < '0' || c > '9' {
			err = Lang.Err(D.NotNumber)
			return
		}
	}

	if dayTxt >= "01" && dayTxt <= "09" {
		day, err = strconv.Atoi(string(dayTxt[1]))
	} else if dayTxt >= "10" && dayTxt <= "31" {
		day, err = strconv.Atoi(dayTxt)
	} else {
		err = errOut
		return
	}

	if err != nil {
		err = errOut
		return
	}

	return
}
