package main

import (
	"endor/db"
	"endor/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	myDb := db.NewRedisDB()
	if myDb == nil {
		panic(fmt.Sprintln("Cannot get new Redis DB"))
		return
	}
	handler := handlers.NewHandler(myDb)

	rGroup := r.Group("/api/objects/v1")
	rGroup.POST("/store", handler.Store)
	rGroup.GET("/id/:id", handler.GetObjectById)
	rGroup.GET("/name/:name", handler.GetObjectByName)
	rGroup.GET("/kind/:kind", handler.ListObjects)
	rGroup.DELETE("/id/:id", handler.DeleteObject)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(fmt.Sprintln("could not start gin", err))
	}
	/*
		animal1 := &model.Animal{
			Name:    "A1",
			ID:      "",
			Type:    "AT1",
			OwnerID: "OwnerA",
		}

		animal2 := &model.Animal{
			Name:    "A2",
			ID:      "",
			Type:    "AT1",
			OwnerID: "OwnerB",
		}

		animal3 := &model.Animal{
			Name:    "A3",
			ID:      "",
			Type:    "AT2",
			OwnerID: "OwnerA",
		}

		person1 := &model.Person{
			Name:     "P1",
			ID:       "",
			LastName: "P1L",
			Birthday: "B1",
		}
		person2 := &model.Person{
			Name:     "P2",
			ID:       "",
			LastName: "P2L",
			Birthday: "B2",
		}
	*/
}
