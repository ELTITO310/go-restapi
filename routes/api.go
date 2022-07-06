package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	db "github.com/sonyarouje/simdb"
)

type User struct {
	CustID  string `json:"custid"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Contact Contact
}

type Contact struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func (c User) ID() (jsonField string, value interface{}) {
	value = c.CustID
	jsonField = "custid"
	return
}

func Register(r fiber.Router) {
	r.Get("/user", handleHome)
	r.Get("/user/:id", handleGetUser)
	r.Post("/user", handleNew)
	r.Put("/user/:id", handleUpdate)
	r.Delete("/user/:id", handleDelete)
}

func handleHome(c *fiber.Ctx) error {
	driver := connect()

	var users []User
	err := driver.Open(User{}).Get().AsEntity(&users)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(users)
}

func handleGetUser(c *fiber.Ctx) error {
	driver := connect()

	id := c.Params("id")

	var u User

	err := driver.Open(User{}).Where("custid", "=", id).First().AsEntity(&u)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(u)

}

func handleUpdate(c *fiber.Ctx) error {
	driver := connect()

	u := new(User)

	if err := c.BodyParser(u); err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var user User

	err := driver.Open(User{}).Where("custid", "=", c.Params("id")).First().AsEntity(&user)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user.Address = u.Address
	user.Name = u.Name
	user.Contact = u.Contact
	user.CustID = u.CustID

	err = driver.Open(User{}).Update(user)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(u)

}

func handleNew(c *fiber.Ctx) error {

	driver := connect()

	u := new(User)

	if err := c.BodyParser(u); err != nil {
		return c.JSON(fiber.Map{
			"type":  "bodyParser",
			"error": err.Error(),
		})
	}

	err := driver.Open(User{}).Insert(u)
	if err != nil {
		return c.JSON(fiber.Map{
			"type":  "Insert",
			"error": err.Error(),
		})
	}

	return c.JSON(u)
}

func handleDelete(c *fiber.Ctx) error {
	driver := connect()
	id := c.Params("id")

	toDel := User{
		CustID: id,
	}

	err := driver.Open(User{}).Delete(toDel)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Deleted User",
	})

}

func connect() *db.Driver {
	driver, err := db.New("data")
	if err != nil {
		log.Fatal(err)
	}
	return driver
}
