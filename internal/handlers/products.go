package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kwesikwaa/toyshop-backend/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Toy struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt   time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updated_at"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Price       float64            `json:"price" bson:"price"`
	Favourites  int                `json:"favourites" bson:"favourites"`
}

func CreateProduct(c *fiber.Ctx) error {
	toy := Toy{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := c.BodyParser(&toy); err != nil {
		return err
	}

	client, err := db.GetMongoClient()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("we no fit connect mongo")
	}

	collection := client.Database(db.DatabaseName).Collection(string(db.ToysCollection))

	_, err = collection.InsertOne(context.TODO(), toy)
	if err != nil {
		return err
	}

	return c.JSON(collection)
}

func GetAllProducts(c *fiber.Ctx) error {
	client, err := db.GetMongoClient()
	var toys []*Toy

	if err != nil {
		return err
	}

	collection := client.Database(db.DatabaseName).Collection(string(db.ToysCollection))

	cur, err := collection.Find(context.TODO(), bson.D{primitive.E{}})

	if err != nil {
		return err
	}
	for cur.Next(context.TODO()) {
		var t Toy
		err := cur.Decode(&t)

		if err != nil {
			return err
		}
		toys = append(toys, &t)
	}

	return c.JSON(toys)

}

func GetSingleProduct(c *fiber.Ctx) {

}

func UpdateProduct(c *fiber.Ctx) {

}
