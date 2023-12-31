package main

import (
	"context"
	"flag"
	"log"

	"github.com/cblokkeel/hotel-reservation/api"
	"github.com/cblokkeel/hotel-reservation/constants"
	"github.com/cblokkeel/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27030"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client, db.DbName))

	listenAddr := flag.String("listenAddr", ":3000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New(config)
	apiv1 := app.Group(constants.ApiV1Route)

	apiv1.Get(constants.UserRoute, userHandler.HandleGetUsers)
	apiv1.Post(constants.UserRoute, userHandler.HandleInsertUser)
	apiv1.Patch(constants.UserByIdRoute, userHandler.HandleUpdateUser)
	apiv1.Delete(constants.UserByIdRoute, userHandler.HandleDeleteUser)
	apiv1.Get(constants.UserByIdRoute, userHandler.HandleGetUser)

	app.Listen(*listenAddr)
}
