def make_album(artist: str, title: str, songs: int):
    album = {"artist": artist, "title": title, "songs": songs}
    return album


print(make_album("Don Moen", "Thank you Lord", 1))
print(make_album("Adele", "Go easy on me", 30))
print(make_album("Jon Bellion", "I feel it", 1))
