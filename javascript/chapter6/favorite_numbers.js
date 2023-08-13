const people = { "Jack": 20, "Jane": 4, "Joe": 10, "Jill": 7, "Jenny": 2 };

Object.entries(people).forEach(([k, v]) => console.log(`${k}'s favorite number is ${v}`));
