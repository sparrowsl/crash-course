names = ["john", "jane", "admin", "jack", "jenny"]
names = list(str())

if len(names) != 0:
    for name in names:
        if name == "admin":
            print("Hello admin, would you like to see a status report?")
        else:
            print(f"Hello {name}, thank you for logging in again.")
else:
    print("We need to find some users!!")
