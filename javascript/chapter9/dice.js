class Dice {
	/** @param {number} sides */
	constructor(sides) {
		this.sides = sides;
	}

	rollDie() {
		console.log(Math.floor(Math.random() * this.sides + 1));
	}
}

const dice = new Dice(6);
for (let i = 1; i <= 10; i++) {
	dice.rollDie();
}
const dice10 = new Dice(10);
for (let i = 1; i <= 10; i++) {
	dice10.rollDie();
}

const dice20 = new Dice(20);
for (let i = 1; i <= 10; i++) {
	dice20.rollDie();
}
