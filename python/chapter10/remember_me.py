import json
from pathlib import Path


def save_user_info():
    name = input("Enter name: ")

    try:
        age = int(input("Enter age: "))
    except ValueError:
        print("Please enter a valid age!!")
        return

    contact = input("Enter contact: ")

    with Path("./remember.json") as file:
        details = json.dumps({"name": name, "age": age, "contact": contact})
        file.write_text(details + "\n")


def display_user_info():
    try:
        file = Path("./remember.json")
        contents = json.loads(file.read_text())
        print(contents)
    except FileNotFoundError:
        print("File does not exist!!")
        return


display_user_info()
save_user_info()
