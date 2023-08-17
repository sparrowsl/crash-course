import * as readline from "readline/promises";

const prompt = readline.createInterface({ input: process.stdin, output: process.stdout });

const input = await prompt.question("Enter 2 numbers: ");
prompt.close();
const numbers = input.split(" ");

if (Number.isNaN(Number(numbers[0])) || Number.isNaN(Number(numbers[1]))) {
	console.log("Please enter two valid numbers!!");
	process.exit(1);
}

console.log(Number(numbers[0]) + Number(numbers[1]));
