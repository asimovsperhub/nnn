package api

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Read(filepath string) []byte {
	f, err := os.Open(filepath)
	if err != nil {
		log.Println("read file fail", err)
		return nil
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("read to fd fail", err)
		return nil
	}

	return fd
}
func UploadFileToS3(raw []byte, filename string) (string, error) {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-east-1"),
		Credentials: credentials.NewStaticCredentials(
			"AKIAWQJWVPP5PAETWGVC",
			"5NQBamaznpwAfBZuM7ODRlPJpOnYY06iJPt/8VpN",
			""), // Sessiontoken是进程相关，应该是连接中可以返回
	})
	if err != nil {
		log.Println("aws  failed", err)
	}
	tempFileName := "cards/" + filename

	_, err1 := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String("dnsdaonftpasscard"),
		Key:         aws.String(tempFileName),
		Body:        bytes.NewReader(raw),
		ContentType: aws.String(http.DetectContentType(raw)),
	})
	if err1 != nil {
		log.Println("PUT err", err1)
		return "", err1
	}
	fileUrl := "https://dnsdaonftpasscard.s3.ap-east-1.amazonaws.com/" + tempFileName

	return fileUrl, nil
}
