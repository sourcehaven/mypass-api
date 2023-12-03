package routes

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sourcehaven/mypass-api/internal/crypt"
	"github.com/sourcehaven/mypass-api/internal/db"
	"github.com/sourcehaven/mypass-api/internal/model"
	"net/http"
	"time"
)

// RegisterUser sends back HTTP 418 I'm a teapot status.
// @Summary I'm a teapot
// @Description Responds with HTTP status 418 I'm a teapot.
// @Tags examples
// @Produce string
// @Success 201 {string} string "I am a teapot!"
// @Router /teapot [get]
func (h *Handler) RegisterUser(c *fiber.Ctx) error {
	user := &db.User{}

	if err := ParseBody(c, user); err != nil {
		return err
	}

	if err := user.Validate(h.Cfg.PasswordLength, h.Cfg.PasswordMinNumber, h.Cfg.PasswordMinCapital, h.Cfg.PasswordMinSpecial); err != nil {
		return err
	}

	hashedPw, err := crypt.HashPassword(user.Password)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "Unable to hash password", nil)
		return err
	}

	user.PasswordHash = hashedPw
	data, err := h.Store.CreateUser(user)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return err
	}

	SendResponse(c, http.StatusCreated, "User created successfully!",
		&fiber.Map{"id": data.ID, "createdAt": data.CreatedAt})
	return nil
}

func (h *Handler) LoginUser(c *fiber.Ctx) error {
	cred := &model.Credentials{}
	if err := ParseBody(c, cred); err != nil {
		return err
	}

	user, err := h.Store.GetUser(cred.Username)

	if err != nil {
		if errors.Is(err, db.NotFound) {
			SendResponse(c, http.StatusNotFound, "Invalid credentials", nil)
		} else {
			SendResponse(c, http.StatusInternalServerError, err.Error(), nil)
		}
		return err
	}

	if user.Username != cred.Username || !crypt.IsValidPassword(cred.Password, user.PasswordHash) {
		SendResponse(c, http.StatusUnauthorized, "Invalid credentials", nil)
		return errors.New("invalid credentials")
	}

	now := time.Now()
	// Create the Claims
	claims := jwt.MapClaims{
		// "aud": "user",                 // The audience of the token
		"iss": h.Cfg.Host,                // The issuer of the token (domain name/IP/service ID)
		"sub": user.ID,                   // The subject of the token
		"exp": now.Add(time.Hour).Unix(), // Expiration
		"iat": now.Unix(),                // Issued at time
		"jti": uuid.New().String(),       // JWT ID
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(h.Cfg.SecretKey))
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, "Could not generate token", nil)
		return err
	}

	SendResponse(c, http.StatusOK, "Token created successfully", model.ResponseMap{"token": t})
	return nil
}
