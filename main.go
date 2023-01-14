package main

import (
	"log"
	"time"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db	*mongo.Database
}

var mg MongoInstance

const dbName = "hrms-api"
const mongoURI = 

type Employee struct {
	ID     string
	Name   string
	Salary float64
	Age    float64
}

func Connect() error {
    //serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	client, err := mongo.NewClient.(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := client.Connect(ctx, clientOptions)
	db := client.Database(dbName)
	if err != nil {
		return err
	}
	mg = MongoInstance{
		Client: client,
		Db: db
	}
	return nil
}

func main() {

	if err:= Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/employee", func(c *fiber.Ctx) error {
		return c.Status(200).
	})
	app.Post("/employee", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World !")
	})
	app.Patch("/employee/:id", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World !")
	})
	app.Delete("/employee/:id", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World !")
	})

	app.Listen(":3000")
}
