package auto

import (
	"log"

	"github.com/milinches/contacts-app-backend/api/database"
	"github.com/milinches/contacts-app-backend/api/models"
	"github.com/milinches/contacts-app-backend/api/utils"
)

func Load() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("error loading the db: %v", err)
	}

	log.Println("successfully connected to the database...")

	if err = db.Debug().Migrator().DropTable(&models.User{}, &models.Contact{}); err != nil {
		log.Fatalf("error dropping table: %v", err)
	}

	if err = db.Debug().AutoMigrate(&models.User{}, &models.Contact{}); err != nil {
		log.Fatalf("error migrating table: %v", err)
	}

	for _, u := range users {
		if err = db.Debug().Model(&models.User{}).Create(&u).Error; err != nil {
			log.Fatalf("cannot create table: %v", err)
		}

		utils.Prettier(u)
	}

	for _, c := range contacts {
		if err = db.Debug().Model(&models.Contact{}).Create(&c).Error; err != nil {
			log.Fatalf("cannot create table: %v", err)
		}

		utils.Prettier(c)
	}
}

// -- Options --

// db.Model(&models.User{}).Association("Contact")
