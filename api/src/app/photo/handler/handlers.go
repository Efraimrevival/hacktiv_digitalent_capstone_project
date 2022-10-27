package handler

import (
	"api/src/app/photo"
	"api/src/app/photo/handler/request"
	"api/src/app/photo/handler/response"
	"api/src/helper"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service photo.Service
}

// @Tags Photo
// @Summary Post photo
// @ID post-photo
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body request.Request true "json request body"
// @Success 201 {object} response.PostResponse
// @Router /photos [post]
func (h *Handler) PostPhotoHandler(c *gin.Context) {
	request := request.Request{}
	userData := helper.GetUserData(c)

	if err := c.ShouldBindJSON(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, helper.ValidateRequest(verr))
			return
		}

		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.PostPhoto(request.MapPostToRecord(userData.ID))
	if err != nil {
		helper.CreateMessageResponse(c, err.Status(), err.Message())
		return
	}

	c.JSON(http.StatusCreated, response.MapToPostResponse(*result))
}

// @Tags Photo
// @Summary Get all photos
// @ID get-all-photos
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {array} response.PhotoResponse
// @Router /photos [get]
func (h *Handler) GetAllPhotosHandler(c *gin.Context) {
	result, err := h.service.GetAllPhotos()
	if err != nil {
		helper.CreateMessageResponse(c, err.Status(), err.Message())
		return
	}

	c.JSON(http.StatusOK, response.MapToBatchPhotoResponse(result))
}

// @Tags Photo
// @Summary Update photo
// @ID update-photo
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param photoId path int true "photoId"
// @Param RequestBody body request.Request true "json request body"
// @Success 200 {object} response.UpdateResponse
// @Router /photos/{photoId} [put]
func (h *Handler) UpdatePhotoHandler(c *gin.Context) {
	request := request.Request{}
	id, _ := strconv.Atoi(c.Param("photoId"))
	if id < 1 {
		helper.CreateMessageResponse(c, http.StatusBadRequest, "param must be a number greater than 0")
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, helper.ValidateRequest(verr))
			return
		}

		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.UpdatePhoto(id, request.MapUpdateToRecord())
	if err != nil {
		helper.CreateMessageResponse(c, err.Status(), err.Message())
		return
	}

	c.JSON(http.StatusOK, response.MapToUpdateResponse(*result))
}

// @Tags Photo
// @Summary Delete photo
// @ID delete-photo
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param photoId path int true "photoId"
// @Success 200 {object} structs.Message
// @Router /photos/{photoId} [delete]
func (h *Handler) DeletePhotoHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("photoId"))
	if id < 1 {
		helper.CreateMessageResponse(c, http.StatusBadRequest, "param must be a number greater than 0")
		return
	}

	if err := h.service.DeletePhoto(id); err != nil {
		helper.CreateMessageResponse(c, err.Status(), err.Message())
		return
	}

	helper.CreateMessageResponse(c, http.StatusOK, "Your photo has been successfully deleted")
}

func NewHandler(service photo.Service) *Handler {
	return &Handler{service}
}
