package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/iris"
	"test-project/internal/db"
	"test-project/web"
)

var db *gorm.DB
var err error
var app *iris.Application

func main() {
	db, _ = gorm.Open("postgres", getConnectionArguments())
	if err != nil {
		panic("failed to connect data base!")
	}
	defer db.Close()
	internal.InitDatabase(db)

	app = iris.New()
	web.RegistryRoutes(app)
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))

}

func getConnectionArguments() string {
	dbArgs := "host=ec2-46-137-117-43.eu-west-1.compute.amazonaws.com "
	dbArgs += "port=5432 "
	dbArgs += "user=znmqbotyalubwj "
	dbArgs += "dbname=d1japsoqvsc0e1 "
	dbArgs += "password=b38e7983d1c697894999432b3b5b6808a4af858eef28d9dc787f558a60fd455f"
	return dbArgs
}
