package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/iris"
	"os"
	"test-project/internal/db"
	"test-project/web"
)

var db *gorm.DB
var err error
var app *iris.Application

func main() {
	db, _ = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("failed to connect data base!")
	}
	defer db.Close()
	internal.InitDatabase(db)

	app = iris.New()
	web.RegistryRoutes(app)
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))

}
