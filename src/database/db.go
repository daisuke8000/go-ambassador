package database

import (
	"ambassador-backend/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	var err error
	// How to
	// gorm.Open(mysql.Open("{MYSQL_USER}:{MYSQL_PASSWORD}@tcp({DOCKER_COMPOSE_DATABASE_NAME}:{DOCKER_COMPOSE_DATABASE_CONTAINER_PORTS})/{{DOCKER_COMPOSE_DATABASE_ENVIRONMENT_DATABASE_NAME}}"), &gorm.Config{})
	DB, err = gorm.Open(mysql.Open("root:root@tcp(db:3306)/ambassador"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database!")
	}
}

func AutoMigrate(){
	DB.AutoMigrate(models.User{}, models.Product{}, models.Link{})
}