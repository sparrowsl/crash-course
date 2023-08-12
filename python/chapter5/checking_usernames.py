current_users = ["john", "jane", "adam", "jack", "jenny", "barry", "joe"]
new_users = ["jonas", "jane", "Jenny", "angela", "jack"]

for user in new_users:
    if user.lower() in current_users:
        print("Username already taken, choose another name.")
    else:
        print("Username is available!")
