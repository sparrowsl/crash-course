let guest_list = ["Melkey", "Johnny Depp", "Linus Trovalds"];

console.log(`Dear ${guest_list[0]}, you are invited to my party!!`);
console.log(`Dear ${guest_list[1]}, you are invited to my party!!`);
console.log(`Dear ${guest_list[2]}, you are invited to my party!!`);

console.log(`\nUnfortunately ${guest_list[2]} can't make it.\n`);

guest_list = ["Melkey", "Johnny Depp", "Rich Harris"];
console.log(`Dear ${guest_list[0]}, you are invited to my party!!`);
console.log(`Dear ${guest_list[1]}, you are invited to my party!!`);
console.log(`Dear ${guest_list[2]}, you are invited to my party!!`);

console.log(`\nI found a bigger table for the guests!!\n`);
guest_list.push("The Primegean");
guest_list.push("Wagslane");

console.log(guest_list);
