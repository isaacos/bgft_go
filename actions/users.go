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
	tx, err := pop.Connect("development")
	if err != nil {
		log.Panic(err)
	}

	users := []models.User{}
	error := tx.All(&users)
	if err != nil {
		log.Panic(error)
	}
	return c.Render(200, r.JSON(users))
}

func (ur UserResource) Create(c buffalo.Context) error{
  body, error := ioutil.ReadAll(c.Request().Body)
        if error != nil {
            fmt.Printf("Error reading body: %v", error)
            return error
        }
    user := &models.User{}
		//converts Request Body to JSON and adds it to a user instance. Then adds ID
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
	//addes user to db but has not connection to the Postgres database
		db[user.ID] = *user
	 return c.Render(201, r.JSON(user))
}


func (ur UserResource) Login(c buffalo.Context) error{
	body, error := ioutil.ReadAll(c.Request().Body)
        if error != nil {
            fmt.Printf("Error reading body: %v", error)
            return error
        }
			//create object and load in the body
			user := models.User{}
			json.Unmarshal([]byte(body), &user)


			tx, err := pop.Connect("development")
			if err != nil {
				log.Panic(err)
			}

			queryError := tx.Where("username = ?", user.UserName).Last(&user)
			if queryError != nil {
				log.Panic(queryError)
			}

	return c.Render(201, r.JSON(user))
}
