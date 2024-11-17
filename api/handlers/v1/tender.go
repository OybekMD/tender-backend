package v1

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"tender/api/models"
	"tender/storage/repo"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

// @Security      BearerAuth
// @Summary       Create Tender
// @Description   This API is for creating a new tender
// @Tags          tenders
// @Accept        json
// @Produce       json
// @Param         TenderCreate body models.TenderCreate true "TenderCreate Model"
// @Success       201 {object} models.TenderResponse
// @Failure       400 {object} models.Error
// @Failure       401 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /api/client/tenders [POST]
func (h *handlerV1) CreateTender(ctx *gin.Context) {
	var body models.TenderCreate

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &models.Error{
			Message: models.WrongInfoMessage,
		})
		log.Println("failed to bind json", err.Error())
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

	pp.Println(body)

	response, err := h.storage.Tender().Create(ctxTime, &repo.Tender{
		ClientID:    body.ClientID,
		Title:       body.Title,
		Description: body.Description,
		Deadline:    body.Deadline,
		Budget:      body.Budget,
		Status:      body.Status,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotCreatedMessage,
		})
		log.Println("failed to create tender", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, &models.TenderResponse{
		ID:          response.ID,
		ClientID:    response.ClientID,
		Title:       response.Title,
		Description: response.Description,
		Deadline:    response.Deadline,
		Budget:      response.Budget,
		Status:      response.Status,
	})
}

// @Security      BearerAuth
// @Summary       List Tenders
// @Description   This API is for getting all tenders
// @Tags          tenders
// @Accept        json
// @Produce       json
// @Success       200 {object} []models.TenderResponse
// @Failure       400 {object} models.Error
// @Failure       401 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /api/client/tenders [GET]
func (h *handlerV1) ListTenders(ctx *gin.Context) {
	duration, err := time.ParseDuration(h.cfg.CtxTimeout)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("failed to parse timeout", err.Error())
		return
	}

	pp.Println("Heloooooooooooooooooooooooo")
	// for key, values := range ctx.Request.Header {
	// 	for _, value := range values {
	// 		log.Printf("%s: %s", key, value)
	// 	}
	// }

	ctxTime, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	tenders, err := h.storage.Tender().List(ctxTime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get all tenders", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, tenders)
}

// @Security      BearerAuth
// @Summary       Update Tender Status
// @Description   This API is for updating tender status
// @Tags          tenders
// @Accept        json
// @Produce       json
// @Param         id path string true "ID"
// @Param         TenderUpdate body models.TenderUpdate true "Update Tender Model"
// @Success       200 {object} models.TenderResponse
// @Failure       400 {object} models.Error
// @Failure       401 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /api/client/tenders/{id} [PUT]
func (h *handlerV1) UpdateTenderStatus(ctx *gin.Context) {
	a := ctx.Param("id")
	fmt.Println(a)
	var body models.TenderUpdate

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &models.Error{
			Message: models.WrongInfoMessage,
		})
		log.Println("failed to bind json", err.Error())
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

	tenderModel := &repo.Tender{
		ID:     body.ID,
		Status: body.Status,
	}

	response, err := h.storage.Tender().Update(ctxTime, tenderModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotUpdatedMessage,
		})
		log.Println("failed to update tender", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// @Security      BearerAuth
// @Summary       Delete Tender
// @Description   This API is for deleting a tender
// @Tags          tenders
// @Accept        json
// @Produce       json
// @Param         id path string true "ID"
// @Success       200 {object} bool
// @Failure       401 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /api/client/tenders/{id} [DELETE]
func (h *handlerV1) DeleteTender(ctx *gin.Context) {
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

	err = h.storage.Tender().Delete(ctxTime, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotDeletedMessage,
		})
		log.Println("failed to delete tender", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, true)
}