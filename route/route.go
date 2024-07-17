package route

import (
	"github.com/gofiber/fiber/v2"
	"syCart/controller"
	"syCart/helper"
)

func Unauthorized(app fiber.Router) {
	app.Post("/login", controller.Login)
	app.Post("/register", controller.Register)

	app.Get("/products", controller.FetchProducts)
	app.Get("/products/:id", controller.GetProduct)
}

func Authorized(app fiber.Router, middleware *helper.AuthMiddleware) {
	group := app.Group("/")
	group.Use(middleware.Authenticate)

	cart := group.Group("/cart")

	cart.Get("/", controller.GetCartItems)
	cart.Post("/", controller.AddToCart)
	cart.Delete("/", controller.DeleteCart)

	group.Post("/checkout", controller.CheckoutCart)
	group.Get("/payments", controller.CheckPayment)
}
