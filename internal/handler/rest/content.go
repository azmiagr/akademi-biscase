package rest

import (
	"akademi-business-case/model"
	"akademi-business-case/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Rest) CreateContent(c *gin.Context) {
	classIDStr := c.Param("classID")
	classID, err := uuid.Parse(classIDStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid class ID", err)
		return
	}

	param := model.CreateContentRequest{}
	err = c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	result, err := r.service.ContentService.CreateContent(&param, classID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to create content", err)
		return
	}

	response.Success(c, http.StatusCreated, "success to create content", result)
}
