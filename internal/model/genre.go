package model

// import libraries
import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Genre struct contains ID that is guid and Genre name.
type Genre struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primary_key;type:uuid"`
	GenreName string    `json:"genrename"`
}
