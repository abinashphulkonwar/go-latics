package handlers

import (
	"log"
	"time"

	"github.com/abinashphulkonwar/go-latics/db"
	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v2"
)

type ViewReqBody struct {
	Id string `json:"ID"`
}

func ViewsHandler(req *fiber.Ctx) error {

	body := req.Body()

	println(string(body))
	Session := db.Session

	if Session == nil {
		return fiber.NewError(500, "Session is nil")
	}
	id := gocql.TimeUUID()
	userId := "user"
	postId := "post456"
	createdAt := time.Now().UTC().Format("2023-01-02 15:04:05")
	views := 1

	if err := Session.Query(
		"INSERT INTO ViewsByUser (Id, UserId, PostId, CreatedAt, Views) VALUES (?, ?, ?, ?, ?)",
		id, userId, postId, createdAt, views).Exec(); err != nil {
		log.Fatal(err)
		return fiber.NewError(500, err.Error())

	}
	data, err := Session.Query("SELECT * from ViewsByUser").Iter().SliceMap()
	if err != nil {
		log.Fatal(err)
		return fiber.NewError(500, err.Error())
	}
	return req.JSON(fiber.Map{
		"Status": "ok",
		"Body":   data,
	})
}
