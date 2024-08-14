package handler

import (
	pb "Api_Gateway/genproto/health_analytics"
	"Api_Gateway/kafka/producer"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddMedicalRecord godoc
// @Summary Adds a new medical record
// @Description Creates a new medical record for a user
// @Tags medical
// @Security ApiKeyAuth
// @Param medical_record body health_analytics.AddMedicalRecordRequest true "Medical record data"
// @Success 201 {object} string "message"
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error while processing request"
// @Router /customer/medical [post]
func (h *Handler) AddMedicalRecord(c *gin.Context) {
	h.Log.Info("AddMedicalRecord handler is invoked")

	var request pb.AddMedicalRecordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.Log.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	writerKafka, err := producer.NewKafkaProducerInit([]string{"kafka:9092"})
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, err.Error())
		return

	}
	defer writerKafka.Close()
	msgBytes, err := json.Marshal(&request)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, err.Error())
		return

	}

	err = writerKafka.Producermessage("create", msgBytes)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, err.Error())
		return
	}


	h.Log.Info("Create Medical Record finished successfully")
	c.JSON(200, gin.H{"message": "Medical Record created successfully"})


}

// GetMedicalRecord godoc
// @Summary Gets a medical record
// @Description Retrieves a medical record by ID
// @Tags medical
// @Security ApiKeyAuth
// @Param id path string true "Medical record ID"
// @Success 200 {object} health_analytics.GetMedicalRecordResponse
// @Failure 400 {object} string "Invalid record ID"
// @Failure 500 {object} string "Server error while processing request"
// @Router /customer/medical/{id} [get]
func (h *Handler) GetMedicalRecord(c *gin.Context) {
	h.Log.Info("GetMedicalRecord handler is invoked")

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing record ID"})
		return
	}

	request := &pb.GetMedicalRecordRequest{RecordId: id}
	resp, err := h.Health.GetMedicalRecord(c.Request.Context(), request)
	if err != nil {
		h.Log.Error("Failed to get medical record", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get medical record"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateMedicalRecord godoc
// @Summary Updates a medical record
// @Description Updates an existing medical record
// @Tags medical
// @Security ApiKeyAuth
// @Param id path string true "Medical record ID"
// @Param medical_record body health_analytics.UpdateMedicalRecordRequest true "Updated medical record data"
// @Success 200 {object} health_analytics.UpdateMedicalRecordResponse
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Server error while processing request"
// @Router /customer/medical/{id} [put]
func (h *Handler) UpdateMedicalRecord(c *gin.Context) {
	h.Log.Info("UpdateMedicalRecord handler is invoked")

	var request pb.UpdateMedicalRecordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.Log.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	resp, err := h.Health.UpdateMedicalRecord(c.Request.Context(), &request)
	if err != nil {
		h.Log.Error("Failed to update medical record", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update medical record"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteMedicalRecord godoc
// @Summary Deletes a medical record
// @Description Deletes a medical record by ID
// @Tags medical
// @Security ApiKeyAuth
// @Param id path string true "Medical record ID"
// @Success 200 {object} health_analytics.DeleteMedicalRecordResponse
// @Failure 400 {object} string "Invalid record ID"
// @Failure 500 {object} string "Server error while processing request"
// @Router /customer/medical/{id} [delete]
func (h *Handler) DeleteMedicalRecord(c *gin.Context) {
	h.Log.Info("DeleteMedicalRecord handler is invoked")

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing record ID"})
		return
	}

	request := &pb.DeleteMedicalRecordRequest{RecordId: id}
	resp, err := h.Health.DeleteMedicalRecord(c.Request.Context(), request)
	if err != nil {
		h.Log.Error("Failed to delete medical record", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete medical record"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ListMedicalRecords godoc
// @Summary Lists medical records
// @Description Retrieves a list of medical records for a user
// @Tags medical
// @Security ApiKeyAuth
// @Param page query int false "Page number"
// @Param limit query int false "Number of records per page"
// @Success 200 {object} health_analytics.ListMedicalRecordsResponse
// @Failure 400 {object} string "Invalid query parameters"
// @Failure 500 {object} string "Server error while processing request"
// @Router /customer/medical [get]
func (h *Handler) ListMedicalRecords(c *gin.Context) {
	h.Log.Info("ListMedicalRecords handler is invoked")

	var request pb.ListMedicalRecordsRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		h.Log.Error("Failed to bind query parameters", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	resp, err := h.Health.ListMedicalRecords(c.Request.Context(), &request)
	if err != nil {
		h.Log.Error("Failed to list medical records", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list medical records"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
