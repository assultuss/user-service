package db

type User struct {
	ID            string `gorm:"primaryKey;column:id" json:"ID"`
	FirstName     string `gorm:"column:first_name" json:"first_name"`
	LastName      string `gorm:"column:last_name" json:"last_name"`
	Age           int    `gorm:"column:age" json:"age"`
	RecordingDate int64  `gorm:"column:recording_date" json:"recording_date"`
}
