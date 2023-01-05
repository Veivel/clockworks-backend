package utils

import "clockworks-backend/models"

/** Return eventField if bodyField is empty. Return bodyField otherwise. */
func CheckForEmptyField(eventField string, bodyField string) string {
	if bodyField == "" {
		return eventField
	} else {
		return bodyField
	}
}

/** Returns true if there are any invalid fields in event. */
func HasInvalidField(event models.Event) bool {
	return (event.AuthorUsername == "" ||
		event.Id == "" ||
		event.Title == "" ||
		len(event.Id) < 4)
}

/*
* Returns true if eventUser is the same user as bodyUser
 */
func isUserEquals(eventUser models.User, bodyUser models.User) bool {
	return eventUser.Username == bodyUser.Username &&
		eventUser.CreatedAt == bodyUser.CreatedAt &&
		eventUser.Email == bodyUser.Email
}

/*
* Retrieve a User instance from a given username string.
* Returns user with username "" if doesn't exist.
 */
func GetUser(username string) (user models.User) {
	models.DB.Table("users").Where("username = ?", username).First(&user)
	return
}
