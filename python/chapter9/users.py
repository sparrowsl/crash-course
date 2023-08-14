class User:
    def __init__(self, first_name: str, last_name: str, age: int):
        self.first_name = first_name
        self.last_name = last_name
        self.age = age

    def describe_user(self):
        print(f"Name: {self.first_name} {self.last_name}")
        print(f"Age: {self.age}")

    def greet_user(self):
        print(f"Hello there {self.first_name} {self.last_name}")


john = User("John", "Smith", 10)
jack = User("Jack", "Willborne", 20)
jenny = User("Jenny", "Trush", 15)

john.describe_user()
john.greet_user()

jack.describe_user()
jack.greet_user()

jenny.describe_user()
jenny.greet_user()
