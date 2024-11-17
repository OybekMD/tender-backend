package v1

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"tender/api/helper/utils"
	"tender/api/models"
	"tender/storage/repo"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
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
			Message: "Invalid tender data",
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

	claims, err := utils.GetClaimsFromToken(ctx.Request, h.cfg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: "Invalid tender data",
		})
		log.Println("failed to get GetClaimsFromToken", err.Error())
		return
	}

	if body.Deadline != "" {
		if err := body.ValidateTimeAndPrice(); err != nil {
			ctx.JSON(http.StatusBadRequest, models.Error{
				Message: "Invalid tender data",
			})
			return
		}
	}

	response, err := h.storage.Tender().Create(ctxTime, &repo.Tender{
		ClientID:    cast.ToString(claims["sub"]),
		Title:       body.Title,
		Description: body.Description,
		Deadline:    body.Deadline,
		Budget:      body.Budget,
		Status:      body.Status,
	})
	if err != nil {
		fmt.Println("eerrrooroo:", err)
		ctx.JSON(http.StatusBadRequest, &models.Error{
			Message: "Invalid input",
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
	tender_id := ctx.Param("id")
	var body models.TenderUpdate

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &models.Error{
			Message: "Invalid tender data",
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

	if err := body.ValidateTenderStatus(); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid tender status",
		})
		return
	}

	tenderModel := &repo.Tender{
		ID:     cast.ToUint(tender_id),
		Status: body.Status,
	}

	_, err = h.storage.Tender().Update(ctxTime, tenderModel)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, &models.Error{
			Message: "Tender not found",
		})
		log.Printf("Failed to update tender: %v\n", err)
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: "Tender not found",
		})
		log.Println("failed to update tender", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, models.AlertMessage{Message: "Tender status updated"})
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
		fmt.Println("ERRorror oor ",err)
		if err.Error() == "not found" {
			ctx.JSON(http.StatusNotFound, &models.Error{
				Message: "Tender not found or access denied",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotDeletedMessage,
		})
		log.Println("failed to delete tender", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, models.AlertMessage{
		Message: "Tender deleted successfully",
	})
}
