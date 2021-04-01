package http

import (
	"FBSTestTask/internal/models"
	"FBSTestTask/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	fibonacciService service.FibonacciService
}

func NewHandler(fibSvc service.FibonacciService) *Handler {
	return &Handler{
		fibonacciService: fibSvc,
	}
}

func (h *Handler) Init() *gin.Engine {
	r := gin.Default()

	r.POST("/api/v1/getFibSlice", h.GetFibSlice)

	return r
}

func (h *Handler) GetFibSlice(c *gin.Context) {
	var req models.FibonacciRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fibSliced, err := h.fibonacciService.GetFibSlice(c, req.Start, req.End)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	r := models.FibonacciResponse{Result: fibSliced}

	c.JSON(http.StatusOK, r)

}
