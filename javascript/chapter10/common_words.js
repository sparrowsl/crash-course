import fs from "fs";

try {
	const fileContents = fs.readFileSync("./romeo_juliet.txt");
	const count = String(fileContents).split("Juliet").length - 1;
	console.log(count);
} catch (e) {
	console.log(e.message);
}
