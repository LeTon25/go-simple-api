package initalize

import (
	"fmt"
	"go-simple-api/global"
)

func RunProgram() {
	LoadConfig()
	fmt.Println("Configuration loaded successfully", global.Config.MySQL.Host)
	InitLogger()
	InitMySQL()
	InitRedis()
	router := InitRouter()
	router.Run()
}
