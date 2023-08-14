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

const restaurant = new Restaurant("Sunny", "Classic");

console.log(restaurant.restaurantName);
console.log(restaurant.cuisineType);

restaurant.openRestaurant();
restaurant.describeRestaurant();
