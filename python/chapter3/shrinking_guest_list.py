guest_list = ["Melkey", "Johnny Depp", "Linus Trovalds"]

print(f"Dear {guest_list[0]}, you are invited to my party!!")
print(f"Dear {guest_list[1]}, you are invited to my party!!")
print(f"Dear {guest_list[2]}, you are invited to my party!!")

print(f"\nUnf{guest_list[2]}tunately %s can't make it.\n")

guest_list = ["Melkey", "Johnny Depp", "Rich Harris"]
print(f"Dear {guest_list[0]}, you are invited to my party!!")
print(f"Dear {guest_list[1]}, you are invited to my party!!")
print(f"Dear {guest_list[2]}, you are invited to my party!!")

print("\nI found a bigger table for the guests!!\n")
guest_list.append("The Primegean")
guest_list.append("Wagslane")

print("\nSorry guys, I can only invite 2 guests!!\n")
guest_list.pop()
guest_list.pop()
guest_list.pop()

print(guest_list)
