import pathlib

try:
    file = pathlib.Path("./romeo_juliet.txt")
except FileNotFoundError as err:
    print(err)
    exit()
else:
    count = str(file.read_text()).lower().count("juliet".lower())
    print(count)
