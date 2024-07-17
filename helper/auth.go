package helper

import (
	"github.com/JPratama7/util/token/paseto"
	"github.com/gofiber/fiber/v2"
	"strings"
	"syCart/model"
)

type AuthMiddleware struct {
	key    string
	paseto *paseto.PASETO
}

func NewAuthMiddleware(paseto *paseto.PASETO, key string) *AuthMiddleware {
	return &AuthMiddleware{paseto: paseto, key: key}
}

func (a AuthMiddleware) Decode(ctx *fiber.Ctx) (data model.Token, err error) {
	payload := ctx.Locals(a.key)
	if payload == nil {
		err = ctx.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		return
	}

	data, ok := payload.(model.Token)
	if !ok {
		err = ctx.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		return
	}

	return
}

func (a AuthMiddleware) Authenticate(ctx *fiber.Ctx) error {
	token := strings.Split(ctx.Get("Authorization"), " ")
	if len(token) != 2 {
		return ctx.Status(fiber.StatusUnauthorized).SendString("Unauthorized, No token found")
	}

	payload := new(paseto.Payload[model.Token])

	err := a.paseto.Decode(token[1], payload)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).SendString("Unauthorized, No token found")
	}

	ctx.Locals(a.key, payload.Data)

	return ctx.Next()
}
