def make_album(artist: str, title: str):
    album = {"artist": artist, "title": title}
    return album


while True:
    album = input("Album name: ")
    if album.lower() == "quit":
        break

    artist = input("Artist name: ")
    if artist.lower() == "quit":
        break

    print(make_album(artist, album))
    print()
