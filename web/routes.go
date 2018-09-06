package web

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"strconv"
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

	/*
		Test marshalling struct and send with response writer
	*/
	app.Get("entities-second-writer-test", func(ctx iris.Context) {
		entities := internal.FindEntitiesSecond()
		content, err := json.Marshal(entities)
		if err != nil {
			ctx.StatusCode(500)
			ctx.JSON(iris.Map{"error": "Internal Server Error"})
		} else {
			ctx.ContentType("application/json")
			ctx.Write(content)
		}
	})

	app.Get("/entities-first", func(ctx iris.Context) {
		entities := internal.FindEntitiesFirst()
		var content []iris.Map
		for _, entity := range entities {
			content = append(content, iris.Map{"idOne": entity.IdOne, "IdTwo": entity.IdTwo, "value": entity.Value})
		}
		ctx.JSON(content)
	})

	app.Get("/entities-second/{idOne}/{idTwo}", func(ctx iris.Context) {
		idOne, err := strconv.Atoi(ctx.Params().Get("idOne"))
		if err != nil {
			fmt.Println(err)
		}
		idTwo, err := strconv.Atoi(ctx.Params().Get("idTwo"))
		if err != nil {
			fmt.Println(err)
		}
		entity := internal.GetEntitySecond(idOne, idTwo)
		ctx.JSON(iris.Map{"idOne": entity.IdOne, "idTwo": entity.IdTwo, "value": entity.Value})
	})

	app.Get("/entities-second/{idOne}/{idTwo}/nestedEntities", func(ctx iris.Context) {
		idOne, err := strconv.Atoi(ctx.Params().Get("idOne"))
		if err != nil {
			fmt.Println(err)
		}
		idTwo, err := strconv.Atoi(ctx.Params().Get("idTwo"))
		if err != nil {
			fmt.Println(err)
		}
		nestedEntities := internal.FindNestedEntitiesSecond(idOne, idTwo)
		var content []iris.Map
		for _, nestedEntity := range nestedEntities {
			content = append(content, iris.Map{"id": nestedEntity.IdThree, "valueOne": nestedEntity.ValueOne})
		}
		ctx.JSON(content)
	})

}
