function ageToBirthDayOLD(age_in, birthday) {
	let age = parseInt(age_in);
	let out_date = "0000-00-00";
	if (!isNaN(age)) {

		let date_now = new Date();

		// let birth_month, birth_day;
		const now_year = parseInt(date_now.getFullYear());

		if (birthday.length === 10) {

			date_now = new Date(birthday);
			date_now.setMonth(date_now.getMonth - 1);
		}

		date_now.setFullYear(now_year - age);

		out_date = date_now.toISOString().slice(0, 10);
	};
	return out_date
};

// date format: 1980-03-01
function calculateAgeOld(birthday) {
	const date_now = new Date();
	const current_year = parseInt(date_now.getFullYear());
	const current_month = parseInt(date_now.getMonth()) + 1;
	const today = parseInt(date_now.getDate());

	// 2016-07-11
	const birth_year = parseInt(String(birthday).substring(0, 4));
	const birth_month = parseInt(String(birthday).substring(5, 7));
	const birth_day = parseInt(String(birthday).substring(8, 10));

	let edad = current_year - birth_year;
	if (current_month < birth_month) {
		edad--;
	} else if (current_month === birth_month) {
		if (today < birth_day) {
			edad--;
		}
	};
	return edad;
};