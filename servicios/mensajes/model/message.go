package model

type Message struct {
	Id         int    `gorm:"primaryKey;AUTO_INCREMENT"`
	Content    string `gorm:"type:varchar(511);not null"`
	ReceiverId int    `gorm:"type:integer;not null"`
}

type Messages []Message
