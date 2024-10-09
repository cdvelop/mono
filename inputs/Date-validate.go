package inputs

import (
	"errors"
	"strconv"
)

// verifica formato 2006-01-02 y si los rangos de el año, mes y dia son validos
// y si los Dias existen según año y mes bisiesto

func (d date) CheckDateExists(date string) error {
	const this = "CheckDateExists "
	err := d.CorrectFormatDate(date)
	if err != nil {
		return errors.New(this + err.Error())
	}

	year, month, day, er := stringToDateNumberSeparate(date)
	if er != "" {
		return errors.New(this + er)
	}

	// Verificar los rangos para año, mes y día

	if year < 1000 || year > 9999 {
		return errors.New("año fuera de rango")
	}

	if month < 1 || month > 12 {
		return errors.New("mes fuera de rango")
	}

	month_days := d.MonthDays(year)[month]
	if day < 1 {
		return errors.New("día no puede ser cero")
	}

	if day > month_days {
		errMsg := d.SpanishMonth()[month] + " no contiene " + strconv.Itoa(day) + " días."

		if !d.IsLeap(year) && month == 2 {
			errMsg += " año " + date[:4] + " no es bisiesto."
		}

		return errors.New(errMsg)
	}

	return nil
}

// verifica formato y valores numéricos en sus posiciones ej: "2006-01-02"
func (d date) CorrectFormatDate(date string) error {
	if len(date) != 10 {
		return errors.New("formato de fecha ingresado incorrecto ej: 2006-01-02")
	}

	numMap := map[byte]bool{
		'0': true, '1': true, '2': true, '3': true, '4': true,
		'5': true, '6': true, '7': true, '8': true, '9': true,
	}

	// Verificar que los guiones estén en las posiciones correctas y que los caracteres sean números
	for i, char := range date {
		if i == 4 || i == 7 {
			if char != '-' {
				return errors.New("formato de fecha ingresado incorrecto ej: 2006-01-02")
			}
		} else {
			if !numMap[byte(char)] {
				return errors.New("formato de fecha ingresado incorrecto ej: 2006-01-02")
			}
		}
	}

	return nil
}

func (d date) SpanishMonth() map[int]string {
	return map[int]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
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
func stringToDateNumberSeparate(date string) (year, month, day int, err string) {

	//YEAR
	year, e := strconv.Atoi(date[:4])
	if e != nil {
		err = e.Error()
		return
	}

	//MONTH
	monthText := date[5:7]

	if monthText >= "01" && monthText <= "09" {
		month, e = strconv.Atoi(string(monthText[1]))
	} else if monthText >= "10" && monthText <= "12" {
		month, e = strconv.Atoi(monthText)
	} else {
		err = "error " + monthText + " es un formato de mes incorrecto ej: 01 a 12"
		return
	}
	if e != nil {
		err = e.Error()
		return
	}

	//DAY
	dayTxt := date[8:10]

	if dayTxt >= "01" && dayTxt <= "09" {
		day, e = strconv.Atoi(string(dayTxt[1]))
	} else if dayTxt >= "10" && dayTxt <= "31" {
		day, e = strconv.Atoi(dayTxt)
	}

	if e != nil {
		err = e.Error()
		return
	}

	return
}
