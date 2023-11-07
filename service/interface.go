package service

import "time"

type Todo struct {
	ID          string    `json:"ID"`
	Title       string    `json:"Title"`
	Description string    `json:"Description"`
	Date        time.Time `json:"Date"`
	Status      string    `json:Status`
}
type Interfaces interface {
	Created(title string, description string, image string) error
	GetListsSort(f string, sort string) ([]Todo, error)
	GetListTitle(title string) (*Todo, error)
	GetListDescription(description string) (*Todo, error)
	Updated(uid string, f string, input string) error
}
