package musicbrainz

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	musicBrainzAPIURL = "http://musicbrainz.org/ws/2/artist/%s?&fmt=json&inc=url-rels+release-groups"
)

// ResponseArtist is a struct that have all information about an artist
type ResponseArtist struct {
	Country      string  `json:"country"`
	Name         string  `json:"name"`
	ID           string  `json:"id"`
	Begin        string  `json:"begin"`
	End          string  `json:"end"`
	Albums       []Album `json:"albums"`
	Biography    string  `json:"biography"`
	BiographyRaw string  `json:"biographyRaw"`
}

// Album is a struct with information about albums of an album
type Album struct {
	Title       string
	ReleaseDate string
	ID          string
}

type artist struct {
	Country       string         `json:"country"`
	Type          string         `json:"type"`
	ReleaseGroups []releaseGroup `json:"release-groups"`
	Relations     []relation     `json:"relations"`
	Name          string         `json:"name"`
}

type lifespan struct {
	Begin string `json:"begin"`
	End   string `json:"end"`
	Ended bool   `json:"ended"`
}

type releaseGroup struct {
	FirstReleaseDate string `json:"first-release-date"`
	Title            string `json:"title"`
	ID               string `json:"id"`
	PrimaryTypeID    string `json:"primary-type-id"`
	PrimaryType      string `json:"primary-type"`
}

type relation struct {
	Type string `json:"type"`
	URL  struct {
		Resource string `json:"resource"`
	} `json:"url"`
}

// GetArtistFromMBID is a method that gets information about an artist from MusicBrainz ID
func GetArtistFromMBID(mbid string) (ResponseArtist, error) {
	r, err := http.Get(fmt.Sprintf(musicBrainzAPIURL, mbid))
	if err != nil {
		return ResponseArtist{}, err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return ResponseArtist{}, fmt.Errorf("Response code %d", r.StatusCode)
	}

	mbArtist := new(artist)

	decodeError := json.NewDecoder(r.Body).Decode(mbArtist)
	if decodeError != nil {
		return ResponseArtist{}, decodeError
	}

	respArtist := generateResponseArtist(*mbArtist)

	return *respArtist, nil
}

func generateResponseArtist(a artist) *ResponseArtist {
	r := new(ResponseArtist)
	return r
}
