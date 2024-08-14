package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	pb "Api_Gateway/genproto/health_analytics"
)

// AddWearableData godoc
// @Summary Adds new wearable data
// @Description Creates a new wearable data record for a user
// @Tags wearable
// @Security ApiKeyAuth
// @Param wearable_data body health_analytics.AddWearableDataRequest true "Wearable data"
// @Success 201 {object} health_analytics.AddWearableDataResponse
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error while processing request"
// @Router /customer/wearable [post]
func (h *Handler) AddWearableData(c *gin.Context) {
	h.Log.Info("AddWearableData handler is invoked")

	var request pb.AddWearableDataRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.Log.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := h.Health.AddWearableData(c.Request.Context(), &request)
	if err != nil {
		h.Log.Error("Failed to add wearable data", "error", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetWearableData godoc
// @Summary Gets wearable data
// @Description Retrieves wearable data by ID
// @Tags wearable
// @Security ApiKeyAuth
// @Param id path string true "Wearable data ID"
// @Success 200 {object} health_analytics.GetWearableDataResponse
// @Failure 400 {object} string "Invalid data ID"
// @Failure 500 {object} string "Server error while processing request"
// @Router /customer/wearable/{id} [get]
func (h *Handler) GetWearableData(c *gin.Context) {
	h.Log.Info("GetWearableData handler is invoked")

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing wearable data ID"})
		return
	}

	request := &pb.GetWearableDataRequest{DataId: id}
	resp, err := h.Health.GetWearableData(c.Request.Context(), request)
	if err != nil {
		h.Log.Error("Failed to get wearable data", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get wearable data"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateWearableData godoc
// @Summary Updates wearable data
// @Description Updates existing wearable data
// @Tags wearable
// @Security ApiKeyAuth
// @Param id path string true "Wearable data ID"
// @Param wearable_data body health_analytics.UpdateWearableDataRequest true "Updated wearable data"
// @Success 200 {object} health_analytics.UpdateWearableDataResponse
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error while processing request"
// @Router /customer/wearable/{id} [put]
func (h *Handler) UpdateWearableData(c *gin.Context) {
	h.Log.Info("UpdateWearableData handler is invoked")

	var request pb.UpdateWearableDataRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.Log.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := h.Health.UpdateWearableData(c.Request.Context(), &request)
	if err != nil {
		h.Log.Error("Failed to update wearable data", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update wearable data"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteWearableData godoc
// @Summary Deletes wearable data
// @Description Deletes wearable data by ID
// @Tags wearable
// @Security ApiKeyAuth
// @Param id path string true "Wearable data ID"
// @Success 200 {object} health_analytics.DeleteWearableDataResponse
// @Failure 400 {object} string "Invalid data ID"
// @Failure 500 {object} string "Server error while processing request"
// @Router /customer/wearable/{id} [delete]
func (h *Handler) DeleteWearableData(c *gin.Context) {
	h.Log.Info("DeleteWearableData handler is invoked")

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing wearable data ID"})
		return
	}

	request := &pb.DeleteWearableDataRequest{DataId: id}
	resp, err := h.Health.DeleteWearableData(c.Request.Context(), request)
	if err != nil {
		h.Log.Error("Failed to delete wearable data", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete wearable data"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
