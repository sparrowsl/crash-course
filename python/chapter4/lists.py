pizzas = ["pepper", "pizza", "phone", "cheese", "vanilla", "laptop", "watch"]

print("The first three items in the list are:")
print("\t", pizzas[:3])

print("Three items from the middle of the list are:")
print("\t", pizzas[int(len(pizzas) / 2) : int(len(pizzas) / 2) + 3])

print("The last three items in the list are:")
print("\t", pizzas[:-3])
