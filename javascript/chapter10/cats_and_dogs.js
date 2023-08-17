import fs from "fs";

try {
	const file = fs.readFileSync("./cats.txt");
	console.log(String(file));
} catch (e) {
	console.log("File does not exist!!");
	process.exit();
}

try {
	const file = fs.readFileSync("./dogs.txt");
	console.log(String(file));
} catch (e) {
	console.log("File does not exist!!");
	process.exit();
}
