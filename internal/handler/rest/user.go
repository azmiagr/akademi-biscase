package rest

import (
	"akademi-business-case/entity"
	"akademi-business-case/model"
	"akademi-business-case/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) Register(c *gin.Context) {
	param := model.UserRegister{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	token, err := r.service.UserService.Register(&param)
	if err != nil {
		if err.Error() == "email already registered" {
			response.Error(c, http.StatusBadRequest, "failed to register new user", err)
			return
		} else if err.Error() == "password doesn't match" {
			response.Error(c, http.StatusBadRequest, "failed to register new user", err)
			return
		} else {
			response.Error(c, http.StatusInternalServerError, "failed to register new user", err)
			return
		}

	}

	response.Success(c, http.StatusCreated, "success to register new user", token)

}

func (r *Rest) VerifyUser(c *gin.Context) {
	var param model.VerifyUser
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	err = r.service.UserService.VerifyUser(param)
	if err != nil {
		if err.Error() == "invalid otp code" {
			response.Error(c, http.StatusUnauthorized, "otp code is wrong", err)
			return
		} else if err.Error() == "otp expired" {
			response.Error(c, http.StatusUnauthorized, "otp code is expired", err)
			return
		} else {
			response.Error(c, http.StatusInternalServerError, "failed to verify user", err)
			return
		}
	}

	response.Success(c, http.StatusOK, "success to verify user", nil)

}

func (r *Rest) Login(c *gin.Context) {
	param := model.UserLogin{}

	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	result, err := r.service.UserService.Login(param)
	if err != nil {
		if err.Error() == "email or password is wrong" {
			response.Error(c, http.StatusUnauthorized, "email or password is wrong", err)
			return
		} else {
			response.Error(c, http.StatusInternalServerError, "failed to login user", err)
			return
		}
	}

	response.Success(c, http.StatusOK, "success to login user", result)
}

func (r *Rest) GetUserProfile(c *gin.Context) {
	user := c.MustGet("user").(*entity.User)

	resp, err := r.service.UserService.GetUserProfile(user.UserID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get user profile", err)
		return
	}

	response.Success(c, http.StatusOK, "success to get user profile", resp)
}

func (r *Rest) GetMentors(c *gin.Context) {
	mentors, err := r.service.UserService.GetMentors()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to get members", err)
		return
	}

	response.Success(c, http.StatusOK, "success to get members", mentors)
}
