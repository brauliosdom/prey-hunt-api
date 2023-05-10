package server

import (
	"functional/prey"
	"functional/shark"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	shark shark.Shark
	prey  prey.Prey
}

func NewHandler(shark shark.Shark, prey prey.Prey) *Handler {
	return &Handler{shark: shark, prey: prey}
}

// PUT: /v1/shark

func (h *Handler) ConfigureShark() gin.HandlerFunc {
	type request struct {
		XPosition float64 `json:"x_position"`
		YPosition float64 `json:"y_position"`
		Speed     float64 `json:"speed"`
	}
	type response struct {
		Success bool `json:"success"`
	}

	return func(context *gin.Context) {
		var request request
		var response response
		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, response)
			return
		}

		position := [2]float64{request.XPosition, request.YPosition}
		h.shark.Configure(position, request.Speed)
		response.Success = true
		context.JSON(http.StatusOK, response)
	}
}

// PUT: /v1/prey

func (h *Handler) ConfigurePrey() gin.HandlerFunc {
	type request struct {
		Speed float64 `json:"speed"`
	}
	type response struct {
		Success bool `json:"success"`
	}

	return func(context *gin.Context) {
		var request request
		var response response
		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, response)
			return
		}

		h.prey.SetSpeed(request.Speed)
		response.Success = true
		context.JSON(http.StatusOK, response)
	}
}

// POST: /v1/simulate

func (h *Handler) SimulateHunt() gin.HandlerFunc {
	type response struct {
		Success bool    `json:"success"`
		Message string  `json:"message"`
		Time    float64 `json:"time"`
	}

	return func(context *gin.Context) {
		var response response
		err, time := h.shark.Hunt(h.prey)
		if err != nil {
			context.JSON(http.StatusInternalServerError, response)
			return
		}

		response.Success = true
		response.Message = "ok"
		response.Time = time
		context.JSON(http.StatusOK, response)
	}
}
