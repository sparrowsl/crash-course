sandwichItems("cake", "sugar");
sandwichItems("mint", "pepper", "cream", "milk");
sandwichItems("chocolate");

/** @param {string[]} items */
function sandwichItems(...items) {
	console.log("The sandwich items:");

	for (const item of items) console.log(item);

	console.log();
}
