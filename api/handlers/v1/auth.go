package v1

import (
	"context"
	"tender/api/helper/hashing"
	"tender/api/models"
	"tender/api/tokens"
	"tender/storage/repo"

	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
)

// @Summary 	  Login
// @Description   This Api for login users login with email and username
// @Tags 		  Register
// @Accept 		  json
// @Produce 	  json
// @Param 		  login body models.LoginRequest true "LoginRequest"
// @Success 	  200 {object} models.LoginResponse
// @Success 	  400 {object} models.Error
// @Failure 	  404 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /login [POST]
func (h *handlerV1) Login(ctx *gin.Context) {
	var (
		body        models.LoginRequest
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	// Bind request body to LoginRequest model
	err := ctx.ShouldBindJSON(&body)
	if err != nil || body.Username == "" || body.Password == "" {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Username and password are required",
		})
		log.Println("Login failed - missing required fields:", err)
		return
	}

	duration, err := time.ParseDuration(h.cfg.CtxTimeout)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("Failed to parse timeout", err)
		return
	}

	ctxTime, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	// Authenticate user
	response, err := h.storage.Auth().Login(ctxTime, body.Username) // Change to use Username
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.Error{
			Message: "User not found",
		})
		log.Println("Invalid login or password", err)
		return
	}

	// Check password
	check := hashing.CheckPasswordHash(body.Password, response.Password)
	if !check {
		ctx.JSON(http.StatusUnauthorized, models.Error{
			Message: "Invalid username or password",
		})
		log.Println("Failed to check password in login", err)
		return
	}

	// Generate JWT
	h.jwtHandler = tokens.JWTHandler{
		Sub:       response.Email, // Use Email for JWT
		Iat:       cast.ToString(time.Now().Unix()),
		Role:      response.Role,
		SigninKey: h.cfg.SigningKey,
		Timeout:   cast.ToInt(h.cfg.AccessTokenTimeout),
	}

	accessToken, refreshToken, err := h.jwtHandler.GenerateAuthJWT()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("Failed to generate token", err)
		return
	}

	// Respond with user info and JWT tokens
	ctx.JSON(http.StatusOK, models.LoginResponse{
		ID:       response.ID,
		Username: response.Username,
		Role:     response.Role,
		Email:    response.Email,
		Token:    accessToken,
		Refresh:  refreshToken,
	})
}

// @Summary 	  Register
// @Description   This Api for sign
// @Tags 		  Register
// @Accept 		  json
// @Produce 	  json
// @Param 		  signup body models.Register true "Register"
// @Success 	  201 {object} models.AlertMessage // Changed to 201 for successful creation
// @Success 	  400 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /register [POST]
func (h *handlerV1) Register(ctx *gin.Context) {
	var (
		body        models.Register
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := ctx.ShouldBindJSON(&body)
	if err != nil || body.Email == "" || body.Password == "" || body.Username == "" {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "username or email cannot be empty",
		})
		log.Println("Register failed - missing required fields:", err)
		return
	}

	duration, err := time.ParseDuration(h.cfg.CtxTimeout)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("Failed to parse timeout", err)
		return
	}

	ctxTime, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	// Validate email
	if err := body.ValidateEmail(); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "invalid email format",
		})
		log.Println("Error Incorrect email for validation:", err)
		return
	}

	// Validate password
	// if err := body.ValidatePassword(); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, models.Error{
	// 		Message: "Incorrect password for validation",
	// 	})
	// 	log.Println("Error Incorrect password for validation:", err)
	// 	return
	// }

	if err := body.ValidateRole(); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "invalid role",
		})
		return
	}

	// Check if email already exists
	responseEmail, err := h.storage.User().CheckField(ctx, "email", body.Email)
	if err != nil || responseEmail {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Email already exists",
		})
		return
	}

	// Check if username already exists
	responseUsername, err := h.storage.User().CheckField(ctx, "username", body.Username)
	if err != nil || responseUsername {
		ctx.JSON(http.StatusConflict, models.Error{
			Message: "Username already exists",
		})
		log.Println("Username already exists:", body.Username)
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "while making id error",
		})
		log.Println("Error Making UUID:", err)
		return
	}

	hashPassword, err := hashing.HashPassword(body.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: "Error while hashing password",
		})
		log.Println("Error while hashing password:", err)
		return
	}

	// Authenticate user
	response, err := h.storage.Auth().Register(ctxTime, &repo.User{
		ID:       id.String(),
		Username: body.Username,
		Password: hashPassword,
		Role:     body.Role,
		Email:    body.Email,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("Error during registration:", err)
		return
	}

	// Generate JWT
	h.jwtHandler = tokens.JWTHandler{
		Sub:       response.Email, // Use Email for JWT
		Iat:       cast.ToString(time.Now().Unix()),
		Role:      response.Role,
		SigninKey: h.cfg.SigningKey,
		Timeout:   cast.ToInt(h.cfg.AccessTokenTimeout),
	}

	accessToken, refreshToken, err := h.jwtHandler.GenerateAuthJWT()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("Failed to generate token", err)
		return
	}

	ctx.JSON(http.StatusCreated, models.RegisterResponse{
		ID:       response.ID,
		Username: response.Username,
		Password: response.Password,
		Role:     response.Role,
		Email:    response.ID,
		Token:    accessToken,
		Refresh:  refreshToken,
	}) // Changed to 201 for successful creation
}
