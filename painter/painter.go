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

const INSPIRATION_STORE = "inspiration.json"
const CursorStore = "cursor.json"

var page = 1

func New() *Painter {
	return &Painter{
		sourceOfInspiration: fromPage(1),
		inspirationStore:    INSPIRATION_STORE,
		cursor: cursor{
			page:  1,
			index: 0,
		},
	}
}

func (p *Painter) PaintOn(canvas Canvas) error {
	// if the cursor turned a new page, fetch new inspiration
	if p.cursor.turnedPage {
		p.sourceOfInspiration = fromPage(p.cursor.page)
		p.GetInspiration()
	}

	if p.hasNoInspiration() {
		return errors.New("No inspiration found")
	}

	inspirations := p.loadInspirations()

	manifest := inspirations.Manifest
	artWork := manifest[p.cursor.index]
	artConfig := inspirations.Config
	artSrc := fmt.Sprintf("%s/%s/full/843,/0/default.jpg", artConfig.ImgUrl, artWork.Id)

	fmt.Println(p.cursor.index, artWork.Title)
	err := canvas.Draw(artSrc, artWork.Title, artWork.Artist)
	if err != nil {
		return err
	}

	p.cursor.next()
	return nil
}

func (p *Painter) GetInspiration() error {
	inspirations, err := http.Get(p.sourceOfInspiration)
	if err != nil {
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
	if err != nil {
		panic("I can't store inspirations")
	}
	defer storage.Close()

	io.Copy(storage, inspirations.Body)
}

func (p *Painter) loadInspirations() Result {
	var res Result
	inspirations, _ := os.ReadFile(p.inspirationStore)
	json.Unmarshal(inspirations, &res)
	return res
}

func fromPage(page int) string {
	ART_INSTITUTE_OF_CHICAGO := "https://api.artic.edu/api/v1/artworks?page=%d&limit=30&fields=title,artist_title,image_id"
	return fmt.Sprintf(ART_INSTITUTE_OF_CHICAGO, page)
}

type Canvas interface {
	Draw(art, title, artist string) error
}

type Painter struct {
	sourceOfInspiration string
	inspirationStore    string
	cursor              cursor
}

type Result struct {
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
