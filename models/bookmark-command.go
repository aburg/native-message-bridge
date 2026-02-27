package models

type BookmarkCommand struct {
	SubCmd string `json:"subcmd"`
	Path   string `json:"path"`
	URL    string `json:"url"`
}
