package auto

import (
	"log"

	"github.com/milinches/contacts-app-backend/api/database"
	"github.com/milinches/contacts-app-backend/api/models"
	"github.com/milinches/contacts-app-backend/api/utils/console"
)

func Load() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("successfully connected to the database...")

	// err = db.Debug().AutoMigrate(&models.User{} ,&models.Contact{})
	err = db.Debug().AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		err = db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			log.Fatal(err)
		}

		console.Prettier(user)
	}

	// for _, c := range contacts {
	// 	err = db.Debug().Model(&models.Contact{}).Create(&c).Error
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	console.Prettier(c)
	// }
}
