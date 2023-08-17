from pathlib import Path

file = Path("./learning_python.txt")
print(file.read_text())

for line in file.read_text().splitlines():
    print(line)
