class Restaurant:
    def __init__(self, name: str, cuisine_type: str, number_served: int):
        self.name = name
        self.cuisine_type = cuisine_type
        self.number_served = number_served

    def describeRestaurant(self):
        print(f"Name: {self.name}")
        print(f"Cuisine Type: {self.cuisine_type}")

    def openRestaurant(self):
        print(f"The {self.name} Restaurant is now open!")

    def setNumberServed(self, num: int):
        self.number_served = num

    def incrementNumberServed(self, num: int):
        self.number_served += num


class IceCreamStand(Restaurant):
    def __init__(self, flavors: list[str]):
        self.flavors = flavors

    def show_flavors(self):
        for flavor in self.flavors:
            print(flavor)


ics = IceCreamStand(["vanilla", "chocolate", "cream"])
ics.show_flavors()
