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
	Price       int      `json:"price"`
	Bedrooms    int      `json:"bedrooms"`
	Bathrooms   int      `json:"bathrooms"`
	Mts2        int      `json:"mts2"`
	UserId      int      `json:"userId"`
}

type Items []Item

type RespuestaItem struct {
	Items          Items `json:"items"`
	HttpStatusCode int   `json:"http_status_code"`
}
