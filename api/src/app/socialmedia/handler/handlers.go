package handler

import (
	"errors"
	"
	
	api/src/app/socialmedia"
	"
	
	api/src/app/socialmedia/handler/request"
	"
	
	api/src/app/socialmedia/handler/response"
	"
	
	api/src/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service socialmedia.Service
}

// @Tags Social Media
// @Summary Create social media
// @ID create-social-media
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body request.Request true "json request body"
// @Success 201 {object} response.PostResponse
// @Router /socialmedias [post]
func (h *Handler) CreateSocialMediaHandler(c *gin.Context) {
	request := request.Request{}
	if err := c.ShouldBindJSON(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, helper.ValidateRequest(verr))
			return
		}

		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userData := helper.GetUserData(c)
	result, err := h.service.CreateSocialMedia(request.MapPostToRecord(userData.ID))
	if err != nil {
		helper.CreateMessageResponse(c, err.Status(), err.Message())
		return
	}

	c.JSON(http.StatusCreated, response.MapToPostResponse(*result))
}

// @Tags Social Media
// @Summary Get all social medias
// @ID get-all-social-medias
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {array} response.Response
// @Router /socialmedias [get]
func (h *Handler) GetAllSocialMediasHandler(c *gin.Context) {
	result, err := h.service.GetAllSocialMedias()
	if err != nil {
		helper.CreateMessageResponse(c, err.Status(), err.Message())
		return
	}

	c.JSON(http.StatusOK, response.MapToBatchSocialMediaResponse(result))
}

// @Tags Social Media
// @Summary Update social media
// @ID update-social-media
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param socialMediaId path int true "socialMediaId"
// @Param RequestBody body request.Request true "json request body"
// @Success 200 {object} response.UpdateResponse
// @Router /socialmedias/{socialMediaId} [put]
func (h *Handler) UpdateSocialMediaHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("socialMediaId"))
	if id < 1 {
		helper.CreateMessageResponse(c, http.StatusBadRequest, "param must be a number greater than 0")
		return
	}

	request := request.Request{}
	if err := c.ShouldBindJSON(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, helper.ValidateRequest(verr))
			return
		}

		helper.CreateMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.service.UpdateSocialMedia(id, request.MapUpdateToRecord())
	if err != nil {
		helper.CreateMessageResponse(c, err.Status(), err.Message())
		return
	}

	c.JSON(http.StatusOK, response.MapToUpdateResponse(*result))
}

// @Tags Social Media
// @Summary Delete social media
// @ID delete-social-media
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param socialMediaId path int true "socialMediaId"
// @Success 200 {object} structs.Message
// @Router /socialmedias/{socialMediaId} [delete]
func (h *Handler) DeleteSocialMediaHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("socialMediaId"))
	if id < 1 {
		helper.CreateMessageResponse(c, http.StatusBadRequest, "param must be a number greater than 0")
		return
	}

	if err := h.service.DeleteSocialMedia(id); err != nil {
		helper.CreateMessageResponse(c, err.Status(), err.Message())
		return
	}

	helper.CreateMessageResponse(c, http.StatusOK, "Your social media has been successfully deleted")
}

func NewHandler(service socialmedia.Service) *Handler {
	return &Handler{service}
}
