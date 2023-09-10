package test

import (
	"mime/multipart"
	"os"
	"sondr-backend/src/service"
	"testing"
)

func TestUploadKycPhotoServiceValid(t *testing.T) {
	type args struct {
		userid      uint
		kycFile1    multipart.File
		kycFile2    multipart.File
		kycHandler1 *multipart.FileHeader
		kycHandler2 *multipart.FileHeader
	}
	var err error
	var fileheader1 multipart.FileHeader
	var fileheader2 multipart.FileHeader
	file1, err := os.Open("../../demo.jpg")
	if err != nil {
		t.Error("error in file length", err)
	}
	ff1, err := file1.Stat()
	if err != nil {
		t.Error("error in file length", err)
	}
	fileheader1.Size = ff1.Size()
	fileheader1.Filename = ff1.Name()

	//file1.Close()
	file2, err := os.Open("../../demo.jpg")
	if err != nil {
		t.Error("error in file2 length", err)
	}
	ff2, err := file2.Stat()
	if err != nil {
		t.Error("error in file length", err)
	}
	fileheader2.Size = ff2.Size()
	fileheader2.Filename = ff2.Name()
	//	file2.Close()

	//	var file multipart.Form

	test := args{
		//pass the parameter here
		kycFile1:    file2,
		kycFile2:    file1,
		kycHandler1: &fileheader2,
		kycHandler2: &fileheader1,
		userid:      1,
	}
	// Add External package name in ____
	ks := &service.KycService{}
	_, err = ks.UploadKycPhotoService(test.userid, test.kycFile1, test.kycFile2, test.kycHandler1, test.kycHandler2)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestUploadKycPhotoServiceInvalid(t *testing.T) {
	type args struct {
		userid      uint
		kycFile1    multipart.File
		kycFile2    multipart.File
		kycHandler1 *multipart.FileHeader
		kycHandler2 *multipart.FileHeader
	}
	var err error
	var fileheader1 multipart.FileHeader
	var fileheader2 multipart.FileHeader
	file1, err := os.Open("../../demo.jpg")
	if err != nil {
		t.Error("error in file length", err)
	}
	ff1, err := file1.Stat()
	if err != nil {
		t.Error("error in file length", err)
	}
	fileheader1.Size = ff1.Size()
	fileheader1.Filename = ff1.Name()

	//file1.Close()
	file2, err := os.Open("../../demo.jpg")
	if err != nil {
		t.Error("error in file2 length", err)
	}
	ff2, err := file2.Stat()
	if err != nil {
		t.Error("error in file length", err)
	}
	fileheader2.Size = ff2.Size()
	fileheader2.Filename = ff2.Name()
	test := args{
		//pass the parameter here
		userid:      0,
		kycFile1:    file2,
		kycFile2:    file1,
		kycHandler1: &fileheader2,
		kycHandler2: &fileheader1,
	}
	// Add External package name in ____
	ks := &service.KycService{}
	_, err = ks.UploadKycPhotoService(test.userid, test.kycFile1, test.kycFile2, test.kycHandler1, test.kycHandler2)

	if err == nil {
		t.Error(err.Error())
	}

}

func TestListAllKycReqUsersValid(t *testing.T) {
	type args struct {
		pageNo   int
		pageSize int
	}

	test := args{
		//pass the parameter here
		pageNo:   1,
		pageSize: 10,
	}
	// Add External package name in ____
	ks := &service.KycService{}
	_, err := ks.ListAllKycReqUsers(test.pageNo, test.pageSize)

	if err != nil {
		t.Error(err.Error())
	}

}

// func TestListAllKycReqUsersInvalid(t *testing.T) {
// 	type args struct {
// 		pageNo   int
// 		pageSize int
// 	}

// 	test := args{
// 		//pass the parameter here

// 	}
// 	// Add External package name in ____
// 	ks := &service.KycService{}
// 	_, err := ks.ListAllKycReqUsers(test.pageNo, test.pageSize)

// 	if err == nil {
// 		t.Error(err.Error())
// 	}

// }

func TestLisAllKycVerifyStatusValid(t *testing.T) {
	type args struct {
		pageNo   int
		pageSize int
	}

	test := args{
		//pass the parameter here
		pageNo:   1,
		pageSize: 10,
	}
	// Add External package name in ____
	ks := &service.KycService{}
	_, err := ks.LisAllKycVerifyStatus(test.pageNo, test.pageSize)

	if err != nil {
		t.Error(err.Error())
	}

}

// func TestLisAllKycVerifyStatusInvalid(t *testing.T) {
// 	type args struct {
// 		pageNo   int
// 		pageSize int
// 	}

// 	test := args{
// 		//pass the parameter here

// 	}
// 	// Add External package name in ____
// 	ks := &service.KycService{}
// 	_, err := ks.LisAllKycVerifyStatus(test.pageNo, test.pageSize)

// 	if err == nil {
// 		t.Error(err.Error())
// 	}

// }

func TestKycApproveValid(t *testing.T) {
	type args struct {
		id     uint
		status string
		email  string
	}

	test := args{
		//pass the parameter here
		id:     1,
		status: "Approved",
		email:  "admin@sondr.com",
	}
	// Add External package name in ____
	ks := &service.KycService{}
	_, err := ks.KycApprove(test.id, test.status, test.email)

	if err != nil {
		t.Error(err.Error())
	}

}

func TestKycApproveInvalid(t *testing.T) {
	type args struct {
		id     uint
		status string
		email  string
	}

	test := args{
		//pass the parameter here
		id:     0,
		status: "DisApprove",
		email:  "admin@sondr.com",
	}
	// Add External package name in ____
	ks := &service.KycService{}
	_, err := ks.KycApprove(test.id, test.status, test.email)

	if err == nil {
		t.Error(err.Error())
	}

}
