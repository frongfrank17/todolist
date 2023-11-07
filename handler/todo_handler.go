package handler

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	service "Hugeman/service"
)

type Required struct {
	svc service.Interfaces
}

func NewTodoHandler(svc service.Interfaces) Required {
	return Required{svc: svc}
}

type FormDataRequestBody struct {
	Title       string                `form:"title"`
	Description string                `form:"description"`
	Image       *multipart.FileHeader `form:"image"`
}

func (requir Required) Created(c *gin.Context) {
	var formdata FormDataRequestBody

	if err := c.ShouldBindWith(&formdata, binding.FormMultipart); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	tempDir := "./temp"
	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		//fmt.Println("Error creating temporary directory:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	imagePath := filepath.Join(tempDir, formdata.Image.Filename)
	out, err := os.Create(imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer func() {
		out.Close()
		// Remove the temporary file after closing it
		if err := os.Remove(imagePath); err != nil {
			fmt.Println("Error deleting temporary file:", err)
		}
	}()
	file, err := formdata.Image.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	data, err := ioutil.ReadFile(imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	base64String := base64.StdEncoding.EncodeToString(data)
	//fmt.Println(base64String)
	//requir.svc.Created(formdata.Title, formdata.Description, ""
	//fmt.Println(formdata.Title)
	err = requir.svc.Created(formdata.Title, formdata.Description, base64String)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Created Todo List ",
	})
	return
}
func (r Required) GetList(c *gin.Context) {
	type_ := c.Param("type")
	search := c.Param("search")

	if type_ == "title" {
		resp, err := r.svc.GetListTitle(search)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Todo List ",
			"list":    resp,
		})

		return
	} else if type_ == "description" {
		resp, err := r.svc.GetListDescription(search)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Todo List ",
			"list":    resp,
		})

		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "Todo List ",
		"list":    nil,
	})
	return
}
func (r Required) GetListsAll(c *gin.Context) {
	sorted := c.Query("sorted")
	f := c.Query("f")
	if sorted != "0" || sorted != "1" {
		sorted = "0"
	}
	fmt.Println(sorted)
	resp, err := r.svc.GetListsSort(f, sorted)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo List ",
		"list":    resp,
	})
	return
}
func (r Required) UpdateTodoList(c *gin.Context) {
	// :field/:set
	f := c.Param("field")
	s := c.Param("set")
	uid := c.Param("uid")
	err := r.svc.Updated(uid, f, s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": " Update Todo Lists ",
	})
	return
}

func (r Required) UpdateStatusCOMPLETED(c *gin.Context) {
	// :field/:set

	uid := c.Param("uid")
	err := r.svc.Updated(uid, "status", "COMPLETED")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": " Update Todo Lists ",
	})
	return
}
func (r Required) UpdateStatusIN_PROGRESS(c *gin.Context) {
	// :field/:set
	uid := c.Param("uid")
	err := r.svc.Updated(uid, "status", "IN_PROGRESS")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": " Update Todo Lists ",
	})
	return
}
func (r Required) UpdateImageFiles(c *gin.Context) {
	// :field/:set
	type FormDataImage struct {
		ID string `form:"id"`

		Image *multipart.FileHeader `form:"image"`
	}

	var formdata FormDataImage

	if err := c.ShouldBindWith(&formdata, binding.FormMultipart); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	tempDir := "./temp"
	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		//fmt.Println("Error creating temporary directory:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	imagePath := filepath.Join(tempDir, formdata.Image.Filename)
	out, err := os.Create(imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer func() {
		out.Close()
		// Remove the temporary file after closing it
		if err := os.Remove(imagePath); err != nil {
			fmt.Println("Error deleting temporary file:", err)
		}
	}()
	file, err := formdata.Image.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	data, err := ioutil.ReadFile(imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	base64String := base64.StdEncoding.EncodeToString(data)
	//fmt.Println(base64String)
	//requir.svc.Created(formdata.Title, formdata.Description, ""
	//fmt.Println(formdata.Title)
	err = r.svc.Updated(formdata.ID, "image", base64String)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Upload Image to  Todo List ",
	})
	return
}
