const number = prompt("Enter a number: ");

// Trim all spaces of the input and convert to number
if (Number.isNaN(number?.trim())) {
	alert("Invalid number");
	window.close();
}

if (Number(number?.trim()) % 10 !== 0) {
	alert(`${number} is not a multiple of 10!`);
} else {
	alert(`${number} is a multiple of 10!`);
}
