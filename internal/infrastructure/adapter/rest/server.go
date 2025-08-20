package rest

import (
	"net/http"

	"go-prisma-calculator/internal/domain/ports/in"

	"github.com/gin-gonic/gin"
)

// Adapter is the REST API adapter.
type Adapter struct {
	usecase in.CalculatorPort
}

// NewAdapter creates a new instance of the REST adapter.
func NewAdapter(usecase in.CalculatorPort) *Adapter {
	return &Adapter{usecase: usecase}
}

// calcRequest defines the structure for JSON requests.
type calcRequest struct {
	A int32 `json:"a"`
	B int32 `json:"b"`
}

// AddHandler handles HTTP POST requests to the /add endpoint.
func (a *Adapter) AddHandler(c *gin.Context) {
	var req calcRequest
	// Parse the JSON body from the request.
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Call the exact same application core logic.
	calculation, err := a.usecase.Add(c.Request.Context(), req.A, req.B)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to perform addition"})
		return
	}

	// Send a successful JSON response.
	c.JSON(http.StatusOK, gin.H{"result": calculation.Result})
}

// DivideHandler handles HTTP POST requests to the /divide endpoint.
func (a *Adapter) DivideHandler(c *gin.Context) {
	var req calcRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	calculation, err := a.usecase.Divide(c.Request.Context(), req.A, req.B)
	if err != nil {
		// Here we can handle specific business logic errors.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": calculation.Result})
}
