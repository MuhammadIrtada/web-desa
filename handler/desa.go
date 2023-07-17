package handler

import (
	"net/http"
	"web-desa/helper"
	"web-desa/model"
	"web-desa/request"

	"github.com/gin-gonic/gin"
)

type desaHandler struct {
	desaService model.DesaService
}

func NewDesaHandler(desaService model.DesaService) model.DesaHandler {
	return &desaHandler{desaService: desaService}
}

func (h *desaHandler) Mount(group *gin.RouterGroup) {
	group.POST("", h.StoreDesaHandler)
	group.GET("", h.FetchDesaHandler)
	group.PATCH("", h.EditDesaHandler)
	group.DELETE("", h.DeleteDesaHandler)
}

func (h *desaHandler) StoreDesaHandler(c *gin.Context) {
	var req request.DesaRequest
	
	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())
		return
	}

	desa, err := h.desaService.StoreDesa(&req)
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusBadRequest, err)
		return
	}

	helper.ResponseSuccessJson(c, "success", desa)
}

func (h *desaHandler) FetchDesaHandler(c *gin.Context) {
	desa, err := h.desaService.FetchDesa()
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusBadRequest, err)
		return
	}

	helper.ResponseSuccessJson(c, "success", desa)
}

func (h *desaHandler) EditDesaHandler(c *gin.Context) {
	var req request.DesaRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())
		return
	}

	desa, err := h.desaService.EditDesa(1, &req)
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
		return
	}

	helper.ResponseSuccessJson(c, "success", desa)
}

func (h *desaHandler) DeleteDesaHandler(c *gin.Context) {
	err := h.desaService.DestroyDesa()
	if err != nil {
		helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
		return
	}

	helper.ResponseSuccessJson(c, "delete success", "")
}