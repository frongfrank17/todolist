package service

import (
	"errors"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"

	repository "Hugeman/repository"
)

type Required struct {
	repo repository.Interfaces
}

func NewTodoService(repo repository.Interfaces) Required {
	return Required{repo: repo}
}

var DataT = [5]string{"title", "description", "status", "image", "createdAt"}

func (rqeuir Required) Created(title string, description string, image string) error {
	Base64Image := image

	Input := &repository.Todo{
		ID:          uuid.NewV4(),
		Title:       title,
		Description: description,
		Status:      "IN_PROGRESS",
		Image:       Base64Image,
		CreatedAt:   time.Now().Local(),
	}
	fmt.Println(title, ": ", Input.Status)
	err := rqeuir.repo.Created(Input)
	if err != nil {
		return err
	}
	return nil
}
func (rqeuir Required) GetListsSort(f string, sort string) ([]Todo, error) {
	ChkField := false
	for _, value := range DataT {

		if value == f {
			ChkField = true
		}
	}
	if ChkField == false {
		return nil, errors.New("Error Sorted Required")
	}
	OrderBy := f + " "
	if sort == "0" {
		OrderBy = OrderBy + "ASC"
	} else {
		OrderBy = OrderBy + "DESC"
	}

	result, err := rqeuir.repo.GetListSort(OrderBy)
	if err != nil {
		return nil, err
	}
	//	fmt.Println(result)
	todoResps := []Todo{}
	for _, td := range result {

		todoResp := Todo{
			ID:          td.ID.String(),
			Title:       td.Title,
			Description: td.Description,
			Date:        td.CreatedAt,
			Status:      td.Status,
		}
		todoResps = append(todoResps, todoResp)
	}
	return todoResps, nil
}
func (rqeuir Required) GetListTitle(title string) (*Todo, error) {
	result, err := rqeuir.repo.GetListTitle(title)
	if err != nil {
		return nil, err
	}
	//uuid := hex.EncodeToString(result.ID)

	resp := &Todo{
		ID:          result.ID.String(),
		Title:       result.Title,
		Description: result.Description,
		Date:        result.CreatedAt,
	}
	return resp, nil
}
func (rqeuir Required) GetListDescription(description string) (*Todo, error) {
	result, err := rqeuir.repo.GetListDescription(description)
	if err != nil {
		return nil, err
	}
	resp := &Todo{
		ID:          result.ID.String(),
		Title:       result.Title,
		Description: result.Description,
	}
	return resp, nil

}
func (rqeuir Required) Updated(uid string, f string, input string) error {
	fmt.Println(uid, " : ", f, " : ", input)
	ChkField := false
	for _, value := range DataT {
		if value == f {
			ChkField = true
		}
	}
	if ChkField == false {
		return errors.New("Error Case Updated")
	}
	if f == "status" {
		fmt.Println(f)
		err := rqeuir.repo.Updated(uid, f, input)
		fmt.Println(err)
		if err != nil {
			return err
		}
		return nil
	}
	err := rqeuir.repo.Updated(uid, f, input)
	fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}
