package actions

import "github.com/gobuffalo/buffalo"

// SightingsCreate default implementation.
func SightingsCreate(c buffalo.Context) error {
	return c.Render(200, r.HTML("sightings/create.html"))
}

