try:
    num1, num2 = input("Enter 2 numbers: ").split(" ")
    num1 = int(num1)
    num2 = int(num2)
except ValueError:
    print("Please enter a valid number!!")
    exit()
else:
    print(num1 + num2)
