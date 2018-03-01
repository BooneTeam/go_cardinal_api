package main

import (
	"gopkg.in/mgo.v2"
	"github.com/kataras/iris"
	"cardinal_api/router"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/rs/cors"
)


func main() {

	app := iris.New()


	app.Logger().SetLevel("debug")

	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}

	corsWrapper := cors.New(corsOptions).ServeHTTP

	app.WrapRouter(corsWrapper)

	session, err := mgo.Dial("localhost")

	if nil != err {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// DB Name and Collection Name
	mongoConnection := session.DB("cardinal-db")

	//mongoCollection := mongoConnection.C("app")
	//mongoCollection .Insert(&CardinalService{"Lightspeed"})

	router.Load(app, mongoConnection)

	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}