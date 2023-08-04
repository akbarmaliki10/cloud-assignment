package route

import (
	"amitshekar-clean-arch/api/controller"
	"amitshekar-clean-arch/bootstrap"
	repository "amitshekar-clean-arch/repository/mysql"
	"amitshekar-clean-arch/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewTodoRouter(env *bootstrap.Env, db *gorm.DB, group *gin.RouterGroup) {
	tr := repository.NewMysqlTodoRepo(db)
	tc := controller.TodoController{
		TodoUsecase: usecase.NewTodoUsecase(tr),
	}
	group.POST("/create", tc.CreateTodo)
	group.GET("/get/:name", tc.GetTodo)
	group.GET("/get", tc.GetAll)
	group.POST("/update", tc.UpdateTodo)
	group.DELETE("/delete/:name", tc.DeleteTodo)
}
