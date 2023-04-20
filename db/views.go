package db

type ViewsByUser struct {
	Id        string
	UserID    string
	PostID    string
	Views     int
	CreatedAt string
}

type ViewsByPost struct {
	Id     string
	UserID string
	PostId string
	Views  int
}
