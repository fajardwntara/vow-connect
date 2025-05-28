package user

import (
	"context"
	"time"
)

/* ======= User Entity ======= */
type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetAll(ctx context.Context) ([]User, error)
}

/* ======= Organizer Entity ======= */
type Organizer struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null" json:"password"`
	Phone     string    `gorm:"size:20" json:"phone"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type OrganizerRepository interface {
	Create(org *Organizer) error
	Update(org *Organizer) error
	Delete(id uint) error
	GetByID(id uint) (*Organizer, error)
	GetByEmail(email string) (*Organizer, error)
	GetAll() ([]Organizer, error)
}

/* ======= Guest Entity ======= */
type Guest struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	WeddingID  uint32    `gorm:"not null" json:"wedding_id"`
	Name       string    `gorm:"size:255;not null" json:"name"`
	Phone      string    `gorm:"size:20" json:"phone"`
	RSVPStatus string    `gorm:"size:20;default:'pending'" json:"rsvp_status"`
	Message    string    `gorm:"type:text" json:"message"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type GuestRepository interface {
	Create(guest *Guest) error
	Update(guest *Guest) error
	Delete(id uint) error
	GetByID(id uint) (*Guest, error)
	GetAll() ([]Guest, error)
}
