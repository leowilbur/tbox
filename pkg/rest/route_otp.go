package rest

import (
	"math"
	"net/http"

	"github.com/leowilbur/tbox/pkg/model"
	"github.com/leowilbur/tbox/pkg/service"
	validator "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// OTPGenerate generate otp for input phone number
func (a *API) OTPGenerate(r *gin.Context) {
	var req model.OTP

	if err := r.BindJSON(&req); err != nil {
		r.JSON(http.StatusBadRequest, model.APIResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Can not parse request body input",
		})
		return
	}

	if _, err := validator.ValidateStruct(req); err != nil {
		r.JSON(http.StatusBadRequest, model.APIResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	if err := service.OTPGenerate(a.DB, req); err != nil {
		r.JSON(http.StatusBadRequest, model.APIResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, model.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "SUCCESS",
	})
}

// OTPValidate validate otp for input phone number
func (a *API) OTPValidate(r *gin.Context) {
	var req model.OTP

	if err := r.BindJSON(&req); err != nil {
		r.JSON(http.StatusBadRequest, model.APIResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Can not parse request body input",
		})
		return
	}

	if err := service.OTPValidate(a.DB, req); err != nil {
		r.JSON(http.StatusBadRequest, model.APIResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	token, err := service.TokenGenerate(map[string]interface{}{
		"phoneNumber": req.PhoneNumber,
		"exp":         math.MaxInt64, // Token never expired
	})
	if err != nil {
		r.JSON(http.StatusBadRequest, model.APIResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, model.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "SUCCESS",
		Data: gin.H{
			"token": token,
		},
	})
}
