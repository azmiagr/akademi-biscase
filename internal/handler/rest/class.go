package rest

import (
	"akademi-business-case/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Rest) GetAllClasses(c *gin.Context) {
	classes, err := r.service.ClassService.GetAllClasses()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get all classes", err)
		return
	}

	response.Success(c, http.StatusOK, "success to get all classes", classes)
}

func (r *Rest) GetClassDetail(c *gin.Context) {
	classIDStr := c.Param("classID")
	classID, err := uuid.Parse(classIDStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid class ID", err)
		return
	}

	class, err := r.service.ClassService.GetClassDetail(classID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get class detail", err)
		return
	}

	response.Success(c, http.StatusOK, "success to get class detail", class)
}

func (r *Rest) GetClassesByType(c *gin.Context) {
	classTypeIDStr := c.Param("classTypeID")
	classTypeID, err := uuid.Parse(classTypeIDStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid class type ID", err)
		return
	}

	classes, err := r.service.ClassService.GetClassesByType(classTypeID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get classes by type", err)
		return
	}

	response.Success(c, http.StatusOK, "success to get classes by type", classes)
}

func (r *Rest) GetClassByName(c *gin.Context) {
	name := c.Query("name")

	classTypeIDStr := c.Query("classTypeID")
	var classTypeID uuid.UUID
	var err error

	if classTypeIDStr != "" {
		classTypeID, err = uuid.Parse(classTypeIDStr)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "invalid class type ID", err)
			return
		}
	}

	classes, err := r.service.ClassService.GetClassByName(name, classTypeID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get classes by name", err)
		return
	}

	response.Success(c, http.StatusOK, "success to get classes by name", classes)
}
