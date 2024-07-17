package main

import (
	"context"
	fiberUtil "github.com/JPratama7/util/fiber"
	"github.com/JPratama7/util/token/paseto"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"syCart/helper"
	"syCart/repository"
	"syCart/route"
	"time"
)

func main() {
	app := fiber.New(fiberUtil.GenerateConfig())

	app.Use(logger.New())

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	conn, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(os.Getenv("MONGODB_URI")),
	)
	if err != nil {
		panic(err)
	}

	if er := conn.Ping(context.Background(), readpref.Primary()); er != nil {
		panic(er)
	}

	defer log.Println(conn.Disconnect(context.Background()))

	token := paseto.NewPASETO(os.Getenv("PUBLIC_KEY"), os.Getenv("PRIVATE_KEY"), time.Duration(60)*time.Minute)

	middleware := helper.NewAuthMiddleware(&token, "payload")

	userRepo := repository.NewUserRepository(conn.Database("cart"), "users")
	cartRepo := repository.NewCartRepository(conn.Database("cart"), "carts")
	orderRepo := repository.NewOrderRepository(conn.Database("cart"), "orders")
	paymentRepo := repository.NewPaymentRepository(conn.Database("cart"), "payments")
	productRepo := repository.NewProductRepository(conn.Database("cart"), "products")

	app.Use(helper.InjectToContext("user", userRepo))
	app.Use(helper.InjectToContext("cart", cartRepo))
	app.Use(helper.InjectToContext("order", orderRepo))
	app.Use(helper.InjectToContext("payment", paymentRepo))
	app.Use(helper.InjectToContext("product", productRepo))
	app.Use(helper.InjectToContext("token", token))
	app.Use(helper.InjectToContext("auther", middleware))

	group := app.Group("/api/v1")

	route.Unauthorized(group)

	route.Authorized(group, middleware)

	log.Fatal(app.Listen(":5000"))
}
