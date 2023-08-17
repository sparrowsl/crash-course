import pathlib

try:
    file = pathlib.Path("./cats.txt")
    print(file.read_text())
except FileNotFoundError as err:
    print(err)
    exit()


try:
    file = pathlib.Path("./dogs.txt")
    print(file.read_text())
except FileNotFoundError as err:
    print(err)
    exit()
