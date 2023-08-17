import fs from "fs";
import * as readline from "readline/promises";

const prompt = readline.createInterface({ input: process.stdin, output: process.stdout });
const filename = "./fav_number.json";

const number = await prompt.question("Enter your favorite number: ");
prompt.close();

if (Number.isNaN(Number(number))) {
	console.log("That is not a valid number!!");
	process.exit();
}

saveNumber(Number(number));
readSavedNumber(filename);


/** @param {string|fs.PathOrFileDescriptor} filename */
function readSavedNumber(filename) {
	try {
		const contents = JSON.parse(String(fs.readFileSync(filename)));
		console.log(`I know your favorite number!!. It's ${contents?.number}`);
	} catch (e) {
		console.log(e.message);
	}
}

/** @param {number} number */
function saveNumber(number) {
	try {
		fs.writeFileSync(filename, JSON.stringify({ number }));
	} catch (e) {
		console.log(e.message);
	}
}
