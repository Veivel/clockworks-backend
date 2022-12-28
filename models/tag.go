package models

type Tag struct {
	// Id         string `gorm:"primaryKey" json:"id"`
	EventId    string `gorm:"type:foreignKey" json:"event_id"`
	Username   string `gorm:"type:VARCHAR(32)" json:"username"`
	TimePeriod string `gorm:"type:TIME" json:"time_period"`
	TagType    int8   `gorm:"type:INT(8)" json:"tag_type"`
}
