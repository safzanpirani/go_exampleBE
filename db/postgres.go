package db

import (
	"fmt"
	"gin/basic/model"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDB *gorm.DB

func GetDBConn() *gorm.DB {
	return PostgresDB
}
func GetPostgresDBConnection() (*gorm.DB, error) {
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", host, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(model.Post{}) //helping to create table automatically
	db.AutoMigrate(model.User{})
	PostgresDB = db
	return db, err
}

//Fly command line tool configuration for PostgreSQL
//flyctl secrets set DB_NAME=~
//flyctl secrets set HOST=postgresql://HOSTNAME_URL
//flyctl secrets set PASSWORD=~
//flyctl secrets set USER=~

