package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"./controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	uc := controllers.NewAccountController(getSession())

	r := gin.Default()
	r.GET("/api/users", uc.GetUsers)
	r.GET("/api/users/:id", uc.GetOneUser)
	r.POST("/api/users/", uc.CreateUser)
	r.DELETE("/api/users/:id", uc.RemoveUser)
	r.PUT("/api/users/:id", uc.UpdateUser)

	if err := r.Run(":8898"); err != nil {
		log.Fatal(err.Error())
	}
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return s
}
