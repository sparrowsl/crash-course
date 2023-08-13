const sandwichOrders = ["bread", "apple", "vanilla", "brown", "tuna"];
const finishedSandwiches = [];

for (const value of sandwichOrders) {
	console.log(`I made your ${value} sandwich`);
	finishedSandwiches.push(value);
}

console.log(finishedSandwiches);
