package routes

import (
	"net/http"

	"github.com/milinches/contacts-app-backend/api/controllers"
)

var userRoutes = []Route{
	{
		URI:     "/users",
		Method:  http.MethodGet,
		Handler: controllers.GetUsers,
	},
	{
		URI:     "/user/create",
		Method:  http.MethodPost,
		Handler: controllers.CreateUser,
	},
	{
		URI:     "/user/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetUser,
	},
	{
		URI:     "/user/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateUser,
	},
	{
		URI:     "/user/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteUser,
	},
}
