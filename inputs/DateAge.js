function AgeInputChange(input) {
	// console.log("AgeInputChange INPUT: ", input);
	let input_date = input.parentNode.nextElementSibling.querySelector('input[type="date"]');
	if (input.value > 0) {
		// let input_date = input.parentNode.nextElementSibling.querySelector('input[data-name="dateage"]');
		let new_birthDay = ageToBirthDay(input.value);
		inputRight(input_date);
		input_date.value = new_birthDay
		// console.log("new_birthDay: ", new_birthDay, "input_date: ", input_date);
	} else {
		input_date.defaultValue

		if (input_date.hasAttribute('required')) {
			inputWrong(input_date,"0 invalido, 0.1 = un mes");
		} else {
			inputNormal(input_date);
		}
	}
};

function DateAgeChange(input_date) {
	let input_age = input_date.parentNode.previousElementSibling.querySelector('input[type="number"]');
	let age = calculateAge(input_date.value)
	if (age >= 0) {
		input_age.value = age;
		userFormTyping(input_date);
	} else {
		input_age.value = '';
	}
}

// date format: 1980-03-01
function calculateAge(birthday) {
	const date_now = new Date();
	const current_year = parseInt(date_now.getFullYear());
	const current_month = parseInt(date_now.getMonth()) + 1;
	const today = parseInt(date_now.getDate());

	const birth_year = parseInt(String(birthday).substring(0, 4));
	const birth_month = parseInt(String(birthday).substring(5, 7));
	const birth_day = parseInt(String(birthday).substring(8, 10));

	let edad = current_year - birth_year;

	if (current_month < birth_month || (current_month === birth_month && today < birth_day)) {
		edad -= 1;
	}

	// Calcular los meses en decimal
	const monthsElapsed = (current_month - birth_month + (today >= birth_day ? 1 : 0));
	const ageWithDecimal = edad + monthsElapsed / 12; // Calcular la edad con decimal

	// Redondear a 1 decimal y convertir a cadena
	const ageString = ageWithDecimal.toFixed(1);

	return ageString;
}

// age ej 0.1, 40.0
function ageToBirthDay(decimal_age) {
	let out_date = "";
	if (!isNaN(decimal_age) && decimal_age != "" && decimal_age != "0" && decimal_age != "0.") {
		// Obtén la fecha actual
		const now_date = new Date();

		// Calcula los años y meses a partir de la edad decimal
		const years = Math.floor(decimal_age);
		const decimal_month = (decimal_age - years) * 12;
		const months = Math.floor(decimal_month);

		// Resta años y meses a la fecha actual para obtener la fecha de nacimiento
		now_date.setFullYear(now_date.getFullYear() - years);
		now_date.setMonth(now_date.getMonth() - months);

		// Formatea la fecha de nacimiento en el formato YYYY-MM-DD
		const year = now_date.getFullYear();
		const month = String(now_date.getMonth() + 1).padStart(2, '0');
		const day = String(now_date.getDate()).padStart(2, '0');

		out_date = year + "-" + month + "-" + day;
	}
	return out_date
}


