package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Album name: ")
		scanner.Scan()
		album := scanner.Text()

		if album == "quit" {
			return
		}

		fmt.Printf("Artist name: ")
		scanner.Scan()
		artist := scanner.Text()

		if artist == "quit" {
			return
		}

		fmt.Println(makeAlbum(artist, album))
		fmt.Println()
	}

}

func makeAlbum(artist, title string) map[string]string {
	album := map[string]string{
		"artist": artist,
		"title":  title,
	}
	return album
}
