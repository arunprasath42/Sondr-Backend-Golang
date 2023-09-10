package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"sondr-backend/src/models"
	"sondr-backend/src/service"
	"sondr-backend/utils/constant"
	"sondr-backend/utils/response"
	"sondr-backend/utils/validator"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListAllKycRequestStatus(c *gin.Context) {
	reqModel := &models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if reqModel.Status != "Requested" {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, errors.New("status is invalid or missing")))
		return
	}
	ks := service.KycService{}
	resp, err := ks.ListAllKycReqUsers(reqModel.PageNo, reqModel.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(resp))

}
func ListAllKycVerifiedStatus(c *gin.Context) {
	reqModel := &models.Request{}
	if err := c.ShouldBindQuery(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if reqModel.Status != "Approved" {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, errors.New("status is invalid or missing")))
		return
	}
	ks := service.KycService{}
	resp, err := ks.LisAllKycVerifyStatus(reqModel.PageNo, reqModel.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(resp))

}

func KycApprove(c *gin.Context) {
	verifedBy := c.GetString("email")
	reqModel := &models.Request{}
	if err := c.ShouldBindJSON(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	fmt.Println(reqModel.UserId)
	if err := validator.ValidateVariable(reqModel.UserId, "required", "userId"); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}
	if reqModel.Status != "Approved" && reqModel.Status != "DisApproved" {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, errors.New("status is invalid or missing")))
		return
	}
	ks := service.KycService{}
	resp, err := ks.KycApprove(reqModel.UserId, reqModel.Status, verifedBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(resp))

}

func UploadKycPhotos(c *gin.Context) {
	//var req models.Request
	id := c.GetString("id")
	userid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	kycFile1, kycHandler1, err := c.Request.FormFile("kycphoto1")

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}

	kycFile2, kycHandler2, err := c.Request.FormFile("kycphoto2")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(constant.BADREQUEST, err))
		return
	}

	//fmt.Println("id", id)
	ks := service.KycService{}
	resp, err := ks.UploadKycPhotoService(uint(userid), kycFile1, kycFile2, kycHandler1, kycHandler2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(constant.INTERNALSERVERERROR, err))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(resp))
}
