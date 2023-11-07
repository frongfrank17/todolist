package repository

import (
	"gorm.io/gorm"
)

type Required struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) Required {
	return Required{db: db}
}

func (requir Required) Created(todo *Todo) error {
	err := requir.db.Create(&todo).Error
	if err != nil {
		return err
	}
	return nil
}
func (requir Required) GetLists() ([]Todo, error) {
	var todo []Todo
	err := requir.db.Find(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (requir Required) GetListSort(sort string) ([]Todo, error) {
	var todo []Todo
	err := requir.db.Order(sort).Find(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func (requir Required) GetListTitle(title string) (*Todo, error) {
	var todo *Todo
	err := requir.db.Where("title = ?", title).First(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func (requir Required) GetListDescription(description string) (*Todo, error) {
	var todo *Todo
	err := requir.db.Where("description = ?", description).First(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func (requir Required) Updated(uid string, f string, input string) error {
	// uid_, err := uuid.Parse(uid)
	// if err != nil {
	// 	return nil
	// }
	err := requir.db.Model(&Todo{}).Where("id = ?", uid).Update(f, input).Error
	if err != nil {
		return err
	}
	return nil
}
