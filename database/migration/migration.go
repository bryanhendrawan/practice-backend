package migration

import (
	"fmt"
	"log"

	"practice-backend/database"
	"practice-backend/model/entity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.Article{})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database migration success")
}
