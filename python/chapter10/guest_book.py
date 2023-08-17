while True:
    name = input("Enter name: ")
    if name.lower() == "quit":
        break

    with open("./guest_book.txt", "a") as file:
        file.write(name + "\n")
