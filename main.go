package main

import (
	"github.com/gin-gonic/gin"
	"github.com/subintp/notes/controllers"
	"github.com/subintp/notes/utils"
	"github.com/subintp/tracer/test"
)

func main() {
	r := gin.Default()

	utils.InitTracer("NoteService")

	r.GET("/notes/:id", controllers.GetNote)
	r.GET("/notes", controllers.GetNotes)
	r.POST("/notes", controllers.CreateNote)
	r.PUT("/notes/:id", controllers.UpdateNote)
	r.DELETE("/notes/:id", controllers.DeleteNote)

	test.Hello()

	r.Run()
}
