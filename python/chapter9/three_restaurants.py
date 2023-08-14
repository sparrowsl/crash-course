class Restaurant:
    def __init__(self, restaurant_name: str, cuisine_type: str):
        self.restaurant_name = restaurant_name
        self.cuisine_type = cuisine_type

    def describe_restaurant(self):
        print(f"Restaurant Name: {self.restaurant_name}")
        print(f"Cuisine Type: {self.cuisine_type}")

    def open_restaurant(self):
        print(f"The {self.restaurant_name} Restaurant is now open!")


res1 = Restaurant("Sunny", "Classic")
res1.open_restaurant()
res1.describe_restaurant()

res2 = Restaurant("Nineties", "VIP")
res2.open_restaurant()
res2.describe_restaurant()

res3 = Restaurant("Tech", "Nerds")
res3.open_restaurant()
res3.describe_restaurant()
