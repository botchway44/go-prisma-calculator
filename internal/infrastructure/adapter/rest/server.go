package rest

import (
	"log/slog"
	"net/http"

	"go-prisma-calculator/internal/domain/ports/in"

	"github.com/gin-gonic/gin"
)

// Adapter is the REST API adapter.
type Adapter struct {
	usecase in.CalculatorPort
	logger  *slog.Logger
}

// NewAdapter is the constructor that fx uses to create an instance.
// It receives the application port and logger as dependencies.
func NewAdapter(usecase in.CalculatorPort, logger *slog.Logger) *Adapter {
	return &Adapter{usecase: usecase, logger: logger}
}

// calcRequest defines the structure for incoming JSON requests.
type calcRequest struct {
	A int32 `json:"a"`
	B int32 `json:"b"`
}

// AddHandler handles HTTP POST requests to the /add endpoint.
// @Summary      Add two numbers
// @Description  Takes two integers and returns their sum.
// @Accept       json
// @Produce      json
// @Param        request body rest.calcRequest true "Add Request"
// @Success      200  {object} map[string]int
// @Router       /add [post]
func (a *Adapter) AddHandler(c *gin.Context) {
	var req calcRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		a.logger.Error("Failed to bind JSON request", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	a.logger.Info("Handling REST Add request", slog.Int("a", int(req.A)), slog.Int("b", int(req.B)))

	calculation, err := a.usecase.Add(c.Request.Context(), req.A, req.B)
	if err != nil {
		a.logger.Error("Usecase failed for REST Add", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to perform addition"})
		return
	}

	a.logger.Info("REST Add request successful", slog.Int("result", calculation.Result))
	c.JSON(http.StatusOK, gin.H{"result": calculation.Result})
}

// DivideHandler handles HTTP POST requests to the /divide endpoint.
// @Summary      Divide two numbers
// @Description  Takes two integers and returns their quotient.
// @Accept       json
// @Produce      json
// @Param        request body rest.calcRequest true "Divide Request"
// @Success      200  {object} map[string]int
// @Router       /divide [post]
func (a *Adapter) DivideHandler(c *gin.Context) {
	var req calcRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		a.logger.Error("Failed to bind JSON request", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	a.logger.Info("Handling REST Divide request", slog.Int("a", int(req.A)), slog.Int("b", int(req.B)))
	
	calculation, err := a.usecase.Divide(c.Request.Context(), req.A, req.B)
	if err != nil {
		a.logger.Error("Usecase failed for REST Divide", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a.logger.Info("REST Divide request successful", slog.Int("result", calculation.Result))
	c.JSON(http.StatusOK, gin.H{"result": calculation.Result})
}