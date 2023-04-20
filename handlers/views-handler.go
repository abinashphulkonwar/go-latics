package handlers

import (
	"fmt"
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
	userId := "userid"
	postId := "post456"
	createdAt := time.Now().UTC().Format("2023-01-02 15:04:05")
	views := 1

	if err := Session.Query(
		"INSERT INTO ViewsByUser (Id, UserId, PostId, CreatedAt, Views) VALUES (?, ?, ?, ?, ?) IF NOT EXISTS",
		id, userId, postId, createdAt, views).Exec(); err != nil {
		log.Fatal(err)
		return fiber.NewError(500, err.Error())
	}
	// if err := Session.Query(
	// 	"UPDATE ViewsByUser SET Views = Views+? WHERE UserID =? AND PostId =? IF EXISTS",
	// 	1, userId, postId).Exec(); err != nil {
	// 	log.Fatal(err)
	// 	return fiber.NewError(500, err.Error())

	// }
	data := Session.Query("SELECT * from ViewsByUser WHERE UserId = ? AND PostId = ? limit 1", userId, postId).Consistency(gocql.One).Iter().Scanner()

	viewByUser := db.ViewsByUser{}

	for data.Next() {
		var (
			userid    string
			postid    string
			createdat string
			id        gocql.UUID
			views     int
		)
		err := data.Scan(&userid, &postid, &createdat, &id, &views)
		if err != nil {
			return fiber.NewError(500, err.Error())

		}
		viewByUser.Id = id.String()
		viewByUser.UserID = userid
		viewByUser.PostID = postid
		viewByUser.CreatedAt = createdat
		viewByUser.Views = views
		fmt.Println(id, userid, postid, createdat, views)
	}

	return req.JSON(viewByUser)
}
