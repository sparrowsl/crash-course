package main

import "fmt"

func main() {
	fmt.Println(makeAlbum("Don Moen", "Thank you Lord", 1))
	fmt.Println(makeAlbum("Adele", "Go easy on me", 30))
	fmt.Println(makeAlbum("Jon Bellion", "I feel it", 1))
}

func makeAlbum(artist, title string, songs int) map[string]interface{} {
	album := map[string]interface{}{
		"artist": artist,
		"title":  title,
		"songs":  songs,
	}
	return album
}
