package delivery

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func UploadPhotoHandler(c *gin.Context) {
	maxSize := int64(5240000) // allow only 5MB of file size

	err := c.Request.ParseMultipartForm(maxSize)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": fmt.Sprintf("Size terlalu besar. Max Size: %v", maxSize),
		})
		return
	}

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": "Tidak bisa mengambil file",
		})
		return
	}
	defer file.Close()

	// create an AWS session which can be
	// reused if we're uploading many files
	s, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("APP_ACCESS_KEY_ID"),     // id
			os.Getenv("APP_SECRET_ACCESS_KEY"), // secret
			""),                                // token can be left blank for now
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
		})
		return
	}

	fileName, err := UploadFileToS3(s, file, fileHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"responseMessage": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode":    0,
		"responseMessage": "Sukses upload gambar",
		"urlPhoto":        fileName,
	})
}

// UploadFileToS3 saves a file to aws bucket and returns the url to // the file and an error if there's any
func UploadFileToS3(s *session.Session, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	// get the file size and read
	// the file content into a buffer
	size := fileHeader.Size
	buffer := make([]byte, size)
	_, errors := file.Read(buffer)
	if errors != nil {
		return "", errors
	}

	// create a unique file name for the file
	tempFileName := "pictures/" + bson.NewObjectId().Hex() + filepath.Ext(fileHeader.Filename)

	// config settings: this is where you choose the bucket,
	// filename, content-type and storage class of the file
	// you're uploading
	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String("chaerul-bucket"),
		Key:                  aws.String(tempFileName),
		ACL:                  aws.String("public-read"), // could be private if you want it to be access by only authorized users
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("STANDARD"),
	})

	if err != nil {
		return "", err
	}

	url := "https://%s.s3.%s.amazonaws.com/%s"
	url = fmt.Sprintf(url, "chaerul-bucket", "ap-southeast-1", tempFileName)

	return url, err
}
