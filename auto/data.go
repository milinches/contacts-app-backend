package auto

import "github.com/milinches/contacts-app-backend/api/models"

var users = []models.User{
	{
		Username:  "milinches",
		Email:     "made@mail.com",
		FirstName: "John",
		LastName:  "Inu",
		OtherName: "WAGMI",
		Password:  "1233345",
	},
}

var contacts = []models.Contact{
	{
		Email:       "ilysm@mail.com",
		FirstName:   "James",
		LastName:    "Rodrigo",
		OtherName:   "Inu",
		Address:     "Acapella 16, @mail.com in coza repulic",
		PhoneNumber: []string{"081678911", "+23451652388"},
	},
}
