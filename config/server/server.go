package server

import (
	"github.com/gin-gonic/gin"
	route_user "github.com/guizot/go-gin-mongodb/src/routes/user"
)

type Server struct {
}

func (c Server) StartGin() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/users", route_user.GetAllUser)
		api.POST("/users", route_user.CreateUser)
		api.GET("/users/:id", route_user.GetUser)
		api.PUT("/users/:id", route_user.UpdateUser)
		api.DELETE("/users/:id", route_user.DeleteUser)
	}
	r.Run(":8000")
}

/*
	File ini men-define endpoint dari macam-macam routes
	yang di buat pada folder src/routes
*/
