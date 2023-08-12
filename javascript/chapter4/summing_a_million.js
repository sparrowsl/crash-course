let numbers = [];

for (let i = 1; i <= 1_000_000; i++) {
	numbers.push(i);
}

console.log(min(numbers));
console.log(max(numbers));
console.log(sum(numbers));


/**
 * @param {Number[]} numbers
 * @returns {Number}
 */
function min(numbers) {
	let min = numbers[0];

	numbers.forEach(n => {
		if (n < min) min = n;
	});

	return min;
}

/**
 * @param {Number[]} numbers
 * @returns {Number}
 */
function max(numbers) {
	let max = numbers[0];

	numbers.forEach(n => {
		if (n > max) max = n;
	});

	return max;
}

/**
 * @param {Number[]} numbers
 * @returns {Number}
 */
function sum(numbers) {
	let total = 0;

	numbers.forEach(n => total += n);

	return total;
}
