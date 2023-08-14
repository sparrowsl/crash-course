class User {
	/**
	 * @param {string} firstName
	 * @param {string} lastName
	 * @param {number} age
	 */
	constructor(firstName, lastName, age) {
		this.firstName = firstName;
		this.lastName = lastName;
		this.age = age;
	}

	describeUser() {
		console.log(`Name: ${this.firstName} ${this.lastName}`);
		console.log(`Age: ${this.age}`);
	}

	greetUser() {
		console.log(`Hello there ${this.firstName}, ${this.lastName}`);
	}
}

const john = new User("John", "Smith", 10);
const jack = new User("Jack", "Willborne", 20);
const jenny = new User("Jenny", "Trush", 15);

john.describeUser();
john.greetUser();

jack.describeUser();
jack.greetUser();

jenny.describeUser();
jenny.greetUser();
