package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/subintp/notes/models"
	"github.com/subintp/notes/utils"
)

// GetNote - get note by id
func GetNote(c *gin.Context) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("GetNote")
	defer span.Finish()

	var note models.Note
	id := c.Params.ByName("id")
	db := utils.GetDbConnection()

	if err := db.First(&note, id).Error; err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"error": "Note not found",
			},
		)
	} else {
		c.JSON(http.StatusOK, note)
	}
}

// GetNotes - get all notes
func GetNotes(c *gin.Context) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("GetNotes")
	defer span.Finish()

	var notes []models.Note
	db := utils.GetDbConnection()

	if err := db.Find(&notes).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, notes)
	}
}

// CreateNote - create note
func CreateNote(c *gin.Context) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("CreateNote")
	defer span.Finish()

	var note models.Note
	db := utils.GetDbConnection()
	c.BindJSON(&note)

	if err := db.Create(&note).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, note)
	}
}

// UpdateNote - update note
func UpdateNote(c *gin.Context) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("UpdateNote")
	defer span.Finish()

	var note models.Note
	id := c.Params.ByName("id")
	db := utils.GetDbConnection()

	if err := db.First(&note, id).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.BindJSON(&note)

	if err := db.Save(&note).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, note)
	}
}

// DeleteNote - delete note by id
func DeleteNote(c *gin.Context) {
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("DeleteNote")
	defer span.Finish()

	var note models.Note
	id := c.Params.ByName("id")
	db := utils.GetDbConnection()

	if err := db.First(&note, id).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	if err := db.Delete(&note).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Deleted",
		})
	}
}
