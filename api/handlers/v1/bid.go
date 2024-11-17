package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"log"
	"net/http"
	"tender/api/models"
	"tender/storage/repo"
	"time"
)

func (h *handlerV1) SubmitBid(c *gin.Context) {
	timeout, err := time.ParseDuration(h.cfg.CtxTimeout)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println(http.StatusInternalServerError, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
	defer cancel()

	var request models.Bid

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Message: "Invalid bid data",
		})
		log.Println(http.StatusBadRequest, err.Error())
		return
	}

	tenderID := c.Query("tender_id")

	response, err := h.storage.Bid().SubmitBid(ctx, &repo.SubmitBidRequest{
		TenderID:     cast.ToInt(tenderID),
		ContractorID: "",
		Price:        request.Price,
		DeliveryTime: request.DeliveryTime,
		Comments:     request.Comments,
	})

	c.JSON(http.StatusCreated, response)
}

func (h *handlerV1) ViewSubmittedBids(c *gin.Context) {}

func (h *handlerV1) AwardBid(c *gin.Context) {}

func (h *handlerV1) DeleteBid(c *gin.Context) {}
