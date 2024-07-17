package controller

import (
	fiberHelper "github.com/JPratama7/util/fiber"
	"github.com/JPratama7/util/token/paseto"
	"github.com/JPratama7/util/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"syCart/helper"
	"syCart/model"
	"syCart/model/request"
	"syCart/model/response"
	"syCart/repository"
)

func Register(ctx *fiber.Ctx) (err error) {
	user, err := helper.ExtractLocal[repository.UserRepository](ctx, "user")
	if err != nil {
		return err
	}

	data, err := helper.ParseValidate[request.Register](ctx)
	if err != nil {
		return err
	}

	if _, er := user.UsernameOrEmailExist(ctx.Context(), data.Username, data.Email); er == nil {
		return fiberHelper.
			NewReturnData(fiber.StatusConflict, false, "Username or email already exist", types.NilInt()).
			WriteResponseBody(ctx)
	}

	salt, err := helper.GenerateSalt(4)
	if err != nil {
		return err
	}

	hashedPassword := helper.HashPassword(data.Password, salt)

	err = user.CreateUser(ctx.Context(), &model.User{
		UserId:       primitive.ObjectID{},
		Username:     data.Username,
		Email:        data.Email,
		Salt:         salt,
		PasswordHash: hashedPassword,
		CreatedAt:    helper.NewDatetime(),
		UpdatedAt:    helper.NewDatetime(),
	})

	if err != nil {
		return
	}

	err = fiberHelper.
		NewReturnData(fiber.StatusCreated, true, "User created successfully", types.NilInt()).
		WriteResponseBody(ctx)

	return
}

func Login(ctx *fiber.Ctx) (err error) {
	user, err := helper.ExtractLocal[repository.UserRepository](ctx, "user")
	if err != nil {
		return err
	}

	token, err := helper.ExtractLocal[paseto.PASETO](ctx, "token")
	if err != nil {
		return err
	}

	req, err := helper.ParseValidate[request.Login](ctx)
	if err != nil {
		return err
	}

	if req.Email == "" && req.Username == "" {
		return fiberHelper.
			NewReturnData(fiber.StatusBadRequest, false, "Email or username is required", types.NilInt()).
			WriteResponseBody(ctx)
	}

	data, err := user.UsernameOrEmailExist(ctx.Context(), req.Username, req.Email)
	if err != nil {
		return err
	}

	if !helper.CheckPasswordHash(req.Password, data.Salt, data.PasswordHash) {
		return fiberHelper.
			NewReturnData(fiber.StatusUnauthorized, false, "Invalid Credentials", types.NilInt()).
			WriteResponseBody(ctx)
	}

	key, err := token.EncodeWithStruct("data", &model.Token{
		Username: data.Username,
		Email:    data.Email,
	})
	if err != nil {
		return err
	}

	err = fiberHelper.
		NewReturnData(fiber.StatusOK, true, "Login successful", response.Login{Token: key}).
		WriteResponseBody(ctx)

	return
}
