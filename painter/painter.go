package painter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
)

const ART_INSTITUTE_OF_CHICAGO = "https://api.artic.edu/api/v1/artworks?page=1&limit=30&fields=title,artist_title,image_id"
const INSPIRATION_STORE = "inspiration.json"

var cursor = 0

func NewPainter() *Painter {
	return &Painter{
		sourceOfInspiration: ART_INSTITUTE_OF_CHICAGO,
		inspirationStore:    INSPIRATION_STORE,
	}
}

func (p *Painter) PaintOn(canvas Canvas) error {
	p.getInspiration()
	if p.hasNoInspiration() {
		return errors.New("No inspiration found")
	}

	inspirations := p.loadInspirations()

	manifest := inspirations.Manifest
	artWork := manifest[cursor]
	artConfig := inspirations.Config
	art := fmt.Sprintf("%s/%s/full/843,/0/default.jpg", artConfig.ImgUrl, artWork.Id)

	fmt.Println("Drawing now")
	canvas.Draw(art, artWork.Title, artWork.Artist)
	return nil
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

func (p *Painter) hasNoInspiration() bool {
	if _, err := os.Stat(p.inspirationStore); errors.Is(err, fs.ErrNotExist) {
		return true
	}
	return false
}

func (p *Painter) save(inspirations *http.Response) {
	storage, err := os.Create(p.inspirationStore)
	if weGotAnError(err) {
		panic("I can't store inspirations")
	}
	defer storage.Close()

	io.Copy(storage, inspirations.Body)

}

func (p *Painter) loadInspirations() Resp {
	fmt.Println("Loading")
	var insp Resp
	inspirations, _ := os.ReadFile(p.inspirationStore)
	json.Unmarshal(inspirations, &insp)
	return insp
}

func weGotAnError(err error) bool {
	if err != nil {
		return true
	}
	return false
}

type Canvas interface {
	Draw(art, title, artist string)
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
