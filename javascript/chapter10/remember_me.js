import fs from "fs";
import * as readline from "node:readline/promises";
import { exit, stdin, stdout } from "process";


displayUserInfo();
saveUserInfo();

async function saveUserInfo() {
	let prompt = readline.createInterface({ input: stdin, output: stdout });
	const name = await prompt.question("Enter name: ");
	prompt.close();

	prompt = readline.createInterface({ input: stdin, output: stdout });
	const age = await prompt.question("Enter age: ");
	prompt.close();

	if (Number.isNaN(Number(age))) {
		console.log("Please enter a valid age!!");
		exit();
	}

	prompt = readline.createInterface({ input: stdin, output: stdout });
	const contact = await prompt.question("Enter contact: ");
	prompt.close();

	const userDetails = JSON.stringify({ name, age: Number(age), contact });
	fs.writeFileSync("./remember.json", userDetails);
}

function displayUserInfo() {
	try {
		const contents = JSON.parse(String(fs.readFileSync("./remember.json")));
		console.log(contents);
	} catch (e) {
		console.log(e.message);
	}
}
