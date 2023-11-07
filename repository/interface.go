package repository

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primary_key;json:"id";type:uuid;default:uuid_generate_v4()"`
	Title       string    `json:"title"gorm:"type:text;size:100"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
}
type Interfaces interface {
	Created(todo *Todo) error
	GetLists() ([]Todo, error)
	GetListSort(sort string) ([]Todo, error)
	GetListTitle(title string) (*Todo, error)
	GetListDescription(description string) (*Todo, error)
	Updated(uid string, f string, input string) error
}
