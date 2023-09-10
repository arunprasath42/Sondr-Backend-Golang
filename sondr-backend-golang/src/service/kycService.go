package service

import (
	"errors"
	"mime/multipart"
	"sondr-backend/src/models"
	"sondr-backend/src/repository"
	"sondr-backend/utils/constant"
	"sondr-backend/utils/logging"
	"sondr-backend/utils/s3"

	"github.com/spf13/viper"
)

type KycService struct{}

func (ks *KycService) ListAllKycReqUsers(pageNo, pageSize int) (*models.KycResponse, error) {
	var response models.KycResponse
	var kycStatusReq []*models.ListKycstatusReq
	if pageSize == 0 {
		pageSize = 10
	}
	count, err := repository.Repo.FindAllKycStatusRequest(&kycStatusReq, pageNo, pageSize)
	if err != nil {
		return nil, err
	}
	response.KycStatusReq = kycStatusReq
	response.ReqCount = count

	return &response, nil
}

func (ks *KycService) LisAllKycVerifyStatus(pageNo, pageSize int) (*models.KycResponse, error) {
	var resp models.KycResponse
	var kycStatusVerify []*models.ListKycStatusVerify
	if pageSize == 0 {
		pageSize = 10
	}
	count, err := repository.Repo.FindAllKycStatusVerify(&kycStatusVerify, pageNo, pageSize)
	if err != nil {
		return nil, err
	}
	resp.KycStatusVerify = kycStatusVerify
	resp.VerifyCount = count
	return &resp, nil
}

func (ks *KycService) KycApprove(id uint, status string, email string) (string, error) {
	if err := repository.Repo.KycApproveAndDisApprove(id, status, email); err != nil {
		return "", err
	}
	return status + " Successfully", nil
}

func (ks *KycService) UploadKycPhotoService(userid uint, kycFile1, kycFile2 multipart.File, kycHandler1, kycHandler2 *multipart.FileHeader) (*models.UserResponse, error) {
	var kyc models.Kycs
	var user models.Users
	var response models.UserResponse
	url := viper.GetString("s3.url")
	fileName1 := "kyc/" + kycHandler1.Filename
	fileName2 := "kyc/" + kycHandler2.Filename
	errFile1 := make(chan error)
	errFile2 := make(chan error)

	if err := repository.Repo.FindById(&user, int(userid)); err != nil {
		return nil, err
	}

	go UploadFileS3(kycFile1, fileName1, kycHandler1.Size, errFile1)
	go UploadFileS3(kycFile2, fileName2, kycHandler2.Size, errFile2)

	err1, err2 := <-errFile1, <-errFile2
	if err1 != nil {
		logging.Logger.WithField("error in Uploading the s3 image1", err1).WithError(err1).Error(constant.INTERNALSERVERERROR, err1)
		return nil, err1
	}
	if err2 != nil {
		logging.Logger.WithField("error in Uploading the s3 image2", err2).WithError(err2).Error(constant.INTERNALSERVERERROR, err2)
		return nil, err2

	}

	kyc.UserId = userid
	kyc.Status = "Requested"
	kyc.KycPhoto1 = url + fileName1
	kyc.KycPhoto2 = url + fileName2
	if err := repository.Repo.InsertOrUpdateKYC(&kyc); err != nil {
		logging.Logger.Infof("error in Inserting in kyc table", err)
		return nil, err
	}
	if err := repository.Repo.UpdateUserProfile(&models.Users{}, int(userid), map[string]interface{}{
		"ProfileStatus": "Completed",
	}); err != nil {
		return nil, errors.New("unable to update user profile")
	}
	response.Message = "KYC photos uploaded successfully"
	return &response, nil
}

func UploadFileS3(file multipart.File, tempFile string, size int64, errfile chan error) {
	_, err := s3.UploadFileToS3(file, tempFile, size)
	errfile <- err
}
