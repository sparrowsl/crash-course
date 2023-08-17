import fs from "fs";
import * as readline from "node:readline/promises";

// @ts-ignore
const prompt = readline.createInterface({ input: process.stdin, output: process.stdout });

const name = await prompt.question("Enter name: ");
prompt.close();

fs.writeFileSync("./guest.txt", name);

console.log(name);
