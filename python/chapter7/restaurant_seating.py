try:
    people = int(input("How many people are in the group? "))
except ValueError:
    print("Please enter a valid number!")
    exit(1)
else:
    if people > 8:
        print("Sorry, but you will have to wait for a table")
    else:
        print("The table is ready!!")
