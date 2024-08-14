package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	pb "Api_Gateway/genproto/health_analytics"
)

// AddLifestyleData godoc
// @Summary Adds new lifestyle data
// @Description Creates a new lifestyle data record for a user
// @Tags lifestyle
// @Security ApiKeyAuth
// @Param lifestyle_data body health_analytics.AddLifestyleDataRequest true "Lifestyle data"
// @Success 201 {object} health_analytics.AddLifestyleDataResponse
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error while processing request"
// @Router /customer/lifestyle [post]
func (h *Handler) AddLifestyleData(c *gin.Context) {
	h.Log.Info("AddLifestyleData handler is invoked")

	var request pb.AddLifestyleDataRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.Log.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := h.Health.AddLifestyleData(c.Request.Context(), &request)
	if err != nil {
		h.Log.Error("Failed to add lifestyle data", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add lifestyle data"})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetLifestyleData godoc
// @Summary Gets lifestyle data
// @Description Retrieves lifestyle data by ID
// @Tags lifestyle
// @Security ApiKeyAuth
// @Param id path string true "Lifestyle data ID"
// @Success 200 {object} health_analytics.GetLifestyleDataResponse
// @Failure 400 {object} string "Invalid data ID"
// @Failure 500 {object} string "Server error while processing request"
// @Router /customer/lifestyle/{id} [get]
func (h *Handler) GetLifestyleData(c *gin.Context) {
	h.Log.Info("GetLifestyleData handler is invoked")

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing lifestyle data ID"})
		return
	}

	request := &pb.GetLifestyleDataRequest{DataId: id}
	resp, err := h.Health.GetLifestyleData(c.Request.Context(), request)
	if err != nil {
		h.Log.Error("Failed to get lifestyle data", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get lifestyle data"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateLifestyleData godoc
// @Summary Updates lifestyle data
// @Description Updates existing lifestyle data
// @Tags lifestyle
// @Security ApiKeyAuth
// @Param id path string true "Lifestyle data ID"
// @Param lifestyle_data body health_analytics.UpdateLifestyleDataRequest true "Updated lifestyle data"
// @Success 200 {object} health_analytics.UpdateLifestyleDataResponse
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error while processing request"
// @Router /customer/lifestyle/{id} [put]
func (h *Handler) UpdateLifestyleData(c *gin.Context) {
	h.Log.Info("UpdateLifestyleData handler is invoked")

	var request pb.UpdateLifestyleDataRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.Log.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := h.Health.UpdateLifestyleData(c.Request.Context(), &request)
	if err != nil {
		h.Log.Error("Failed to update lifestyle data", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update lifestyle data"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteLifestyleData godoc
// @Summary Deletes lifestyle data
// @Description Deletes lifestyle data by ID
// @Tags lifestyle
// @Security ApiKeyAuth
// @Param id path string true "Lifestyle data ID"
// @Success 200 {object} health_analytics.DeleteLifestyleDataResponse
// @Failure 400 {object} string "Invalid data ID"
// @Failure 500 {object} string "Server error while processing request"
// @Router /customer/lifestyle/{id} [delete]
func (h *Handler) DeleteLifestyleData(c *gin.Context) {
	h.Log.Info("DeleteLifestyleData handler is invoked")

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing lifestyle data ID"})
		return
	}

	request := &pb.DeleteLifestyleDataRequest{DataId: id}
	resp, err := h.Health.DeleteLifestyleData(c.Request.Context(), request)
	if err != nil {
		h.Log.Error("Failed to delete lifestyle data", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete lifestyle data"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
