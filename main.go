package main

import (
	"github.com/gin-gonic/gin"
	"github.com/subintp/notes/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/notes/:id", controllers.GetNote)
	r.GET("/notes", controllers.GetNotes)
	r.POST("/notes", controllers.CreateNote)
	r.PUT("/notes/:id", controllers.UpdateNote)
	r.DELETE("/notes/:id", controllers.DeleteNote)

	r.Run()
}
