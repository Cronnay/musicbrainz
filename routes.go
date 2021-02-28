package main

import (
	"fmt"
	"net/http"

	mb "github.com/cronnay/musicbrainz/musicbrainz"
	"github.com/labstack/echo/v4"
)

// Index is the initial route
func Index(c echo.Context) error {
	return c.JSON(http.StatusOK, "Cygni - MusicBrainz API")
}

// GetArtist is the route to get artist based on MBID
func GetArtist(c echo.Context) error {
	mbid := c.Param("mbid")
	artist, error := mb.GetArtistFromMBID(mbid)
	if error != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Error received: %s", error))
	}

	return c.JSON(http.StatusOK, artist)
}
