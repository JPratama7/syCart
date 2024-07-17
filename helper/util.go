package helper

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/JPratama7/util/convert"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"lukechampine.com/blake3"
	"time"
)

// NewTimestamp creates a new timestamp.
func NewTimestamp() primitive.Timestamp {
	return primitive.Timestamp{T: uint32(time.Now().Unix())}
}

// NewDatetime creates a new datetime.
func NewDatetime() primitive.DateTime {
	return primitive.NewDateTimeFromTime(time.Now())
}

// ExtractLocal extracts a value from the context's locals.
func ExtractLocal[T any](ctx *fiber.Ctx, key string) (res T, err error) {
	value := ctx.Locals(key)
	if value == nil {
		return
	}

	res, ok := value.(T)
	if !ok {
		err = errors.New("invalid type")
		return
	}
	return
}

// validate is a global validator instance.
var validate = validator.New(validator.WithRequiredStructEnabled())

// ParseValidate parses the request body into the provided data structure.
func ParseValidate[T any](c *fiber.Ctx) (data T, err error) {
	err = c.BodyParser(&data)
	if err != nil {
		return
	}

	err = validate.Struct(data)
	if err != nil {
		return
	}
	return
}

// GenerateSalt creates a new random salt of the given length.
func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

// HashPassword generates a BLAKE3 hash of the password using the provided salt.
func HashPassword(password, salt string) string {
	// Combine the password and salt
	combined := password + salt

	// Generate BLAKE3 hash
	hash := blake3.Sum256(convert.UnsafeBytes(combined))

	// Return the hash as a hexadecimal string
	return hex.EncodeToString(hash[:])
}

// CheckPasswordHash compares a BLAKE3 hashed password with its possible plaintext equivalent.
func CheckPasswordHash(password, salt, hash string) bool {
	// Hash the password with the salt
	hashedPassword := HashPassword(password, salt)

	// Compare the hashed password with the provided hash
	return hashedPassword == hash
}

func InjectToContext[T any](key string, data T) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		ctx.Locals(key, data)
		return ctx.Next()
	}

}
