package models

type Tag struct {
	// Id      string 	 `gorm:"primaryKey" json:"id"`
	EventId  string `gorm:"ForeignKey:Id" json:"event_id"`
	Username string `gorm:"type:VARCHAR(32)" json:"username"`
	Period   string `json:"period"`
	TagType  int8   `gorm:"type:INT(8)" json:"tag_type"`
}

type TagsData struct {
	Periods []string `gorm:"type:VARCHAR(8)" json:"periods"`
}
