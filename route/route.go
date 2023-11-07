package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	handler "Hugeman/handler"
	repository "Hugeman/repository"
	service "Hugeman/service"
)

func ApirountingInterfaces(route *gin.Engine, db *gorm.DB) {
	repo := repository.NewTodoRepository(db)
	svc := service.NewTodoService(repo)
	hler := handler.NewTodoHandler(svc)
	apiGroup := route.Group("/api/todo")
	apiGroup.GET("/lists", hler.GetListsAll)
	apiGroup.GET("/list/:type/:search", hler.GetList)
	apiGroup.POST("/created", hler.Created)
	apiGroup.PUT("/list/update/:field/:set/:uid", hler.UpdateTodoList)
	apiGroup.PUT("/list/update_status/inprogress/:uid", hler.UpdateStatusIN_PROGRESS)
	apiGroup.PUT("/list/update_status/completed/:uid", hler.UpdateStatusCOMPLETED)
	apiGroup.PATCH("/list/update/image", hler.UpdateTodoList)

}
