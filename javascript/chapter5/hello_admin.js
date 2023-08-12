let names = ["john", "jane", "admin", "jack", "jenny"];
names = [];

if (names.length !== 0) {
	for (const name of names) {
		if (name === "admin") {
			console.log("Hello admin, would you like to see a status report?");
		} else {
			console.log(`Hello ${name}, thank you for logging in again.`);
		}
	}
} else {
	console.log("We need to find some users!!");
}

