package db

type ViewsByUser struct {
	Id     string
	UserID string
	PostId string
	Views  int
}

type ViewsByPost struct {
	Id     string
	UserID string
	PostId string
	Views  int
}
