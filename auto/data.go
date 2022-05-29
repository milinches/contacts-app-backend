package auto

import "github.com/milinches/contacts-app-backend/api/models"

var (
	users = []models.User{
		{
			Email:     "john@mail.com",
			FirstName: "John",
			LastName:  "Inu",
			Password:  "neverhaveiever",
		},
		{
			Email:     "johns@mail.com",
			FirstName: "John",
			LastName:  "Inu",
			Password:  "neverhaveiever",
		},
	}

	contacts = []models.Contact{
		{
			Email:       "joma@mail.com",
			Name:        "Joma Tatsumi",
			Address:     "Acapella 16, in coza repulic",
			PhoneNumber: []string{"081678911", "+23451652388", "08167875645"},
			Type:        []string{"Personal"},
			UserID: 1,
		},
	}
)
