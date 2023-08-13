while (true) {
	const topping = ("Enter a pizza topping: ");

	if (topping.toLowerCase() === "quit") window.close();

	alert(`I will add the ${topping} on your pizza.`);
}
