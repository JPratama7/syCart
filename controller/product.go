package controller

import (
	fiberHelper "github.com/JPratama7/util/fiber"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"syCart/helper"
	"syCart/model"
	"syCart/repository"
)

func FetchProducts(ctx *fiber.Ctx) (err error) {
	product, err := helper.ExtractLocal[repository.ProductRepository](ctx, "product")
	if err != nil {
		return
	}

	var products []model.Product

	switch {
	case ctx.Query("category", "") != "":
		category := ctx.Query("category", "")
		products, err = product.GetProductByCategoryName(ctx.Context(), category)
	default:
		products, err = product.FetchProduct(ctx.Context())
	}

	if err != nil {
		return fiberHelper.
			NewReturnData(fiber.StatusNotFound, false, "product not found", err.Error()).
			WriteResponseBody(ctx)
	}

	return fiberHelper.
		NewReturnData(fiber.StatusOK, true, "products fetched", products).
		WriteResponseBody(ctx)
}

func GetProduct(ctx *fiber.Ctx) (err error) {
	product, err := helper.ExtractLocal[repository.ProductRepository](ctx, "product")
	if err != nil {
		return
	}

	productId, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		return fiberHelper.
			NewReturnData(fiber.StatusNotFound, false, "product not found", err.Error()).
			WriteResponseBody(ctx)
	}

	data, err := product.GetProduct(ctx.Context(), productId)
	if err != nil {
		return fiberHelper.
			NewReturnData(fiber.StatusNotFound, false, "product not found", err.Error()).
			WriteResponseBody(ctx)
	}

	return fiberHelper.
		NewReturnData(fiber.StatusOK, true, "product fetched", data).
		WriteResponseBody(ctx)
}
