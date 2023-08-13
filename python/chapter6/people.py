people = {
    "Jack": {
        "firstName": "Jack",
        "lastName": "Doe",
        "age": 10,
        "city": "Freetown",
    },
    "Kate": {
        "firstName": "Kathrine",
        "lastName": "Smith",
        "age": 13,
        "city": "Detroit",
    },
}

for key, value in people.items():
    print(f"{key}")

    for k, v in value.items():
        print(f"{k}: {v}")

    print()
