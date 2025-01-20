package inputs

import (
	"errors"
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
		return Lang.Err(D.MonthOutOfRange)
	}

	month_days := d.MonthDays(year)[month]
	if day < 1 {
		return Lang.Err(D.DayCannotBeZero)
	}

	if day > month_days {
		errMsg := d.NameMonths()[month] + " " + Lang.T(D.DoesNotExist) + " " + strconv.Itoa(day) + " " + Lang.T(D.Day) + "."

		if !d.IsLeap(year) && month == 2 {
			errMsg += " " + Lang.T(D.Year) + " " + date[:4] + " " + Lang.T(D.NotValid) + "."
		}

		return errors.New(errMsg)
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

// formato fecha "2006-01-02" retorna: 2006,1,2. NOTA: NO VERIFICA EL FORMATO INGRESADO
func stringToDateNumberSeparate(date string) (year, month, day int, err error) {

	//YEAR
	year, err = strconv.Atoi(date[:4])
	if err != nil {
		return
	}

	//MONTH
	monthText := date[5:7]

	if monthText >= "01" && monthText <= "09" {
		month, err = strconv.Atoi(string(monthText[1]))
	} else if monthText >= "10" && monthText <= "12" {
		month, err = strconv.Atoi(monthText)
	} else {
		err = Lang.Err(monthText, D.Format, D.Month, D.NotValid, D.Example, ':', "01 a 12")
		return
	}
	if err != nil {
		return
	}

	//DAY
	dayTxt := date[8:10]

	if dayTxt >= "01" && dayTxt <= "09" {
		day, err = strconv.Atoi(string(dayTxt[1]))
	} else if dayTxt >= "10" && dayTxt <= "31" {
		day, err = strconv.Atoi(dayTxt)
	}

	if err != nil {
		return
	}

	return
}
