while True:
    try:
        age = int(input("Enter the age: "))
    except ValueError:
        print("Enter a valid age!!\n")
    else:
        if age < 3:
            print("Movie ticket is free!")
        if age >= 3 and age <= 12:
            print("Movie ticket is $10.")
        if age > 12:
            print("Movie ticket is $15.")
