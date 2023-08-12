let threes = [];

for (let i = 3; i <= 30; i++) {
	if (i % 3 === 0) threes.push(i);
}

threes.forEach(t => console.log(t));
