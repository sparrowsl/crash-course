class Restaurant:
    def __init__(self, restaurant_name: str, cuisine_type: str):
        self.restaurant_name = restaurant_name
        self.cuisine_type = cuisine_type
        self.number_served = 0

    def describe_restaurant(self):
        print(f"Restaurant Name: {self.restaurant_name}")
        print(f"Cuisine Type: {self.cuisine_type}")

    def open_restaurant(self):
        print(f"The {self.restaurant_name} Restaurant is now open!")

    def set_number_served(self, num: int):
        self.number_served = num

    def increment_number_served(self, num: int):
        self.number_served += num


restaurant = Restaurant("Sunny", "Classic")

print(restaurant.restaurant_name)
print(restaurant.cuisine_type)

restaurant.open_restaurant()
restaurant.describe_restaurant()


print(
    f"The {restaurant.restaurant_name} has served {restaurant.number_served} customers."
)
restaurant.number_served = 3
print(
    f"The {restaurant.restaurant_name} has served {restaurant.number_served} customers."
)

restaurant.set_number_served(10)
print(
    f"The {restaurant.restaurant_name} has served {restaurant.number_served} customers."
)
restaurant.increment_number_served(5)
print(
    f"The {restaurant.restaurant_name} has served {restaurant.number_served} customers."
)
