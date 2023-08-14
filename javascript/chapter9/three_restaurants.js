class Restaurant {
	/**
	 * @param {string} restaurantName
	 * @param {string} cuisineType
	 */
	constructor(restaurantName, cuisineType) {
		this.restaurantName = restaurantName;
		this.cuisineType = cuisineType;
	}

	describeRestaurant() {
		console.log(`Restaurant Name: ${this.restaurantName}`);
		console.log(`Cuisine Type: ${this.cuisineType}`);
	}

	openRestaurant() {
		console.log(`The ${this.restaurantName} Restaurant is now open!`);
	}
}

const res1 = new Restaurant("Sunny", "Classic");
res1.openRestaurant();
res1.describeRestaurant();

const res2 = new Restaurant("Nineties", "VIP");
res2.openRestaurant();
res2.describeRestaurant();

const res3 = new Restaurant("Tech", "Nerds");
res3.openRestaurant();
res3.describeRestaurant()

