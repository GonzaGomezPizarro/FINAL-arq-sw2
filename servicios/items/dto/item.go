package dto

type Item struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Country     string   `json:"country"`
	State       string   `json:"state"`
	City        string   `json:"city"`
	Address     string   `json:"address"`
	Photos      []string `json:"photos"`
	UserId      int      `json:"user"`
}

type Items []Item
