while (true) {
	const age = prompt("Enter the age: ");

	if (Number.isNaN(age)) {
		alert("Enter a valid age!!");
		continue;
	}


	if (Number(age) < 3) {
		alert("Movie ticket is free!");
	} else if (Number(age) >= 3 && Number(age) <= 12) {
		alert("Movie ticket is $10.");
	} else if (Number(age) > 12) {
		alert("Movie ticket is $15.");
	}
}
