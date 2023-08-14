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
		this.loginAttempts = 0;
	}

	describeUser() {
		console.log(`Name: ${this.firstName} ${this.lastName}`);
		console.log(`Age: {this.age}`);
	}

	incrementLoginAttempts() {
		this.loginAttempts += 1;
	}

	resetLoginAttempts() {
		this.loginAttempts = 0;
	}

	greetUser() {
		console.log(`Hello there ${this.firstName}, ${this.lastName}`);
	}
}

const john = new User("John", "Smith", 10);

john.incrementLoginAttempts();
john.incrementLoginAttempts();
john.incrementLoginAttempts();
john.incrementLoginAttempts();
console.log(john.loginAttempts);

john.resetLoginAttempts();
console.log(john.loginAttempts);
