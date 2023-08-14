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
		console.log("Name: %s\n", this.name);
		console.log("Cuisine Type: %s\n", this.cuisineType);
	}

	openRestaurant() {
		console.log("The %s Restaurant is now open!\n", this.name);
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

class IceCreamStand extends Restaurant {
	/** @param {string[]} flavors */
	constructor(flavors) {
		super("", "");
		this.flavors = flavors;
	}

	showFlavors() {
		for (const flavor of this.flavors) console.log(flavor);
	}
}

const ics = new IceCreamStand(["vanilla", "chocolate", "cream"]);
ics.showFlavors();
