package painter

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const ART_INSTITUTE_OF_CHICAGO = "https://api.artic.edu/api/v1/artworks?page=1&limit=30&fields=title,artist_title,image_id"
const INSPIRATION_STORE = "inspiration.json"

func NewPainter() *Painter {
	return &Painter{
		sourceOfInspiration: ART_INSTITUTE_OF_CHICAGO,
		inspirationStore:    INSPIRATION_STORE,
	}
}

func (p *Painter) getInspiration() error {
	inspirations, err := http.Get(p.sourceOfInspiration)
	if weGotAnError(err) {
		return err
	}
	defer inspirations.Body.Close()

	if inspirations.StatusCode >= 400 && inspirations.StatusCode < 600 {
		return fmt.Errorf("Received status code: %d", inspirations.StatusCode)
	}

	p.save(inspirations)
	return nil
}

func (p *Painter) save(inspirations *http.Response) {
	storage, err := os.Create(p.inspirationStore)
	if weGotAnError(err) {
		panic("I can't store inspirations")
	}
	defer storage.Close()

	io.Copy(storage, inspirations.Body)

}

func weGotAnError(err error) bool {
	if err != nil {
		return true
	}
	return false
}

type Painter struct {
	sourceOfInspiration string
	inspirationStore    string
}

type Resp struct {
	Manifest []ArtWork `json:"data"`
	Config   ArtConfig `json:"config"`
}

type ArtWork struct {
	Title  string `json:"title"`
	Artist string `json:"artist_title"`
	Id     string `json:"image_id"`
}

type ArtConfig struct {
	ImgUrl string `json:"iiif_url"`
}
