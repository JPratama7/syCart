package controller

import (
	fiberHelper "github.com/JPratama7/util/fiber"
	"github.com/JPratama7/util/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"syCart/helper"
	"syCart/model/request"
	"syCart/repository"
)

func AddToCart(ctx *fiber.Ctx) (err error) {
	auth, err := helper.ExtractLocal[*helper.AuthMiddleware](ctx, "auther")
	if err != nil {
		return
	}

	cart, err := helper.ExtractLocal[repository.CartRepository](ctx, "cart")
	if err != nil {
		return
	}

	user, err := helper.ExtractLocal[repository.UserRepository](ctx, "user")
	if err != nil {
		return
	}

	userToken, err := auth.Decode(ctx)
	if err != nil {
		return err
	}

	userData, err := user.GetByEmail(ctx.Context(), userToken.Email)
	if err != nil {
		return err
	}

	req, err := helper.ParseValidate[request.Cart](ctx)
	if err != nil {
		return fiberHelper.NewReturnData(fiber.StatusBadRequest, false, "Invalid Request", types.NilInt()).WriteResponseBody(ctx)
	}

	if req.Quantity <= 0 {
		return fiberHelper.NewReturnData(fiber.StatusBadRequest, false, "Invalid Quantity", types.NilInt()).WriteResponseBody(ctx)
	}

	productId, err := primitive.ObjectIDFromHex(req.ProductId)
	if err != nil {
		return fiberHelper.NewReturnData(fiber.StatusBadRequest, false, "Invalid Product Id", types.NilInt()).WriteResponseBody(ctx)
	}

	product, err := helper.ExtractLocal[repository.ProductRepository](ctx, "product")
	if err != nil {
		return
	}

	productData, err := product.GetProduct(ctx.Context(), productId)
	if err != nil {
		return
	}

	if productData.StockQuantity < req.Quantity {
		return fiberHelper.NewReturnData(fiber.StatusBadRequest, false, "Out of Stock", types.NilInt()).WriteResponseBody(ctx)
	}

	cartItem, err := cart.GetCartItem(ctx.Context(), &userData, productId)
	if err == nil {
		cartItem.Quantity += req.Quantity
		err = cart.UpdateCart(ctx.Context(), &userData, &cartItem)

		if err != nil {
			return fiberHelper.NewReturnData(fiber.StatusInternalServerError, false, "Internal Server Error", types.NilInt()).WriteResponseBody(ctx)
		}

		return fiberHelper.
			NewReturnData(fiber.StatusOK, true, "Cart Updated", types.NilInt()).
			WriteResponseBody(ctx)
	}

	id, err := cart.AddCart(ctx.Context(), &userData, productId, req.Quantity)
	if err != nil {
		return fiberHelper.NewReturnData(fiber.StatusInternalServerError, false, "Internal Server Error", types.NilInt()).WriteResponseBody(ctx)
	}

	return fiberHelper.
		NewReturnData(fiber.StatusOK, true, "Product Added to Cart", id).
		WriteResponseBody(ctx)
}

func DeleteCart(ctx *fiber.Ctx) (err error) {
	auth, err := helper.ExtractLocal[*helper.AuthMiddleware](ctx, "auther")
	if err != nil {
		return
	}

	cart, err := helper.ExtractLocal[repository.CartRepository](ctx, "cart")
	if err != nil {
		return
	}

	user, err := helper.ExtractLocal[repository.UserRepository](ctx, "user")
	if err != nil {
		return
	}

	userToken, err := auth.Decode(ctx)
	if err != nil {
		return
	}

	userData, err := user.GetByEmail(ctx.Context(), userToken.Email)
	if err != nil {
		return
	}

	req, err := helper.ParseValidate[request.CartItem](ctx)
	if err != nil {
		return fiberHelper.NewReturnData(fiber.StatusBadRequest, false, "Invalid Request", types.NilInt()).WriteResponseBody(ctx)
	}

	cartItemId, err := primitive.ObjectIDFromHex(req.CartItemId)

	err = cart.RemoveCart(ctx.Context(), &userData, cartItemId)
	if err != nil {
		return fiberHelper.NewReturnData(fiber.StatusInternalServerError, false, "Internal Server Error", types.NilInt()).WriteResponseBody(ctx)
	}
	return fiberHelper.
		NewReturnData(fiber.StatusOK, true, "Cart Item Removed", types.NilInt()).
		WriteResponseBody(ctx)

}

func GetCartItems(ctx *fiber.Ctx) (err error) {
	auth, err := helper.ExtractLocal[*helper.AuthMiddleware](ctx, "auther")
	if err != nil {
		return
	}

	cart, err := helper.ExtractLocal[repository.CartRepository](ctx, "cart")
	if err != nil {
		return
	}

	user, err := helper.ExtractLocal[repository.UserRepository](ctx, "user")
	if err != nil {
		return
	}

	userToken, err := auth.Decode(ctx)
	if err != nil {
		return
	}

	userData, err := user.GetByEmail(ctx.Context(), userToken.Email)
	if err != nil {
		return err
	}

	cartItems, err := cart.GetCarts(ctx.Context(), &userData)
	if err != nil || len(cartItems) == 0 {
		return fiberHelper.
			NewReturnData(fiber.StatusNotFound, false, "Cart Not Found", types.NilInt()).
			WriteResponseBody(ctx)
	}

	return fiberHelper.
		NewReturnData(fiber.StatusOK, true, "Cart Found", cartItems).
		WriteResponseBody(ctx)
}
