package models

type BookmarkCommand struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}
