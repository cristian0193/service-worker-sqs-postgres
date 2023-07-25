package entity

// Events represents the entity.
type Events struct {
	ID      string `gorm:"NULL;TYPE:VARCHAR(200);COLUMN:id" json:"id"`
	Message string `gorm:"NULL;TYPE:VARCHAR(200);COLUMN:message" json:"message"`
	Date    string `gorm:"NULL;TYPE:VARCHAR(200);COLUMN:date" json:"date"`
}

// TableName definition name for table .
func (Events) TableName() string {
	return "events"
}
