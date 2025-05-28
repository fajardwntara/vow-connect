package wedding

import (
	"time"

	"github.com/fajardwntara/vow-connect/api/domain/user"
)

type Wedding struct {
	ID          uint32       `gorm:"primary_key;auto_increment" json:"id"`
	OrganizerID uint32       `gorm:"not null" json:"organizer_id"`
	Name        string       `gorm:"size:255;not null" json:"name"`
	Date        time.Time    `json:"date"`
	Location    string       `gorm:"size:255" json:"location"`
	Guests      []user.Guest `gorm:"foreignKey:WeddingID" json:"guests"`
	CreatedAt   time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type WeddingRepository interface {
	Create(wd *Wedding) error
	Update(wd *Wedding) error
	Delete(id uint) error
}
