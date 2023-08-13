const people = prompt("How many people are in the group? ");

if (Number.isNaN(people)) {
	alert("Please enter a valid number!");
	window.close();
}

if (Number(people) > 8) {
	alert("Sorry, but you will have to wait for a table");
} else {
	alert("The table is ready!!");
}
