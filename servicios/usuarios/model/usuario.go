package model

type User struct {
	Id        int    `gorm:"primaryKey;AUTO_INCREMENT"`
	Username  string `gorm:"type:varchar(45);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"` //se guarda encryptada
	FirstName string `gorm:"type:varchar(255);not null"`
	LastName  string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);not null;unique"`
	phone     int    `gorm:"type:integer;not null"`
}

type Users []User

type Credenciales struct {
	Username string
	Password string
}
