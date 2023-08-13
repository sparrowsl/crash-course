const pets = {
	"sparrow": {
		"animal": "Bird",
		"owner": "Sparrow",
	},
	"bolt": {
		"animal": "Dog",
		"owner": "Sparrow",
	},
};

for (const [key, value] of Object.entries(pets)) {
	console.log(key);

	for (const [k, v] of Object.entries(value)) {
		console.log(`${k}: ${v}`);
	}

	console.log();
}
