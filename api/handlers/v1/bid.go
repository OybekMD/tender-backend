package v1

import (
	"context"
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

// SubmitBid
// @Security      BearerAuth
// @Summary       Submit Bid
// @Description   This API is for submit a bid to tender by contractor
// @Tags          bids
// @Accept        json
// @Produce       json
// @Param         id path string true "ID"
// @Success       200 {object} bool
// @Failure       401 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /api/contractor/tenders/{id}/bid [POST]
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

	tenderID := c.Param("id")
	fmt.Println("IDDDDDDDDDDDDDDDDDDDDDDDDDDDDD:", tenderID)

	claims, err := utils.GetClaimsFromToken(c.Request, h.cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println(http.StatusInternalServerError, err.Error())
		return
	}
	response, err := h.storage.Bid().SubmitBid(ctx, &repo.SubmitBidRequest{
		TenderID:     cast.ToInt(tenderID),
		ContractorID: cast.ToString(claims["sub"]),
		Price:        request.Price,
		DeliveryTime: request.DeliveryTime,
		Comments:     request.Comments,
	})
	if err != nil {
		fmt.Println("Eroororor:", err)
		c.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h *handlerV1) ViewSubmittedBids(c *gin.Context) {}

func (h *handlerV1) AwardBid(c *gin.Context) {}

func (h *handlerV1) DeleteBid(c *gin.Context) {}
