class User:
    def __init__(self, first_name: str, last_name: str, age: int):
        self.first_name = first_name
        self.last_name = last_name
        self.age = age
        self.login_attempts = 0

    def describe_user(self):
        print(f"Name: {self.first_name} {self.last_name}")
        print(f"Age: {self.age}")

    def increment_login_attempts(self):
        self.login_attempts += 1

    def reset_login_attempts(self):
        self.login_attempts = 0

    def greet_user(self):
        print(f"Hello there {self.first_name} {self.last_name}")


john = User("John", "Smith", 10)

john.increment_login_attempts()
john.increment_login_attempts()
john.increment_login_attempts()
john.increment_login_attempts()
print(john.login_attempts)

john.reset_login_attempts()
print(john.login_attempts)
