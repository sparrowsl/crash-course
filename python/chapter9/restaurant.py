class Restaurant:
    def __init__(self, restaurant_name: str, cuisine_type: str):
        self.restaurant_name = restaurant_name
        self.cuisine_type = cuisine_type

    def describe_restaurant(self):
        print(f"Restaurant Name: {self.restaurant_name}")
        print(f"Cuisine Type: {self.cuisine_type}")

    def open_restaurant(self):
        print(f"The {self.restaurant_name} Restaurant is now open!")


restaurant = Restaurant("Sunny", "Classic")

print(restaurant.restaurant_name)
print(restaurant.cuisine_type)

restaurant.open_restaurant()
restaurant.describe_restaurant()
