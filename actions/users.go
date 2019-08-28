package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/isaac/project/bgft_go/models"
	"github.com/satori/go.uuid"
)

 var db = make(map[uuid.UUID]models.User)

type UserResource struct{}

func (ur UserResource) List(c buffalo.Context) error {
	return c.Render(200, r.JSON(db))
}
