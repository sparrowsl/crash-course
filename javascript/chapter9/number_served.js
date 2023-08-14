class Restaurant {
	/**
	 * @param {string} name
	 * @param {string} cuisineType
	 */
	constructor(name, cuisineType) {
		this.name = name;
		this.cuisineType = cuisineType;
		this.numberServed = 0;
	}


	describeRestaurant() {
		console.log(`Name: ${this.name}`);
		console.log(`Cuisine Type: ${this.cuisineType}`);
	}

	openRestaurant() {
		console.log(`The ${this.name} Restaurant is now open!`);
	}
	/** @param {number} num */
	setNumberServed(num) {
		this.numberServed = num;
	}

	/** @param {number} num */
	incrementNumberServed(num) {
		this.numberServed += num;
	}
}

const restaurant = new Restaurant("Sunny", "Classic");
console.log(`The ${restaurant.name} has served ${restaurant.numberServed} customers.`);

restaurant.numberServed = 3;
console.log(`The ${restaurant.name} has served ${restaurant.numberServed} customers.`);

restaurant.setNumberServed(10);
console.log(`The ${restaurant.name} has served ${restaurant.numberServed} customers.`);

restaurant.incrementNumberServed(5);
console.log(`The ${restaurant.name} has served ${restaurant.numberServed} customers.`);
