package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/isaac/project/bgft_go/models"
  "github.com/gobuffalo/uuid"
   "fmt"
	 "log"
   // "net/url"
   "io/ioutil"
  // "strconv"
  "encoding/json"
	"github.com/gobuffalo/pop"
)

 var db = make(map[uuid.UUID]models.User)

type UserResource struct{}

func (ur UserResource) List(c buffalo.Context) error {
	return c.Render(200, r.JSON(db))
}

func (ur UserResource) Create(c buffalo.Context) error{
  body, error := ioutil.ReadAll(c.Request().Body)
        if error != nil {
            fmt.Printf("Error reading body: %v", error)
            return error
        }
    user := &models.User{}
    json.Unmarshal([]byte(body), user)
    user.ID = uuid.Must(uuid.NewV4())

		tx, err := pop.Connect("development")
		if err != nil {
			log.Panic(err)
		}
		verrs, err := tx.ValidateAndCreate(user)
	if err != nil {
		log.Panic(err)
	}
	tx.ValidateAndSave(verrs)
		db[user.ID] = *user


  //
	 return c.Render(201, r.JSON(user))
}
