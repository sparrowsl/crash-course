pets = {
    "sparrow": {
        "animal": "Bird",
        "owner": "Sparrow",
    },
    "bolt": {
        "animal": "Dog",
        "owner": "Sparrow",
    },
}

for key, value in pets.items():
    print("%s\n", key)

    for k, v in value.items():
        print(f"{k}: {v}")

    print()
