package router

import (
	"GoLearn/accountBank/config"
	"GoLearn/accountBank/controller"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"os"
	"os/exec"
)

func SetupRouter(config config.Config, r *gin.Engine, c *controller.Controller)  {
	// Moi route deu bat dau boi prefix ServiceName
	api := r.Group("/" + config.ServiceName)
	{

		setupAdminRoutes(c, api)
		setupDocumentRoute(api, "")
	}
}

func setupDocumentRoute(api *gin.RouterGroup, ginMode string) {
	if ginMode != "release" {
		goPath := os.ExpandEnv("$GOPATH")
		if goPath == "" {
			panic("Chưa thiết lập GOPATH!")
		}

		_, err := exec.Command("/"+goPath+"/bin/swag", "init").Output()

		if err != nil {
			log.Println("Không thể tạo document bằng Swag!")
			panic(err)
		}

		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

