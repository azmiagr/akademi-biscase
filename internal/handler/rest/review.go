package rest

import (
	"akademi-business-case/entity"
	"akademi-business-case/model"
	"akademi-business-case/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Rest) AddReview(c *gin.Context) {
	user := c.MustGet("user").(*entity.User)

	param := model.AddReviewRequest{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	classIDStr := c.Param("classID")
	classID, err := uuid.Parse(classIDStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid class ID", err)
		return
	}

	resp, err := r.service.ReviewService.AddReview(&param, user.UserID, classID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to add review", err)
		return
	}

	response.Success(c, http.StatusCreated, "success to add review", resp)
}
