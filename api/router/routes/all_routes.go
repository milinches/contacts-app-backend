package routes

import (
	"net/http"

	"github.com/milinches/contacts-app-backend/api/controllers/users"
	"github.com/milinches/contacts-app-backend/api/controllers/contacts"
)

type (
	Route struct {
		URI     string
		Method  string
		Handler func(http.ResponseWriter, *http.Request)
	}
)

var allRoutes = []Route{
	{
		URI:     "/users",
		Method:  http.MethodGet,
		Handler: users.GetUsers,
	},
	{
		URI:     "/user/create",
		Method:  http.MethodPost,
		Handler: users.CreateUser,
	},
	{
		URI:     "/user/{id}",
		Method:  http.MethodGet,
		Handler: users.GetUser,
	},
	{
		URI:     "/user/{id}",
		Method:  http.MethodPut,
		Handler: users.UpdateUser,
	},
	{
		URI:     "/user/{id}",
		Method:  http.MethodDelete,
		Handler: users.DeleteUser,
	},


	{
		URI:     "/contacts",
		Method:  http.MethodGet,
		Handler: contacts.GetContacts,
	},
	{
		URI:     "/contact/create",
		Method:  http.MethodPost,
		Handler: contacts.CreateContact,
	},
	{
		URI:     "/contact/{id}",
		Method:  http.MethodGet,
		Handler: contacts.GetContact,
	},
	{
		URI:     "/contact/{id}",
		Method:  http.MethodPut,
		Handler: contacts.UpdateContact,
	},
	{
		URI:     "/contact/{id}",
		Method:  http.MethodDelete,
		Handler: contacts.DeleteContact,
	},
}
