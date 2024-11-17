package v1

/*
import (
	"tender/api/helper/hashing"
	"tender/api/helper/parsing"
	"tender/api/helper/utils"
	"context"
	"log"
	"net/http"
	"time"

	"tender/api/models"
	"tender/storage/repo"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Security      BearerAuth
// @Summary 	  Create Student
// @Description   This Api for creating a new student
// @Tags 		  students
// @Accept 		  json
// @Produce 	  json
// @Param 		  StudentInfo body models.StudentInfo true "StudentInfo Model"
// @Success 	  201 {object} models.StudentInfo
// @Failure 	  400 {object} models.Error
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/student [POST]
func (h *handlerV1) CreateStudent(ctx *gin.Context) {
	var body models.StudentInfo

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &models.Error{
			Message: models.WrongInfoMessage,
		})
		log.Println("failed to bind json", err.Error())
		return
	}

	err = body.Student.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: models.WrongInfoMessage,
		})
		log.Println("error validating students data", err.Error())
		return
	}

	err = body.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: models.WrongInfoMessage,
		})
		log.Println("error validating students data", err.Error())
		return
	}

	duration, err := time.ParseDuration(h.cfg.CtxTimeout)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("failed to parse timeout", err.Error())
		return
	}

	ctxTime, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	body.Student.Id = uuid.New().String()

	response, err := h.storage.Student().Create(
		ctxTime,
		&repo.Student{
			Id:          body.Student.Id,
			FirstName:   body.Student.FirstName,
			LastName:    body.Student.LastName,
			Age:         body.Student.Age,
			ClassNumber: body.Student.ClassNumber,
			PhoneNumber: body.Student.PhoneNumber,
			Email:       body.Student.Email,
			Gender:      body.Student.Gender,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotCreatedMessage,
		})
		log.Println("failed to create user", err.Error())
		return
	}

	hashPassword, err := hashing.HashPassword(body.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("failed to hashing password (student)", err.Error())
		return
	}

	responsePassword, err := h.storage.Login().SavePassword(
		ctxTime,
		&repo.LoginPassword{
			UserId:   response.Id,
			Role:     "student",
			Password: hashPassword,
		},
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotCreatedMessage,
		})
		log.Println("failed to generate login to student", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, &models.StudentInfo{
		Student: models.Student{
			Id:          response.Id,
			FirstName:   response.FirstName,
			LastName:    response.LastName,
			Age:         response.Age,
			ClassNumber: response.ClassNumber,
			PhoneNumber: response.PhoneNumber,
			Gender:      response.Gender,
			Email:       response.Email,
		},
		Login:    responsePassword.Login,
		Password: responsePassword.Password,
	})
}

// @Security      BearerAuth
// @Summary 	  Update Student
// @Description   This Api for updating student
// @Tags 		  students
// @Accept 		  json
// @Produce 	  json
// @Param 		  Student body models.Student true "Update Student Model"
// @Success 	  200 {object} models.Student
// @Failure 	  400 {object} models.Error
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/student [PUT]
func (h *handlerV1) UpdateStudent(ctx *gin.Context) {
	var body models.Student

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &models.Error{
			Message: models.WrongInfoMessage,
		})
		log.Println("failed to bind json", err.Error())
		return
	}

	err = body.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: models.WrongInfoMessage,
		})
		log.Println("Error validating student", err.Error())
		return
	}

	duration, err := time.ParseDuration(h.cfg.CtxTimeout)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("failed to parse timeout", err.Error())
		return
	}

	ctxTime, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	studentModel := &repo.Student{}
	err = parsing.StructToStruct(&body, studentModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("Error parsing struct to struct", err.Error())
		return
	}

	response, err := h.storage.Student().Update(ctxTime, studentModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotUpdatedMessage,
		})
		log.Println("failed to update user", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// @Security      BearerAuth
// @Summary 	  Delete Student
// @Description   This Api for deleting student
// @Tags 		  students
// @Accept 		  json
// @Produce 	  json
// @Param         id path string true "ID"
// @Success 	  200 {object} bool
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/student/{id} [DELETE]
func (h *handlerV1) DeleteStudent(ctx *gin.Context) {
	id := ctx.Param("id")

	duration, err := time.ParseDuration(h.cfg.CtxTimeout)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("failed to parse timeout", err.Error())
		return
	}

	ctxTime, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	response, err := h.storage.Student().Delete(ctxTime, id)
	if err != nil || !response {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotDeletedMessage,
		})
		log.Println("failed to delete user", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, true)
}

// @Security      BearerAuth
// @Summary 	  Get Student
// @Description   This Api for get student
// @Tags 		  students
// @Accept        json
// @Produce       json
// @Param         id path string true "ID"
// @Success 	  200 {object} models.Student
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/student/{id} [GET]
func (h *handlerV1) GetStudent(ctx *gin.Context) {
	id := ctx.Param("id")

	duration, err := time.ParseDuration(h.cfg.CtxTimeout)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("failed to parse timeout", err.Error())
		return
	}

	ctxTime, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	student, err := h.storage.Student().Get(ctxTime, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get student", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, student)
}

// @Security      BearerAuth
// @Summary       ListStudents
// @Description   This Api for get all students
// @Tags          students
// @Accept        json
// @Produce       json
// @Param         page query uint64 true "Page"
// @Param         limit query uint64 true "Limit"
// @Success 	  200 {object} []models.Student
// @Failure		  400 {object} models.Error
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/students [GET]
func (h *handlerV1) ListStudents(ctx *gin.Context) {
	queryParams := ctx.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to parse query params json" + errStr[0])
		return
	}

	duration, err := time.ParseDuration(h.cfg.CtxTimeout)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("failed to parse timeout", err.Error())
		return
	}

	ctxTime, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	students, err := h.storage.Student().GetAll(ctxTime, uint64(params.Page), uint64(params.Limit))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get all users", err.Error())
		return
	}
	if len(students) == 0 {
		ctx.JSON(http.StatusOK, nil)
		log.Println("Not found students")
		return
	}

	ctx.JSON(http.StatusOK, students)
}
*/
