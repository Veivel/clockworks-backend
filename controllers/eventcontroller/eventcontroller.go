package eventcontroller

import (
	"clockworks-backend/models"
	"clockworks-backend/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/** Indexes ALL events. */
func All(c *gin.Context) {
	var events []models.Event

	models.DB.Find(&events)
	c.JSON(200, gin.H{
		"data": events,
	})
}

/** Index all of current user's events. No user matching. */
func Index(c *gin.Context) {
	var events []models.Event
	var user, _ = c.Get("Username")

	models.DB.Where("author_username = ?", user).Find(&events)
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

/** Return a event with a specified ID. No user matching. */
func Show(c *gin.Context) {
	var event models.Event
	var id = c.Param("id")

	result := models.DB.Where("id = ?", id).First(&event)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"id":      id,
			"message": "No Event with specified ID found.",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"data": event,
		})
	}
}

/*
* Creates a new Event, given event data.
Handles duplicate ID. No user matching.
*/
func Create(c *gin.Context) {
	var event models.Event
	var body models.Event

	err := c.BindJSON(&body)

	if err != nil || utils.HasInvalidField(body) {
		fmt.Println("error creating. body: ", body)
		c.JSON(400, gin.H{
			"message": "Invalid body.",
		})
	} else {
		// will not be blank. taken care in jwt middleware
		var username, _ = c.Get("Username")

		// structs are DEEP-COPIED by default. wowzers???
		event = body
		event.AuthorUsername = username.(string)

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

/*
* Updates an existing event.
All fields are optional. Uses user matching.
*/
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

		// if wrong user
		username, _ := c.Get("Username")
		if event.AuthorUsername != username.(string) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var newEvent = models.Event{
			Id:             event.Id,
			Title:          utils.GetUpdatedField(event.Title, body.Title), // can change
			AuthorUsername: event.AuthorUsername,
			UseWhitelist:   body.UseWhitelist, // can change
		}

		models.DB.Where("id = ?", id).Assign(newEvent).FirstOrCreate(&event)

		c.JSON(200, gin.H{
			"data":    event,
			"message": "Successfully updated event.",
		})
	}
}

/*
* Delete an Event with specified ID.
Performs user matching.
*/
func Delete(c *gin.Context) {
	var event models.Event
	var id = c.Param("id")

	models.DB.Where("id = ?", id).First(&event)
	username, _ := c.Get("Username")

	if event.Id == "" {
		c.JSON(400, gin.H{
			"id":      id,
			"message": "Event with specified ID does not exist.",
		})
	} else if event.AuthorUsername != username {
		// wrong user
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	} else {
		models.DB.Delete(&event)
		c.JSON(200, gin.H{
			"id":      event.Id,
			"message": "Successfully deleted event.",
		})
	}
}
