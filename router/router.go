package router

import (
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2"
)



func Load(app *iris.Application, mongoConnection *mgo.Database) {
	LoadServiceRoutes(app, mongoConnection)
}