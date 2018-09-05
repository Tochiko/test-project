package web

import (
	"github.com/kataras/iris"
	"test-project/internal/db"
)

func RegistryRoutes(app *iris.Application) {
	app.Handle("GET", "/", func(context iris.Context) {
		context.HTML("<h1>Welcome</h1>")
	})

	app.Get("/entities-second", func(ctx iris.Context) {
		entities := internal.FindEntitiesSecond()
		var content []iris.Map
		for _, entity := range entities {
			content = append(content, iris.Map{"idOne": entity.IdOne, "idTwo": entity.IdTwo, "value": entity.Value})
		}
		ctx.JSON(content)
	})

	app.Get("/entities-first", func(ctx iris.Context) {
		entities := internal.FindEntitiesFirst()
		var content []iris.Map
		for _, entity := range entities {
			content = append(content, iris.Map{"idOne": entity.IdOne, "IdTwo": entity.IdTwo, "value": entity.Value})
		}
		ctx.JSON(content)
	})
}
