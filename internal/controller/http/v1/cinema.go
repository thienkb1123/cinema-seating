package v1

import (
	"cinema-seating/internal/entity"
	settings "cinema-seating/internal/usecase/cinema"
	"cinema-seating/pkg/errors"
	"cinema-seating/pkg/logger"
	"cinema-seating/pkg/utils"
	"github.com/gin-gonic/gin"
)

type seatingRoutes struct {
	uc     settings.Cinema
	logger *logger.Logger
}

func NewSeatingRoutes(handler *gin.RouterGroup, logger *logger.Logger, uc settings.Cinema) {
	r := seatingRoutes{uc, logger}
	h := handler.Group("/cinema")
	{
		h.POST("/configure", r.configureCinema)
		h.GET("/available-seats", r.getAvailableSeats)
		h.POST("/reserve", r.reserveSeat)
		h.POST("/cancel", r.cancelSeat)
	}
}

// @BasePath /v1
// @Summary Configure cinema
// @Description Configure cinema
// @Tags Seating
// @Accept json
// @Produce json
// @Param body body entity.Cinema true "Cinema"
// @Success 201 {object} entity.Cinema
// @Router /cinema/configure [post]
func (h *seatingRoutes) configureCinema(c *gin.Context) {
	var req entity.Cinema
	if err := utils.ReadBodyRequest(c, &req); err != nil {
		h.logger.Error(errors.WithMessage(err, "http - v1 - configure cinema read body request"))
		errorResponse(c, err)
		return
	}

	ctx := c.Request.Context()
	err := h.uc.ConfigureCinema(ctx, &req)
	if err != nil {
		h.logger.Error(errors.WithMessage(err, "http - v1 - configure cinema"))
		errorResponse(c, err)
		return
	}

	successResponseWithCreated(c)
}

// @BasePath /v1
// @Summary Get available seats
// @Description Get available seats
// @Tags Seating
// @Accept json
// @Produce json
// @Success 200 {object} entity.Seat
// @Router /cinema/available-seats [get]
func (h *seatingRoutes) getAvailableSeats(c *gin.Context) {
	ctx := c.Request.Context()
	seats, err := h.uc.GetAvailableSeats(ctx)
	if err != nil {
		h.logger.Error(errors.WithMessage(err, "http - v1 - get available seats"))
		errorResponse(c, err)
		return
	}

	successResponseWithOK(c, seats)
}

// @BasePath /v1
// @Summary Reserve seat
// @Description Reserve seat
// @Tags Seating
// @Accept json
// @Produce json
// @Param body body entity.SeatAction true "SeatAction"
// @Success 200 {object} entity.SeatAction
// @Router /cinema/reserve [post]
func (h *seatingRoutes) reserveSeat(c *gin.Context) {
	var req entity.SeatAction
	if err := utils.ReadBodyRequest(c, &req); err != nil {
		h.logger.Error(errors.WithMessage(err, "http - v1 - reserve seat read body request"))
		errorResponse(c, err)
		return
	}

	ctx := c.Request.Context()
	err := h.uc.ReserveSeat(ctx, &req)
	if err != nil {
		h.logger.Error(errors.WithMessage(err, "http - v1 - reserve seat"))
		errorResponse(c, err)
		return
	}

	successResponseWithOK(c, nil)
}

// @BasePath /v1
// @Summary Cancel seat
// @Description Cancel seat
// @Tags Seating
// @Accept json
// @Produce json
// @Param body body entity.SeatAction true "SeatAction"
// @Success 200 {object} entity.SeatAction
// @Router /cinema/cancel [post]
func (h *seatingRoutes) cancelSeat(c *gin.Context) {
	var req entity.SeatAction
	if err := utils.ReadBodyRequest(c, &req); err != nil {
		h.logger.Error(errors.WithMessage(err, "http - v1 - cancel seat read body request"))
		errorResponse(c, err)
		return
	}

	ctx := c.Request.Context()
	err := h.uc.CancelSeat(ctx, &req)
	if err != nil {
		h.logger.Error(errors.WithMessage(err, "http - v1 - cancel seat"))
		errorResponse(c, err)
		return
	}

	successResponseWithOK(c, nil)
}
