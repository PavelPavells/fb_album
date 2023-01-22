package fbAlbum

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	fb "github.com/huandu/facebook"
)

type FbAlbum struct {
	Token string
}

func NewFbAlbum(token string) *FbAlbum {
	if token == "" {
		return nil;
	}

	newAlbum := new(FbAlbum);
	newAlbum.Token = token;

	return newAlbum;
}

func (myAlbums *FbAlbum) GetMyAlbums() (*FBAlbums, error) {
	return myAlbums.GetAlbumsByUserId("me")
}

func (myAlbums *FbAlbum) GetAlbumsByUserId(uuid string) (*FBAlbums, error) {
	if uuid == "" {
		return nil, errors.New("uuid is empty")
	}

	resAlbum := myAlbums.RunFBGraphAPI("/" + uuid + "/albums")
	retAlbum := FBAlbums{}

	ParseMapToStruct(resAlbum, &retAlbum)

	return &retAlbum, nil
}

func (myAlbums *FbAlbum) GetPhotoByAlbum(albumId string, count int) (*FBPhotos, error) {
	if albumId == "" {
		return nil, errors.New("albumId is empty")
	}

	photoRet := FBPhotos{}
	query := fmt.Sprintf("/%s/photos?limit=%d", albumId, count)
	resPhoto := myAlbums.RunFBGraphAPI(query)

	ParseMapToStruct(resPhoto, &photoRet)

	return &photoRet, nil
}

func (myAlbums *FbAlbum) RunFBGraphAPI(query string) (result interface{}) {
	result, error := fb.Get(query, fb.Params{
		"access_token": myAlbums.Token,
	})

	if error != nil {
		log.Fatalln("Facebook connect error, error = ", error.Error())
	}

	return;
}

func ParseMapToStruct(data interface{}, decode interface{}) {
	jret, _ := json.Marshal(data)
	err := json.Unmarshal(jret, &decode)

	if err != nil {
		log.Fatalln(err)
	}
}

