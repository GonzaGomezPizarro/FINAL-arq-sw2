package dto

type Message struct {
	Id         int    `json:"id"`
	Content    string `json:"content"`
	ReceiverId int    `json:"receiver"`
	ItemId     string `json:"item"`
}

type Messages []Message
