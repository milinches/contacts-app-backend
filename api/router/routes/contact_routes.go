package routes

import (
	"net/http"

	"github.com/milinches/contacts-app-backend/api/controllers"
)

var postRoutes = []Route{
	{
		URI:     "/contacts",
		Method:  http.MethodGet,
		Handler: controllers.GetContacts,
	},
	{
		URI:     "/contact/create",
		Method:  http.MethodPost,
		Handler: controllers.CreateContact,
	},
	{
		URI:     "/contact/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetContact,
	},
	{
		URI:     "/contact/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateContact,
	},
	{
		URI:     "/contact/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteContact,
	},
}