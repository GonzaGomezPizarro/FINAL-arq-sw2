package db

import (
	client "github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/mensajes/client" // puede faltar el /user
	"github.com/GonzaGomezPizarro/FINAL-arq-sw2/servicios/mensajes/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {

	dsn := "user:pass@tcp(127.0.0.1:3307)/mensajes_db?charset=utf8mb4&parseTime=True&loc=Local" // 3307 o 3306 ??
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	client.Db = db
}
func StartDbEngine() {

	db.AutoMigrate(&model.Message{}) // crea una tabla en plural de "Message" o la usa si esta creada

	log.Info("Finishing Migration Database Tables")
}
