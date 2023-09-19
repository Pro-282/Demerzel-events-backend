package models

type User struct {
	// Add user fields

	// Events Relationship
	Events           []Event `gorm:"foreignKey:Creator"`
	InterestedEvents []Event `gorm:"many2many:interested_events;"`
	UserGroup []UserGroup `json:"user_group" gorm:"foreignkey:UserID;association_foreignkey:ID"`
}
