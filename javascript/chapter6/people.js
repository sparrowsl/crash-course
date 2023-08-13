const people = {
	"Jack": {
		"firstName": "Jack",
		"lastName": "Doe",
		"age": 10,
		"city": "Freetown",
	},
	"Kate": {
		"firstName": "Kathrine",
		"lastName": "Smith",
		"age": 13,
		"city": "Detroit",
	},
};

for (const [key, value] of Object.entries(people)) {
	console.log( key);

	for (const [k, v] of Object.entries(value)) {
		console.log(`${k}: ${v}`);
	}

	console.log();
}
