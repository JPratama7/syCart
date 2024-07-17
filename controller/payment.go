package controller

import (
	fiberHelper "github.com/JPratama7/util/fiber"
	"github.com/gofiber/fiber/v2"
	"syCart/helper"
	"syCart/repository"
)

func CheckPayment(ctx *fiber.Ctx) (err error) {
	auth, err := helper.ExtractLocal[*helper.AuthMiddleware](ctx, "auther")
	if err != nil {
		return
	}

	userToken, err := auth.Decode(ctx)
	if err != nil {
		return
	}

	user, err := helper.ExtractLocal[repository.UserRepository](ctx, "user")
	if err != nil {
		return
	}

	payment, err := helper.ExtractLocal[repository.PaymentRepository](ctx, "payment")
	if err != nil {
		return err
	}

	userData, err := user.GetByEmail(ctx.Context(), userToken.Email)
	if err != nil {
		return
	}

	payments, err := payment.FetchPaymentByUser(ctx.Context(), &userData)
	if err != nil || len(payments) == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Payment Not Found")
	}

	return fiberHelper.NewReturnData(fiber.StatusOK, true, "Payment Found", payments).WriteResponseBody(ctx)
}
