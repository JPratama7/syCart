package controller

import (
	fiberHelper "github.com/JPratama7/util/fiber"
	"github.com/gofiber/fiber/v2"
	"syCart/helper"
	"syCart/model"
	"syCart/repository"
)

func CheckoutCart(ctx *fiber.Ctx) (err error) {
	auther, err := helper.ExtractLocal[*helper.AuthMiddleware](ctx, "auther")
	if err != nil {
		return
	}

	userToken, err := auther.Decode(ctx)
	if err != nil {
		return
	}

	user, err := helper.ExtractLocal[repository.UserRepository](ctx, "user")
	if err != nil {
		return
	}

	userData, err := user.GetByEmail(ctx.Context(), userToken.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
	}

	cart, err := helper.ExtractLocal[repository.CartRepository](ctx, "cart")
	if err != nil {
		return
	}

	order, err := helper.ExtractLocal[repository.OrderRepository](ctx, "order")
	if err != nil {
		return
	}

	payment, err := helper.ExtractLocal[repository.PaymentRepository](ctx, "payment")
	if err != nil {
		return
	}

	listCart, err := cart.CartsWithProduct(ctx.Context(), &userData)
	if err != nil || len(listCart) == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Cart not found, create a new one")
	}

	orderData := new(model.Order)
	orderData.UserId = userData.UserId
	orderData.Status = model.PENDING

	orderItems := make([]model.OrderItem, 0, len(listCart))
	for _, v := range listCart {
		orderItems = append(orderItems, model.OrderItem{
			ProductId:       v.ProductId,
			Quantity:        v.Quantity,
			PriceAtPurchase: v.Product.Price,
		})
		orderData.TotalAmount += v.Product.Price * float64(v.Quantity)
	}

	orderData.OrderItems = &orderItems
	orderData.CreatedAt = helper.NewDatetime()
	orderData.UpdatedAt = helper.NewDatetime()

	orderId, err := order.Checkout(ctx.Context(), orderData)
	if err != nil {
		return
	}

	id, err := payment.CreatePayment(ctx.Context(), orderId, orderData)
	if err != nil {
		return
	}

	err = cart.RemoveAll(ctx.Context(), &userData)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
	}

	return fiberHelper.
		NewReturnData(fiber.StatusOK, true, "Successfully checkout", id).
		WriteResponseBody(ctx)
}
