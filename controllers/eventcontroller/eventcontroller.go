package eventcontroller

import (
	"clockworks-backend/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

/** Indexes ALL events. */
func Index(c *gin.Context) {
	var events []models.Event

	models.DB.Find(&events)
	c.JSON(200, gin.H{
		"data": events,
	})
}

/** Count number of events with a specified ID (should be <= 1) */
func CountIds(c *gin.Context) {
	var id string
	var count int64

	c.BindJSON(&id)
	fmt.Println(id)
	models.DB.Select("id").Where("id = ?", id).Count(&count)

	c.JSON(200, gin.H{"count": count})
}

/** Return a event with a specified ID */
func Show(c *gin.Context) {
	var event models.Event
	var id = c.Param("id")

	result := models.DB.Where("id = ?", id).First(&event)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"id":      id,
			"message": "No Event with specified ID found.",
		})
	} else {
		c.JSON(200, gin.H{
			"data": event,
		})
	}
}

/** Creates a new Event, given event data. Handles duplicate ID */
func Create(c *gin.Context) {
	var event models.Event
	var body models.Event

	err := c.BindJSON(&body)

	if err != nil || hasInvalidField(body) {
		fmt.Println("error creating. body: ", body)
		c.JSON(400, gin.H{
			"message": "Invalid body.",
		})
	} else {
		// structs are DEEP-COPIED by default. wowzers???
		event = body

		// if event with identical PK exists, return it. else, create new.
		models.DB.FirstOrCreate(&body)
		if event == body {
			c.JSON(200, gin.H{
				"data":    event,
				"message": "Successfully created event.",
			})
		} else {
			// only return a message, we don't need the existing even't sdata.
			c.JSON(400, gin.H{
				"message": "Event already exists.",
			})
		}
	}
}

/** Updates an existing event. All fields are optional. */
func Update(c *gin.Context) {
	var body models.Event
	var event models.Event
	var id = c.Param("id")

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid body.",
		})
	} else {
		models.DB.Find("id = ?", id).First(&event)

		var newEvent = models.Event{
			// Id:             event.Id,
			Title:          compareField(event.Title, body.Title),
			AuthorUsername: compareField(event.AuthorUsername, body.AuthorUsername),
			UseWhitelist:   true,
		}

		models.DB.Where("id = ?", id).Assign(newEvent).FirstOrCreate(&event)

		c.JSON(200, gin.H{
			"data":    event,
			"message": "Successfully updated event.",
		})
	}
}

func Delete(c *gin.Context) {
	var event models.Event
	var id = c.Param("id")

	models.DB.Where("id = ?", id).First(&event)

	if event.Id == "" {
		c.JSON(400, gin.H{
			"id":      id,
			"message": "Event with specified ID does not exist.",
		})
	} else {
		models.DB.Delete(&event)
		c.JSON(200, gin.H{
			"id":      event.Id,
			"message": "Successfully deleted event.",
		})
	}
}

// ----- helper functions -----

func compareField(eventField string, bodyField string) string {
	if bodyField == "" {
		return eventField
	} else {
		return bodyField
	}
}

func hasInvalidField(event models.Event) bool {
	return (event.AuthorUsername == "" ||
		event.Id == "" ||
		event.Title == "" ||
		len(event.Id) < 4)
}
