package painter

import (
	"os"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

const TEST_INSPIRATION_STORE = "test_inspiration.json"

func TestPainter_InspiresThePainterWithImages(t *testing.T) {
	defer gock.Off()
	defer os.Remove(TEST_INSPIRATION_STORE)

	gock.New(ART_INSTITUTE_OF_CHICAGO).
		Reply(200).
		JSON(map[string]any{
			"data":   map[string]string{"title": "a painting", "artist_title": "an artist", "image_id": "id"},
			"config": map[string]string{"iiif_url": "https://iiif_url.com"},
		})

	painter := NewPainter()
	painter.inspirationStore = TEST_INSPIRATION_STORE

	err := painter.getInspiration()
	assert.NoError(t, err)

	assert.FileExists(t, TEST_INSPIRATION_STORE)
}

func TestPainter_NoInspirationStoredWhenError(t *testing.T) {
	defer gock.Off()

	gock.New(ART_INSTITUTE_OF_CHICAGO).
		Reply(403)

	painter := NewPainter()
	painter.inspirationStore = TEST_INSPIRATION_STORE

	painter.getInspiration()
	assert.NoFileExists(t, TEST_INSPIRATION_STORE)
}
