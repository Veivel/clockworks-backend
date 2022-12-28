package models

type Event struct {
	Id             string `gorm:"primaryKey" json:"id"`
	Title          string `gorm:"type:VARCHAR(128)" json:"title"`
	AuthorUsername string `gorm:"type:VARCHAR(32)" json:"author_username"`
	UseWhitelist   bool   `gorm:"type:BOOLEAN" json:"use_whitelist"`
}

type EventData struct {
	Title          string `gorm:"type:VARCHAR(128)" json:"title"`
	AuthorUsername string `gorm:"type:VARCHAR(32)" json:"author_username"`
	UseWhitelist   bool   `gorm:"type:BOOLEAN" json:"use_whitelist"`
}
