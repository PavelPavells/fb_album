package main

import (
	"fmt"
	"log"

	"github.com/PavelPavells/fb_album"
)

func main() {
	fbToken := ""
	fb_album.NewFbAlbum(fbToken)

	var targetAlbumId string
	var albumPhotoCount int
	var err error

	albums, err := fb_album.GetMyAlbums()

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	for _, album := range albums.Data {
		fmt.Println(album.Name)
		targetAlbumId = album.ID
		albumPhotoCount = album.Count

		break
	}

	fmt.Println("Album ID: ", targetAlbumId)

	photos, err := fb_album.GetPhotoByAlbum(targetAlbumId, albumPhotoCount)

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	fmt.Println("My photos: ", len(photos.Data))

	for _, photo := range photos.Data {
		fmt.Println(photo.ID)
	}
}
