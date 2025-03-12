package initialize

import (
	"fmt"

	"github.com/anonystick/go-ecommerce-backend-api/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() *gin.Engine {
	// load configuration
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration nysql", m.Username, m.Password)
	InitLogger()
	global.Logger.Info("Config Log ok!!", zap.String("ok", "success"))
	InitMysql()
	InitMysqlC()
	InitServiceInterface()
	InitRedis()
	InitKafka()
	r := InitRouter()
	return r
	// r.Run(":8002")
}
