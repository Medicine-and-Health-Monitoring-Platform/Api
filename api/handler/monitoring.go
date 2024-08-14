package handler

import (
	"net/http"
	pb "Api_Gateway/genproto/health_analytics"

	"github.com/gin-gonic/gin"
)

// GenerateHealthRecommendations godoc
// @Summary Generates health recommendations
// @Description Generates personalized health recommendations based on user's health data
// @Tags monitoring
// @Security ApiKeyAuth
// @Success 200 {object} health_analytics.GenerateHealthRecommendationsResponse
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error while processing request"
// @Router /customer/monitoring/recommendations [post]
func (h *Handler) GenerateHealthRecommendations(c *gin.Context) {
	h.Log.Info("GenerateHealthRecommendations handler is invoked")

	var request pb.GenerateHealthRecommendationsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.Log.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := h.Health.GenerateHealthRecommendations(c.Request.Context(), &request)
	if err != nil {
		h.Log.Error("Failed to generate health recommendations", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate health recommendations"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetRealtimeHealthMonitoring godoc
// @Summary Gets real-time health monitoring data
// @Description Retrieves real-time health monitoring data for the user
// @Tags monitoring
// @Security ApiKeyAuth
// @Success 200 {object} health_analytics.GetRealtimeHealthMonitoringResponse
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error while processing request"
// @Router /customer/monitoring/realtime [get]
func (h *Handler) GetRealtimeHealthMonitoring(c *gin.Context) {
	h.Log.Info("GetRealtimeHealthMonitoring handler is invoked")

	var request pb.GetRealtimeHealthMonitoringRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.Log.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := h.Health.GetRealtimeHealthMonitoring(c.Request.Context(), &request)
	if err != nil {
		h.Log.Error("Failed to get real-time health monitoring data", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get real-time health monitoring data"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetDailyHealthSummary godoc
// @Summary Gets daily health summary
// @Description Retrieves daily health summary for the user
// @Tags monitoring
// @Security ApiKeyAuth
// @Success 200 {object} health_analytics.GetDailyHealthSummaryResponse
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error while processing request"
// @Router /customer/monitoring/daily [get]
func (h *Handler) GetDailyHealthSummary(c *gin.Context) {
	h.Log.Info("GetDailyHealthSummary handler is invoked")

	var request pb.GetDailyHealthSummaryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.Log.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := h.Health.GetDailyHealthSummary(c.Request.Context(), &request)
	if err != nil {
		h.Log.Error("Failed to get daily health summary", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get daily health summary"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetWeeklyHealthSummary godoc
// @Summary Gets weekly health summary
// @Description Retrieves weekly health summary for the user
// @Tags monitoring
// @Security ApiKeyAuth
// @Success 200 {object} health_analytics.GetWeeklyHealthSummaryResponse
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error while processing request"
// @Router /customer/monitoring/weekly [get]
func (h *Handler) GetWeeklyHealthSummary(c *gin.Context) {
	h.Log.Info("GetWeeklyHealthSummary handler is invoked")

	var request pb.GetWeeklyHealthSummaryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.Log.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := h.Health.GetWeeklyHealthSummary(c.Request.Context(), &request)
	if err != nil {
		h.Log.Error("Failed to get weekly health summary", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get weekly health summary"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
