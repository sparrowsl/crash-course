from src.city_country import city_country


def test_city_country():
    value = city_country("paris", "France")
    assert value == "Paris, France - population 0"


def test_city_country_population():
    value = city_country("santiago", "chile", 5_000_000)
    assert value == "Santiago, Chile - population 5000000"
