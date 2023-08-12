const currentUsers = ["john", "jane", "adam", "jack", "jenny", "barry", "joe"];
const newUsers = ["jonas", "jane", "Jenny", "angela", "jack"];

for (let user of newUsers) {
	if (currentUsers.includes(user.toLowerCase())) {
		console.log("Username already taken, choose another name.");
	} else {
		console.log("Username is available!");
	}
}
