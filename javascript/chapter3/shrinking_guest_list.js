let guestList = ["Melkey", "Johnny Depp", "Linus Trovalds"];

console.log(`Dear ${guestList[0]}, you are invited to my party!!`);
console.log(`Dear ${guestList[1]}, you are invited to my party!!`);
console.log(`Dear ${guestList[2]}, you are invited to my party!!`);

console.log(`\nUnfortunately ${guestList[2]} can't make it.\n`);

guestList = ["Melkey", "Johnny Depp", "Rich Harris"];
console.log(`Dear ${guestList[0]}, you are invited to my party!!`);
console.log(`Dear ${guestList[1]}, you are invited to my party!!`);
console.log(`Dear ${guestList[2]}, you are invited to my party!!`);

console.log("\nI found a bigger table for the guests!!\n");
guestList.push("The Primegean");
guestList.push("Wagslane");

console.log("Sorry guys, I can only invite 2 guests!!\n");
guestList.pop();
guestList.pop();
guestList.pop();
console.log(guestList);
