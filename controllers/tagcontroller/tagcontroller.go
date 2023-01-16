package tagcontroller

import (
	"clockworks-backend/models"
	"clockworks-backend/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// All tags for a single event
func Index(c *gin.Context) {
	var tags []models.Tag
	var eventId = c.Param("eventId")

	models.DB.Where("event_id = ?", eventId).Find(&tags)

	c.JSON(200, gin.H{
		"data": tags,
	})
}

func Toggle(c *gin.Context) {
	var body models.TagsData
	var event models.Event
	var eventId = c.Param("eventId")
	var username, _ = c.Get("Username")

	err := c.BindJSON(&body)

	if err != nil { // add another condition for invalid body
		c.JSON(400, gin.H{
			"message": "Invalid body.",
		})
		return
	} else {
		// should be a middleware?
		models.DB.Where("id = ?", eventId).Find(&event)
		if event.AuthorUsername != username.(string) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		} else {
			var count int
			// Toggles a tag for each time period in body
			for i := range body.Periods {
				var destTag models.Tag
				var tmpTag = models.Tag{
					EventId:  eventId,
					Username: username.(string),
					Period:   body.Periods[i],
					TagType:  0,
				}

				fmt.Println("tag: ", tmpTag)
				fmt.Println("tag's time:", utils.GetTime(tmpTag.Period))
				// TODO: validation of time ^

				// Checks if tmpTag exists or not...
				models.DB.Table("tags").First(&destTag, &tmpTag)
				if destTag.EventId == "" {
					// Creates tag (toggles "on") if doesnt exist yet...
					fmt.Printf("not exist. Creating...\n\n")
					models.DB.Table("tags").Create(&tmpTag)
					count++
				} else {
					// Deletes tag (toggles 'off') if already exists.
					fmt.Printf("already exist. Deleting...\n\n")
					models.DB.Delete(&destTag, &tmpTag)
					count++
				}

			}

			c.JSON(200, gin.H{
				"message": fmt.Sprintf("Succesfully toggled %d tags in event.", count),
			})
		}
	}
}
