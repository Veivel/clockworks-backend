package models

type Event struct {
	Id             string `gorm:"PrimaryKey" json:"id"`
	Title          string `gorm:"type:VARCHAR(128)" json:"title"`
	UseWhitelist   bool   `gorm:"type:BOOLEAN" json:"use_whitelist"`
	AuthorUsername string `gorm:"ForeignKey:Username" json:"author_username"`
	// AuthorUsername string `gorm:"type:VARCHAR(32)" json:"author_username"`
}
