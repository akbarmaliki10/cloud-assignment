package route

import (
	"amitshekar-clean-arch/api/controller"
	"amitshekar-clean-arch/bootstrap"
	repository "amitshekar-clean-arch/repository/mysql"
	"amitshekar-clean-arch/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewProductRouter(env *bootstrap.Env, db *gorm.DB, group *gin.RouterGroup) {
	pr := repository.NewMysqlProductRepo(db)
	tc := controller.ProductController{
		ProductUsecase: usecase.NewProductUsecase(pr),
	}
	group.POST("/", tc.CreateProduct)
	group.GET("/:id", tc.GetProduct)
	group.GET("/", tc.GetAll)
	group.PUT("/:id", tc.UpdateProduct)
	group.DELETE("/:id", tc.DeleteProduct)
}
