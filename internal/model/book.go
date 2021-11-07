package model

// import libraries
import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Book struct that contains ID, Name, Number, Title, Author and Genre
type Book struct {
	gorm.Model
	ID         uuid.UUID `gorm:"primary_key;type:uuid"`
	Owned      bool      `json:"owned"`
	BookName   string    `json:"name"`
	BookNumber string    `json:"number"`
	Title      string    `json:"title"`
	AuthorID   uuid.UUID `gorm:"type:uuid"`
	Author     Author    `gorm:"foreignKy:AuthorID" json:"author"`
	GenreID    uuid.UUID `gorm:"type:uuid"`
	Genre      Genre     `gorm:"foreigney:GenreID" json:"genre"`
}
