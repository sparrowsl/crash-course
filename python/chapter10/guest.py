from pathlib import Path

name = input("Enter name: ")

with Path("./guest.txt") as file:
    file.write_text(name)
