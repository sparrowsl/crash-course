import json
from pathlib import Path

filename = "./fav_number.json"


def readSavedNumber(filename: str):
    try:
        file = Path(filename)
    except FileExistsError:
        print("File does not exists!!")
        exit()
    else:
        value = json.loads(file.read_text())
        print(f"I know your favorite number!!. It's  {value['number']}")


def saveNumber(number: int):
    details = json.dumps({"number": number})

    with Path(filename) as file:
        file.write_text(details)


number = input("Enter your favorite number: ")

try:
    number = int(number)
except ValueError:
    print("That is not a valid number!!")
    exit()

saveNumber(number)
readSavedNumber(filename)
