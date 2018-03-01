package router

import(
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	//"fmt"
	//"bytes"
	//"encoding/json"
	"cardinal_api/model"
)


func LoadServiceRoutes(app *iris.Application, mongoConnection *mgo.Database) {
	app.Get("/services", func(ctx iris.Context) {
		mongoCollection := mongoConnection.C("app")
		results := []model.CardinalService{}
		err := mongoCollection.Find(bson.M{"name": "Lightspeed"}).All(&results)
		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(results)
	})

	app.Get("/services/:id/actions", func(ctx iris.Context) {

		mongoCollection := mongoConnection.C("app")
		result := model.CardinalService{}
		err := mongoCollection.Find(bson.M{"name": "Lightspeed"}).One(&result)
		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(result)
	})

	app.Post("/services/:id/actions/:action_id", func(ctx iris.Context) {
		c := &model.ForwardRequest{}
		ctx.ReadJSON(c)
		mongoCollection := mongoConnection.C("app")
		result := model.CardinalService{}
		err := mongoCollection.Find(bson.M{"name": "Lightspeed"}).One(&result)
		if err != nil {
			log.Fatal(err)
		}
		c.SendData()
		ctx.JSON(c)
	})

	app.Get("/services/:id/actions/:action_id", func(ctx iris.Context) {

		mongoCollection := mongoConnection.C("app")
		result := model.CardinalService{}
		err := mongoCollection.Find(bson.M{"name": "Lightspeed"}).One(&result)
		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(result)
	})
}