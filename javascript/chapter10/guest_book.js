import fs from "fs";
import * as readline from "readline/promises";

while (true) {
	const prompt = readline.createInterface({ input: process.stdin, output: process.stdout });
	const name = await prompt.question("Enter name: ");
	prompt.close();

	if (name.toLowerCase() == "quit") break;

	fs.appendFileSync("./guest_book.txt", name + "\n");
}
