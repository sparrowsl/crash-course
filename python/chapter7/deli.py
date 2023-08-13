sandwich_orders = ["bread", "apple", "vanilla", "brown", "tuna"]
finished_sandwiches = list(str())

for v in sandwich_orders:
    print(f"I made your {v} sandwich")
    finished_sandwiches.append(v)

print(finished_sandwiches)
