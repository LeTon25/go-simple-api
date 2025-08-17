package initalize

import (
	"fmt"
	"go-simple-api/global"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, message string) {
	if err != nil {
		global.Logger.Error(message, zap.Error(err))
		panic(err)
	}

}

func InitMySQL() {
	m := global.Config.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.User, m.Password, m.Host, m.Port, m.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	checkErrorPanic(err, "Failed to connect to MySQL")
	global.Mdb = db

	SetPool()
	migrateTables()
}

func SetPool() {
	m := global.Config.MySQL
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		global.Logger.Error("Failed to get SQL DB", zap.Error(err))
		panic(err)
	}

	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)

}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
	//&model.GoUserV2
	)

	if err != nil {
		return
	}
}

func genTableDAO() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(global.Mdb) // reuse your gorm db

	g.GenerateAllTable()

	// Generate the code
	g.Execute()
}
