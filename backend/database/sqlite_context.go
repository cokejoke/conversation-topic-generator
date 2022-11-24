package database

import (
	"conversation-topic/errorhandler"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var DBConn *gorm.DB

func Connect() {
	var err error
	DBConn, err = gorm.Open("sqlite3", "./data/conversation-topic.db")
	errorhandler.Handle(err)

	err = DBConn.DB().Ping()
	errorhandler.Handle(err)

	log.Println("Connected to sqlite3 database")
}
