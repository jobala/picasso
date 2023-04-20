package painter

import "log"

type cursor struct {
	page       int
	index      int
	turnedPage bool
}

func (c *cursor) next() {
	if c.index == 29 {
		c.page += 1
		c.index = 0
		c.turnedPage = true

		log.Printf("\n\nServing results from %d\n\n", c.page)
	} else {
		c.index += 1
		c.turnedPage = false
	}
}
