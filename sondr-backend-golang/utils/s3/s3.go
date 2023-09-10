package s3

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
)

var ss *session.Session

func ConnectS3() error {
	//connect
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(viper.GetString("s3.region")),
		Credentials: credentials.NewStaticCredentials(viper.GetString("s3.accesKey"), viper.GetString("s3.secretKey"), ""),
	})
	if err != nil {
		log.Println("error connecting to s3")
		return err
	}
	ss = sess
	return nil
}

func UploadFileToS3(file multipart.File, TempFileName string, size int64) (string, error) {

	//	s := ConnectS3()

	filebytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Print("error in file reading", err, TempFileName)
		return "", err
	}

	_, err = s3.New(ss).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(viper.GetString("s3.bucket")),
		Key:                  aws.String(TempFileName),
		ACL:                  aws.String("public-read"), // could be private if you want it to be access by only authorized users
		Body:                 bytes.NewReader(filebytes),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(filebytes)),
		ServerSideEncryption: aws.String("AES256"),
	})
	if err != nil {
		fmt.Println("error in inserting the file", err)
		return "", err
	}
	return TempFileName, err
}
