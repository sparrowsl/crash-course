def sandwich_items(*items: str):
    print("The sandwich items:")

    for item in items:
        print(item)

    print()


sandwich_items("cake", "sugar")
sandwich_items("mint", "pepper", "cream", "milk")
sandwich_items("chocolate")
