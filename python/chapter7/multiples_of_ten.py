try:
    number = int(input("Enter a number: "))
except ValueError:
    print("Invalid number")
    exit(1)
else:
    if number % 10 != 0:
        print(f"{number} is not a multiple of 10!")
    else:
        print("{number} is a multiple of 10!")
